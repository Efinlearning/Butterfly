// ═══════════════════════════════════════════════════════════════════
// FYERS backend (Go) — OAuth + Option Chain proxy
// Deploy this on Render as a "Web Service" (Go environment). It is the
// only place your FYERS App ID / Secret ID / Access Token ever live —
// the browser never sees them. Standard library only, no dependencies.
// ═══════════════════════════════════════════════════════════════════
package main

import (
	"bufio"
	"bytes"
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"regexp"
	"strings"
	"sync"
	"time"
)

// FYERS splits its endpoints across two hosts:
//   - api-t1.fyers.in — the interactive login/authorize page (generate-authcode).
//     Hitting this with the plain API host instead returns a JSON 500
//     "Invalid Request, please provide valid method" error.
//   - api.fyers.in     — the pure REST/JSON API (token exchange, quotes, option chain).
const fyersAuthHost = "https://api-t1.fyers.in/api/v3"
const fyersBase = "https://api.fyers.in/api/v3"

var (
	fyersAppID     string
	fyersSecretID  string
	publicURL      string
	frontendOrigin string
	port           string
)

// ── In-memory token store (single-user tool: one dashboard, one FYERS login).
// FYERS access tokens expire daily, so there's no need for a database —
// just reconnect each trading morning via the Connect button. ──
type tokenStore struct {
	mu          sync.RWMutex
	accessToken string
	expiresAt   time.Time
	state       string
}

var store tokenStore

func (s *tokenStore) connected() bool {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return s.accessToken != "" && time.Now().Before(s.expiresAt)
}

func (s *tokenStore) set(token string, ttl time.Duration) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.accessToken = token
	s.expiresAt = time.Now().Add(ttl)
}

func (s *tokenStore) clear() {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.accessToken = ""
	s.expiresAt = time.Time{}
}

func (s *tokenStore) get() string {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return s.accessToken
}

func (s *tokenStore) setState(state string) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.state = state
}

func (s *tokenStore) checkState(state string) bool {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return state != "" && state == s.state
}

// ── Minimal .env loader (no dependencies). Reads KEY=VALUE lines from a
// .env file in the working directory and sets them as process env vars,
// without overriding any var that's already set (so real environment
// variables — e.g. ones set in Render's dashboard — always win). ──
func loadDotEnv(path string) {
	f, err := os.Open(path)
	if err != nil {
		return // no .env file present — that's fine, fall back to real env vars
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" || strings.HasPrefix(line, "#") {
			continue
		}
		eq := strings.Index(line, "=")
		if eq < 0 {
			continue
		}
		key := strings.TrimSpace(line[:eq])
		val := strings.TrimSpace(line[eq+1:])
		// Strip surrounding quotes, if any: KEY="value" or KEY='value'
		if len(val) >= 2 && (val[0] == '"' && val[len(val)-1] == '"' || val[0] == '\'' && val[len(val)-1] == '\'') {
			val = val[1 : len(val)-1]
		}
		if key == "" {
			continue
		}
		if _, already := os.LookupEnv(key); !already {
			os.Setenv(key, val)
		}
	}
}

func main() {
	loadDotEnv(".env")

	fyersAppID = os.Getenv("FYERS_APP_ID")
	fyersSecretID = os.Getenv("FYERS_SECRET_ID")
	publicURL = strings.TrimRight(os.Getenv("PUBLIC_URL"), "/")
	frontendOrigin = strings.TrimRight(os.Getenv("FRONTEND_ORIGIN"), "/")
	port = os.Getenv("PORT")
	if port == "" {
		port = "10000"
	}

	if fyersAppID == "" || fyersSecretID == "" || publicURL == "" || frontendOrigin == "" {
		log.Fatal("Missing required env vars: FYERS_APP_ID, FYERS_SECRET_ID, PUBLIC_URL, FRONTEND_ORIGIN")
	}

	mux := http.NewServeMux()
	mux.HandleFunc("/health", handleHealth)
	mux.HandleFunc("/auth/login", handleLogin)
	mux.HandleFunc("/auth/callback", handleCallback)
	mux.HandleFunc("/auth/status", handleStatus)
	mux.HandleFunc("/auth/disconnect", handleDisconnect)
	mux.HandleFunc("/api/option-chain", handleOptionChain)
	mux.HandleFunc("/api/fyers/", handleProxy)

	log.Printf("FYERS backend listening on :%s", port)
	log.Fatal(http.ListenAndServe(":"+port, withCORS(mux)))
}

// ── CORS ──
func withCORS(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", frontendOrigin)
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusNoContent)
			return
		}
		next.ServeHTTP(w, r)
	})
}

func writeJSON(w http.ResponseWriter, status int, v any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(v)
}

// ── GET /health ──
func handleHealth(w http.ResponseWriter, r *http.Request) {
	writeJSON(w, http.StatusOK, map[string]any{"ok": true, "connected": store.connected()})
}

// ── GET /auth/login — the frontend opens this URL in a popup ──
func handleLogin(w http.ResponseWriter, r *http.Request) {
	state := randomHex(16)
	store.setState(state)

	redirectURI := publicURL + "/auth/callback"
	q := url.Values{}
	q.Set("client_id", fyersAppID)
	q.Set("redirect_uri", redirectURI)
	q.Set("response_type", "code")
	q.Set("state", state)

	authURL := fyersAuthHost + "/generate-authcode?" + q.Encode()
	http.Redirect(w, r, authURL, http.StatusFound)
}

// ── GET /auth/callback?auth_code=&state= — FYERS redirects the popup here ──
func handleCallback(w http.ResponseWriter, r *http.Request) {
	authCode := r.URL.Query().Get("auth_code")
	if authCode == "" {
		authCode = r.URL.Query().Get("code")
	}
	state := r.URL.Query().Get("state")

	closeWith := func(msgType, errMsg, brokerName string) {
		payload := map[string]string{"type": msgType}
		if errMsg != "" {
			payload["error"] = errMsg
		}
		if brokerName != "" {
			payload["brokerName"] = brokerName
		}
		b, _ := json.Marshal(payload)
		originJSON, _ := json.Marshal(frontendOrigin)
		w.Header().Set("Content-Type", "text/html")
		fmt.Fprintf(w, `<!doctype html><html><body>
<script>
  window.opener && window.opener.postMessage(%s, %s);
  window.close();
</script>
<p>You can close this window.</p>
</body></html>`, string(b), string(originJSON))
	}

	if authCode == "" || state == "" || !store.checkState(state) {
		closeWith("BROKER_AUTH_ERROR", "Invalid or expired login attempt. Please try again.", "")
		return
	}

	appIDHash := sha256.Sum256([]byte(fyersAppID + ":" + fyersSecretID))
	body, _ := json.Marshal(map[string]string{
		"grant_type": "authorization_code",
		"appIdHash":  hex.EncodeToString(appIDHash[:]),
		"code":       authCode,
	})

	req, err := http.NewRequest(http.MethodPost, fyersBase+"/validate-authcode", bytes.NewReader(body))
	if err != nil {
		closeWith("BROKER_AUTH_ERROR", err.Error(), "")
		return
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		closeWith("BROKER_AUTH_ERROR", err.Error(), "")
		return
	}
	defer resp.Body.Close()

	var d struct {
		S           string `json:"s"`
		Message     string `json:"message"`
		AccessToken string `json:"access_token"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&d); err != nil {
		closeWith("BROKER_AUTH_ERROR", "Failed to parse FYERS response", "")
		return
	}
	if d.S != "ok" || d.AccessToken == "" {
		msg := d.Message
		if msg == "" {
			msg = "Token exchange failed"
		}
		closeWith("BROKER_AUTH_ERROR", msg, "")
		return
	}

	// FYERS tokens expire at a market-adjacent cutoff; treat as good for 8h to be safe.
	store.set(d.AccessToken, 8*time.Hour)
	closeWith("BROKER_AUTH_SUCCESS", "", "FYERS")
}

// ── GET /auth/status ──
func handleStatus(w http.ResponseWriter, r *http.Request) {
	store.mu.RLock()
	expiresAt := store.expiresAt
	store.mu.RUnlock()
	resp := map[string]any{"connected": store.connected()}
	if !expiresAt.IsZero() {
		resp["expiresAt"] = expiresAt
	} else {
		resp["expiresAt"] = nil
	}
	writeJSON(w, http.StatusOK, resp)
}

// ── POST /auth/disconnect ──
func handleDisconnect(w http.ResponseWriter, r *http.Request) {
	store.clear()
	writeJSON(w, http.StatusOK, map[string]bool{"ok": true})
}

// ── auth guard ──
func requireAuth(w http.ResponseWriter, r *http.Request) bool {
	if !store.connected() {
		writeJSON(w, http.StatusUnauthorized, map[string]string{
			"s": "error", "message": "Not connected to FYERS — click Connect first.",
		})
		return false
	}
	return true
}

var proxyPrefix = regexp.MustCompile(`^/api/fyers/`)

// ── GET /api/fyers/<anything> → GET https://api.fyers.in/api/v3/<anything> ──
// Covers everything the dashboard needs:
//
//	/api/fyers/quotes?symbols=...
//	/api/fyers/data/option-chain?symbol=...&strikecount=...
func handleProxy(w http.ResponseWriter, r *http.Request) {
	if !requireAuth(w, r) {
		return
	}
	path := proxyPrefix.ReplaceAllString(r.URL.Path, "")
	target := fyersBase + "/" + path
	if r.URL.RawQuery != "" {
		target += "?" + r.URL.RawQuery
	}
	fetchAndRelay(w, target)
}

// ── GET /api/option-chain?symbol=NSE:NIFTY50-INDEX&strikecount=25 ──
func handleOptionChain(w http.ResponseWriter, r *http.Request) {
	if !requireAuth(w, r) {
		return
	}
	symbol := r.URL.Query().Get("symbol")
	if symbol == "" {
		symbol = "NSE:NIFTY50-INDEX"
	}
	strikecount := r.URL.Query().Get("strikecount")
	if strikecount == "" {
		strikecount = "25"
	}
	q := url.Values{}
	q.Set("symbol", symbol)
	q.Set("strikecount", strikecount)
	target := fyersBase + "/data/option-chain?" + q.Encode()
	fetchAndRelay(w, target)
}

func fetchAndRelay(w http.ResponseWriter, target string) {
	req, err := http.NewRequest(http.MethodGet, target, nil)
	if err != nil {
		writeJSON(w, http.StatusBadGateway, map[string]string{"s": "error", "message": err.Error()})
		return
	}
	req.Header.Set("Authorization", fyersAppID+":"+store.get())

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		writeJSON(w, http.StatusBadGateway, map[string]string{"s": "error", "message": err.Error()})
		return
	}
	defer resp.Body.Close()

	raw, err := io.ReadAll(resp.Body)
	if err != nil {
		writeJSON(w, http.StatusBadGateway, map[string]string{"s": "error", "message": err.Error()})
		return
	}

	// If FYERS says the token is dead, forget it so the UI shows "disconnected"
	var d struct {
		S       string `json:"s"`
		Message string `json:"message"`
	}
	if json.Unmarshal(raw, &d) == nil && d.S == "error" && strings.Contains(strings.ToLower(d.Message), "token") {
		store.clear()
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(resp.StatusCode)
	w.Write(raw)
}

func randomHex(n int) string {
	b := make([]byte, n)
	if _, err := rand.Read(b); err != nil {
		// Extremely unlikely; fall back to a time-based value rather than crash.
		return hex.EncodeToString([]byte(time.Now().String()))
	}
	return hex.EncodeToString(b)
}
