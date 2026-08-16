# NIFTY 50 — Multi-Expiry Butterfly 1:2:1
# Complete Trade Entry & Exit Rules
### Version 3.0 — Evidence-Based Update

---

## DOCUMENT PURPOSE

This document encodes every entry, management, and exit rule
for the NIFTY 50 Multi-Expiry Butterfly strategy. Rules are
derived from three sources:

1. Core strategy taught in the course (Value Concept,
   Chain Concept, Straddle Calibration)
2. Live data evidence from the 22856 session analysis
3. Backtesting observations on NSE Bhavcopy data

Every rule has a reason. The reason is written next to the rule.
No rule should be followed blindly — understand the reason first.

---

## PART 1 — STRATEGY PHILOSOPHY

### 1.1 What We Are

We are NET OPTION SELLERS structured as butterflies.
Profit source : theta decay + VIX contraction
NOT from : predicting market direction
Position type : non-directional (delta-gamma neutral by construction)
Greek we manage: VEGA only (VIX is the single active risk)


### 1.2 What We Are NOT

We are NOT naked sellers — all sells are covered by buys
We are NOT directional traders — no naked calls, no naked puts
We are NOT momentum players — we do not predict moves
We are NOT short-term scalpers — minimum 1 day holding


### 1.3 Core Equation

PROFIT = Premium Erosion (Theta) + VIX Contraction (Vega gain)
LOSS = VIX Expansion (Vega loss) + Adverse move beyond tolerance


When market moves big → VIX spikes → BF values drop (M2M loss)
When market stabilizes → VIX falls → BF values rise (profit)
M2M loss from VIX spike is NOT real loss — it is a scaling opportunity.

---

## PART 2 — STRUCTURE DEFINITION

### 2.1 The 1:2:1 Butterfly

CALL BUTTERFLY:
Buy 1 lot at K (First Leg)
Sell 2 lots at K + gap (Middle Leg — this is where max profit sits)
Buy 1 lot at K + 2×gap (Last Leg — hedge)

PUT BUTTERFLY:
Buy 1 lot at K (First Leg)
Sell 2 lots at K − gap (Middle Leg)
Buy 1 lot at K − 2×gap (Last Leg — hedge)

Net position: Debit paid (money spent to enter)
Max profit : At expiry when spot = middle leg strike
Max loss : Entry debit (limited, defined risk)


### 2.2 Multi-Expiry Version

First Leg (buy) → Next Week or Monthly expiry
Middle Leg (sell) → Same expiry as first leg
Last Leg (buy) → Current Weekly expiry (cheap hedge)

Why multi-expiry:

Wider range of movement survivable
More theta captured from longer-dated legs
Hedge leg (last) costs almost nothing (1-3 Rs)
Gives margin benefit that allows larger position

### 2.3 Valid Combinations (Lot Ratios)
Combination	Lot Ratio	Exit Shows	Margin/Lot	When to Use
1:2:1	1 buy:2sell	2.0	~₹45,000	Straddle ≤ 160
2:3:2	2 buy:3sell	1.5	~₹35,000	Straddle 160-280
3:4:3	3 buy:4sell	1.33	~₹28,000	Straddle > 280

Note: fractional ratios (1.5 buy) are physically impossible.
So 2:3:2 means: buy 2 first legs, sell 3 middles, buy 2 last legs.
The 0.5 fractional is approximated by using 2:3 ratio.
Hedge last leg is personal judgment — use cheapest OTM strike
available at 1-3 Rs. Not included in value% calculations.


---

## PART 3 — PRE-ENTRY CALIBRATION

### 3.1 Step 1 — Read ATM Straddle

Before any trade, check the current ATM straddle.
ATM Straddle = ATM Call LTP + ATM Put LTP

This single number tells you:
→ How much gap to use
→ Which combination to use
→ How much adverse move you can absorb
→ Whether to use debit BF or back ratio


### 3.2 Step 2 — Calculate Recommended Gap

Formula: Gap = ATM Straddle ÷ 2, rounded to nearest 50

Straddle ≤ 140 → Gap = 50 pts
Straddle 140-240 → Gap = 100 pts ← most common
Straddle 240-380 → Gap = 150 pts
Straddle > 380 → Gap = 200 pts

Example: Straddle = 167.8 → 167.8÷2 = 83.9 → nearest 50 = 100 pts

Why: Gap defines the profit window. Larger gap = wider window =
more room for market to move while staying profitable.
Gap should roughly equal expected weekly move (straddle ÷ 2).


### 3.3 Step 3 — Select Combination

Straddle ≤ 160 → 1:2 combination (exit shows 2.0 per lot)
Straddle 160-280 → 2:3 combination (exit shows 1.5 per lot)
Straddle > 280 → 3:4 combination (exit shows 1.33 per lot)

Why: Higher straddle = more premium available = can afford to sell
more lots for the same risk. Lower straddle = sell fewer lots.


### 3.4 Step 4 — Adverse Move Tolerance

Your position can absorb adverse moves up to ≈ straddle value
without requiring adjustment.

Straddle 167.8 → tolerate ±167 pts move
Straddle 300 → tolerate ±300 pts move
Straddle 500 → tolerate ±500 pts move

This is derived from the premium in the position — not a rule
you set, but what the market itself is pricing in.

DTE at entry vs tolerance:
Expiry day → survive ±50 pts
Day before expiry → survive ±100 pts
Monday (weekly) → survive ±200-250 pts
Next week entry → survive ±300-400 pts
Monthly entry → survive ±500-600 pts


### 3.5 Step 5 — VIX Check

VIX LEVEL tells you the regime:
VIX < 12 → Very low → premiums crushed → consider back ratio
VIX 12-14 → Normal → ideal debit BF zone
VIX 14-17 → Moderate → good BF entry, scale carefully
VIX 17-20 → Elevated → excellent BF entry (more premium)
VIX > 20 → High → scale aggressively on dips

VIX DIRECTION is MORE important than level:
VIX falling today → BF values rising → STRONG ENTRY signal (+30 score)
VIX rising today → BF values falling → M2M loss expected
→ still enter if value% is ideal, else wait
VIX stable → Normal entry

Combined check:
VIX > 17 AND falling → BEST possible entry timing
VIX < 12 AND falling → Check back ratio instead
VIX rising sharply → Wait for peak, then enter on the turn


### 3.6 Step 6 — Straddle Overshoot Signal (NEW RULE)

Overshoot Ratio = Today's market move ÷ ATM Straddle

Ratio > 1.2 → Market moved MORE than weekly expectation today
→ Highest probability stabilization follows
→ PRIME butterfly entry (+40 score)
Ratio 0.9-1.2 → Near full range used → Good entry (+20 score)
Ratio 0.5-0.9 → Moderate move → Normal entry (+5 score)
Ratio < 0.5 → Market barely moved → Theta less urgent (-5 score)

Example from live data:
Market moved 235 pts, straddle = 167.8
Ratio = 235 / 167.8 = 1.40 → PRIME ENTRY
On such days, BF entry is highest priority.

Why: A market that overshoots its weekly expected range in one
session almost always stabilizes for remaining days of the week.
This is the best theta-extraction setup.


### 3.7 Step 7 — OI-Based Middle Strike Selection (NEW RULE)

Do NOT always use ATM as middle strike.
Use the strike with HIGHEST OI BUILD within ±150 pts of ATM.

For CALL butterfly:
Look at call OI change column in option chain
Find strike above ATM with highest positive OI change
→ This is where market makers are building resistance = pin target
→ Place your MIDDLE leg here

For PUT butterfly:
Look at put OI change column
Find strike below ATM with highest positive OI change
→ This is where support is being built = put pin target
→ Place your MIDDLE leg here

Example from live data:
ATM = 22850
Call OI change: 22900 = +7,105,725 ← highest
23000 = +9,207,900 ← second highest
→ Place call BF middle at 22900 (50 pts above ATM)
→ NOT at ATM 22850

Why: Market makers pin option prices at strikes where they have
maximum short interest. OI build shows WHERE they want price.
Your middle leg should be at their target, not at your guess.


### 3.8 Step 8 — PCR Directional Bias

PCR = Total Put OI / Total Call OI

PCR < 0.7 → Bullish OI structure (more calls written = resistance)
→ Shift call BF middles 50-100 pts higher than OI target
→ Put BFs center at or near ATM

PCR 0.7-1.2 → Neutral → symmetric placement
→ Both sides balanced

PCR 1.2-1.8 → Bearish OI structure → support being built
→ Put BF middles 50 pts lower than OI target
→ Call BFs center at ATM or slightly higher

PCR > 1.8 → Extreme bearish → reduce call BF, increase put BF

Example: PCR = 0.70 (live data) → Bullish → calls at higher strikes


---

## PART 4 — VALUE CONCEPT (ENTRY GATE)

### 4.1 Core Formula

Value% = BF Rate ÷ FLP × 100

Where:
BF Rate = Net debit = First Leg LTP − (2 × Middle Leg LTP) + Last Leg LTP
FLP = First Leg Premium = LTP of the buying leg at strike K

This formula ensures the position is non-directional.
When Value% ≤ 20%, delta and gamma are automatically near zero.
You do NOT need to manually hedge. The structure does it.


### 4.2 Entry Thresholds

Value% Grade Action

≤ 10% IDEAL Enter immediately. Allocate 3× base lots.
Structure is maximally non-directional.
Best possible R:R ratio.

10.1% – 15% GOOD Enter. Allocate 2× base lots.
Solid non-directional structure.
Minor directional bias acceptable.

15.1% – 20% PASS Enter. Allocate 1× base lots (base size).
Position is still non-directional.
Do not go above base lots here.

20.1% – 25% BORDERLINE Enter ONLY if:
(a) Chain concept shows ≥2/3 scenarios pass
(b) VIX is falling strongly (>0.5 drop today)
(c) Overshoot ratio > 1.0
If ANY of these 3 fail → SKIP this strike

25% FAIL NEVER ENTER. Non-negotiable hard stop.
Position will be directional.
No exception regardless of other signals.


### 4.3 FLP Minimums (Data Quality Gate)

Before calculating Value%, verify:

FLP must be > ₹8 If FLP < ₹8 → option is too deep OTM
or data is stale → skip this strike

BF Rate must be > ₹0.5 If BF < ₹0.5 → spread too tight
or negative → pricing anomaly → skip

ATM Straddle must be > ₹50 If straddle < ₹50 → premiums
are crushed → back ratio zone, not BF

These gates prevent entering on bad data or wrong market conditions.


### 4.4 Tiered Position Sizing

Value% ≤ 10% AND Score ≥ 120 → 3× base lots
Value% ≤ 15% AND Score ≥ 90 → 2× base lots
Value% ≤ 20% AND Score ≥ 60 → 1× base lots (standard)
Value% 20-25% AND conditions → 0.5× base lots (half size)
Value% > 25% → 0 lots (do not enter)

Base lots = whatever your standard unit is (e.g., 2 buy lots of 2:3)

Why tiered sizing:
Live data proof: 4.2% Value% gives R:R of 1:7.5
17.6% Value% gives R:R of 1:4.6
Difference: 63% better R:R at ideal entry.
It makes mathematical sense to bet more on better setups.


### 4.5 Entry Scoring System

Score combines all factors. Minimum score to enter = 60.

FACTOR 1 — Value% (max 60 points)
≤ 10% → 60 pts
≤ 15% → 45 pts
≤ 20% → 25 pts
≤ 25% → 5 pts

25% → REJECT (return -999, do not enter)

FACTOR 2 — VIX direction (max 40 points)
VIX falling > 1.0 today → +40 pts (best condition)
VIX falling 0.5-1.0 → +25 pts
VIX falling 0-0.5 → +10 pts
VIX rising 0-1.0 → 0 pts
VIX rising > 1.0 → -20 pts (bad, avoid if possible)

FACTOR 3 — Overshoot ratio (max 40 points)
Ratio > 1.2 → +40 pts (prime entry — market overshot range)
Ratio 0.9-1.2 → +20 pts
Ratio 0.5-0.9 → +5 pts
Ratio < 0.5 → -10 pts

FACTOR 4 — Double Middle Lock (30 points)
Call middle strike = Put middle strike → +30 pts
(Both sides peak at same spot = maximum combined profit)
Call middle ≠ Put middle → 0 pts

FACTOR 5 — PCR neutral zone (10 points)
PCR 0.7-1.3 → +10 pts (neutral = butterfly friendly)
PCR < 0.5 → -10 pts (extreme — directional risk)
PCR > 1.8 → -10 pts

FACTOR 6 — Chain Concept (20 points)
3 of 3 scenarios survive → +20 pts
2 of 3 scenarios survive → +10 pts
1 of 3 scenarios survive → 0 pts
0 of 3 scenarios survive → -20 pts (and reject entry if borderline)

Entry decision:
Score ≥ 120 → Strong entry, use tiered sizing
Score 60-119 → Normal entry, use base sizing
Score < 60 → Skip this trade, look for better strike/timing


---

## PART 5 — CHAIN CONCEPT (SCENARIO VALIDATION)

### 5.1 What Chain Concept Is

Before entering any trade, simulate what happens to your butterfly
value in 3 scenarios:

Scenario 1 — Market UP one adverse move (= straddle pts up)
Scenario 2 — Market FLAT (theta erodes premium)
Scenario 3 — Market DOWN one adverse move (= straddle pts down)

For each scenario, compute the BF rate at the adjacent strikes
in the option chain. If BF retains ≥ 40% of entry value in that
scenario → scenario "survives."

Need minimum 2 of 3 scenarios to survive before entering.


### 5.2 How to Run Chain Check

Look at the option chain exit tool (butterfly rates column).

For a 100-gap Call BF at strike 22800 (straddle = 167):
Adverse up : Check 22800 CE BF if market = 22800+167 = 22967
→ Look at 23000 CE butterfly rates
Flat : BF rate stays at current value or improves by ~10%
(theta works for you)
Adverse down : Check 22800 CE BF if market = 22800-167 = 22633
→ Look at 22650 CE butterfly rates

Survival threshold = 40% of entry BF rate:
Entry BF = 18 Rs
Threshold = 18 × 0.40 = 7.2 Rs
If simulated BF in that scenario ≥ 7.2 → scenario survives

Two of three must survive.


### 5.3 Chain Concept During Trade Management

While trade is open, continuously ask:
"If market moves X more pts from HERE, what does my BF become?"

Use adjacent strike BF rates from the exit tool to answer.

Trade going ITM → Check chain HEAVILY
→ If chain shows BF → 0 on further adverse move → EXIT
→ If chain shows BF stays reasonable → HOLD

Trade going OTM → Check VALUE% (Hold% rule) more heavily
→ Chain less relevant once deeply OTM

Market stabilizing → Chain confirms hold → collect theta


---

## PART 6 — HOLD RULES (POSITION MANAGEMENT)

### 6.1 Hold% Formula

Hold% = Entry Debit ÷ Current FLP × 100

Where:
Entry Debit = what you paid when entering (fixed, never changes)
Current FLP = live price of first leg option RIGHT NOW

What Hold% tells you:
When FLP was high (entry) → Hold% is low → good, hold
When FLP falls (position going OTM) → Hold% rises → warning
When Hold% > 20% → position is becoming directional → review
When Hold% > 25% → position IS directional → EXIT NOW


### 6.2 Individual Leg Hold Rules

Hold% ≤ 15% → HOLD comfortably. Theta working. No action.

Hold% 15-20% → HOLD but monitor closely.
Check chain concept for adjacent BF rates.
If chain still favorable → hold.
If chain shows further deterioration → prepare to exit.

Hold% 20-25% → REVIEW immediately.
Three checks:
(a) Is the OTHER side (call/put) compensating?
(b) Does chain concept show improvement possible?
(c) Is straddle falling (VIX dropping)?
If YES to any → hold one more session.
If NO to all → EXIT this specific leg.

Hold% > 25% → EXIT IMMEDIATELY. Non-negotiable.
This leg has become directional.
Book at whatever price is available.
Do not wait for "recovery."
Re-enter at closer-to-ATM strike after exit.


### 6.3 Combined Portfolio Hold Rule (NEW RULE)

When running both call and put butterflies together:

Combined Hold% = Total Entry Debit ÷ Total Current FLP × 100

If Call Hold% = 28% BUT Put Hold% = 8%:
Combined = (call debit + put debit) / (call FLP + put FLP) × 100

If combined < 20% → hold both (put side is compensating call side)
If combined > 20% → exit the breached leg, keep the compensating leg
If combined > 25% → exit both

Why: One leg deteriorating does not mean the strategy failed.
The other side may be making money simultaneously.
Evaluate the pair, not each leg in isolation.


### 6.4 ITM Position Rules

When your butterfly goes ITM (market moves past first leg):

Deeply ITM (first leg 100+ pts ITM):
→ Follow chain concept HEAVILY
→ Check: if market moves X more pts, what does BF become?
→ If BF → 0 on further adverse move → EXIT NOW at small loss
→ Do not hold deeply ITM positions hoping for reversal
→ Small loss now is better than full debit loss later

Slightly ITM (first leg 0-50 pts ITM):
→ Monitor both value% and chain
→ If Hold% < 15% → hold (still has time value)
→ If Hold% 15-25% → review with chain check
→ If Hold% > 25% → exit

Special ITM rule: When deeply ITM AND Hold% > 15% → EXIT
(Lower threshold for ITM because further adverse moves accelerate loss)


---

## PART 7 — EXIT RULES

### 7.1 Stop Loss Rule

Exit when: Current BF < Entry Debit ÷ 3

Example: Entry debit = 18 Rs
Stop loss trigger = 18 ÷ 3 = 6 Rs
If BF falls to 6 Rs → exit immediately

Why ÷ 3: Gives enough room for normal theta erosion and VIX fluctuation
but prevents holding a position that has fundamentally broken.
Maximum realistic loss = 18 Rs (full debit).
Stop at 6 Rs limits actual loss to 12 Rs (67% of max).

Scaling version: If you scaled during VIX spike, average your total
debit paid and apply stop on the combined average.


### 7.2 Profit Booking Rules

PARTIAL BOOKING (book 50% of position):
Trigger: Current BF ≥ Entry Debit × 1.5
Action: Close 50% of lots. Hold remaining 50%.
Why: Locks in profit while letting winners run further.

FULL BOOKING:
Trigger: Current BF ≥ Entry Debit × 2.0
Action: Close 100% of position.
Why: 2× debit = 100% return on risk. Exceptional result.
Take it. Re-enter fresh if conditions permit.

MIDDLE EXCEEDED BOOKING:
Trigger: Market price crosses the middle strike
Action: Book entire position IMMEDIATELY
Why: Maximum profit point has been passed. BF value will fall
from this point onward as market moves further.
Do not wait for market to "come back" to middle.

DTE-ADJUSTED BOOKING THRESHOLDS:
DTE ≥ 3 days → Book at 2.0× (standard rule)
DTE = 2 days → Book at 1.5× (time running out)
DTE = 1 day → Book at 1.3× (take what you can get)
DTE = 0 (expiry) → Let expire or exit if intrinsic < 0.5 Rs


### 7.3 Expiry Day Rules

On Thursday (expiry day):

Morning (9:15 - 11:00):
→ If ATM straddle > 50 Rs → position still has value → hold
→ If any leg is ITM → closely monitor, be ready to exit

Post 11:00:
→ Straddle collapses rapidly
→ If profitable → book by 1:00 PM (don't let it go to zero)
→ If marginally profitable → book by 2:00 PM
→ If at loss → decide: hold to expiry (zero more loss) or exit

Post 3:00 PM on expiry:
→ All remaining positions settle at intrinsic value
→ No action needed unless position is borderline ITM/OTM

Settlement:
BF value at expiry = max(0, spot−K) − 2×max(0, spot−K−gap) + max(0, spot−K−2×gap)
For put BF: max(0, K−spot) − 2×max(0, K−gap−spot) + max(0, K−2×gap−spot)
This is your final P&L if held to expiry.


### 7.4 Chain Concept Exit Rule

This is the most nuanced exit signal.

Step 1: From current market level, look at adjacent BF rates
in the option chain (from your exit tool).

Step 2: Ask "If market moves one more straddle-worth against me,
what does my BF rate become?"
→ Look at the rate at that hypothetical level from the chain.

Step 3: Decision tree:
If hypothetical BF > 40% of entry debit → HOLD (survivable)
If hypothetical BF 20-40% of entry debit → REVIEW (risky, watch)
If hypothetical BF < 20% of entry debit → EXIT NOW (chain says exit)

The chain will tell you before the position actually deteriorates.
This is forward-looking exit, not reactive exit.

Real example from course:
"Agar yahaan se 100 point aur neeche gayi to meri butterfly
kya ho jaayegi? Chain check karo — agar 3 Rs ho jaati hai to
abhi nikal jao, loss kha lo, fresh entry lo."


### 7.5 VIX Spike — Hold vs Scale Decision

When VIX spikes (rises > 1.0 in a session):
→ All BF values drop (M2M loss appears)
→ This is NOT an exit signal
→ This IS a scaling opportunity

Decision process during VIX spike:
Step 1: Check if fundamental structure is intact
(market has NOT moved more than 2× straddle)
Step 2: Check available margin
Step 3: Check scaled position's average Hold%
(make sure combined Hold% stays < 20% after scaling)
Step 4: Scale in at the cheaper BF rates

After event/VIX collapse:
→ ALL positions profit simultaneously
→ Premium erosion accelerates
→ Book in stages: 1/3 first, hold 2/3, book another 1/3 later

What NOT to do during VIX spike:
→ Do NOT panic-exit all positions
→ Do NOT reduce lots because of M2M loss
→ Do NOT stop watching the position — monitor closely
→ Do NOT scale without checking margin first


---

## PART 8 — SCALING RULES

### 8.1 Scale Check Formula

Scale% = Current BF Rate ÷ Current FLP × 100

This is the same as entry Value% but calculated NOW.

If Scale% ≤ 20% AND current BF < original entry debit:
→ Scaling is justified (you are adding at cheaper rate)
→ Average cost comes down → position improves

If Scale% > 20%:
→ Do not scale (not value here even if VIX is high)

If current BF > original entry debit:
→ Position already profitable → no need to scale
→ Book some lots instead


### 8.2 When to Scale

MANDATORY SCALING CONDITION:
All three must be true:
(a) Scale% ≤ 20% (value concept passes)
(b) Current BF < entry debit (scaling is at better price)
(c) Margin is available (do not over-leverage)

IDEAL SCALING CONDITION (scale aggressively):
All of mandatory PLUS:
(d) VIX is elevated (> 15) — means more premium to collect when it falls
(e) Overshoot ratio today > 0.8 — market has already moved, likely to stabilize
(f) VIX showing signs of topping (was higher earlier in session)

SCALE SIZE:
If Scale% ≤ 10% → add up to 2× original lots
If Scale% ≤ 20% → add 1× original lots
Never more than 2× original size in a single scaling event
Maximum total position size = 3× original lots per expiry


### 8.3 Averaging Rule

After scaling, compute:
Average debit = (Original lots × original debit + New lots × new BF)
/ (Original lots + New lots)

New Hold% check = Average debit ÷ Current FLP × 100
Must remain < 20% after scaling.

Scaling example:
Original: 2 buy lots, debit 18 Rs, FLP 103 Rs
VIX spikes, BF falls to 9 Rs, FLP = 80 Rs
Scale% = 9/80 × 100 = 11.3% → IDEAL, scale

Scale in: 2 more buy lots at 9 Rs
New average debit = (2×18 + 2×9) / 4 = 13.5 Rs
New Hold% = 13.5 / 80 × 100 = 16.9% → safe, hold

When VIX normalizes and FLP returns to 103 Rs:
New Hold% = 13.5 / 103 = 13.1% → very comfortable
BF rises back toward 18 Rs or higher
P&L on 4 lots vs original 2 lots → large profit


---

## PART 9 — BACK RATIO SPREAD

### 9.1 When to Use Back Ratio (NOT Butterfly)

Use back ratio ONLY when ALL of these are true:
(a) ATM Straddle < 80 Rs (premiums very compressed)
(b) Historical average straddle for this DTE/VIX > 1.8× current
(current straddle is unusually low)
(c) You expect VIX to expand OR a directional move to come
(d) You are willing to hold for potential move (not pure theta play)

If only (a) is true but not (b) → Wait for better straddle, don't force entry
Back ratio is used in MAXIMUM 10-20% of weekly opportunities.
80-90% of the time you should be in debit butterfly.


### 9.2 Back Ratio Structure

CALL BACK RATIO (expecting upward move or VIX expansion):
Sell 1 lot at K (middle — near ATM)
Buy 2 lots at K + gap (further OTM calls)

Result: Net CREDIT received
Max profit: Unlimited upside (net long 1 call after hedge)
Max loss : At K + gap at expiry (limited by structure)

PUT BACK RATIO (expecting downward move or VIX expansion):
Sell 1 lot at K (middle — near ATM)
Buy 2 lots at K − gap (further OTM puts)

Result: Net CREDIT received
Max profit: Large downside move
Max loss : At K − gap at expiry


### 9.3 Back Ratio Stop Loss

Stop loss = Credit received × 2
Example: Credit = 15 Rs → Stop at 30 Rs loss on the spread

Exit if spread widens to credit + stop loss
Time stop: Exit by Wednesday even if profitable (don't carry to expiry)
If VIX normalizes and straddle expands → exit and enter debit BF instead


---

## PART 10 — THE 8 SCENARIOS

### SCENARIO 1 — Market Up 100 pts

What happens:
Call BFs going toward ITM → first leg being tested
Put BFs moving OTM → value eroding but slowly
VIX typically falls when market rises → helps all BFs

Action:
Call side: Check chain. If BF at next 100 pt level still > 40% of entry → hold
Put side: Check Hold%. If < 20% → hold. OTM theta still working.
Do NOT exit call side prematurely just because market rose.
If call BF still > entry × 1.2 → consider partial booking.


### SCENARIO 1A — Market Up 200+ pts

What happens:
Call BF likely ITM or near ITM
VIX possibly at day's low
Put BFs cheaper now

Action:
Call side: Apply ITM rules. If market > middle strike → exit calls.
Put side: VIX low → put premiums cheaper → SCALE puts at better rates.
Scale put BFs to compensate for call BF loss.
The put scaling at cheaper rates recovers the call loss when market stabilizes.


### SCENARIO 2 — Market Down 100 pts

What happens:
Put BFs going toward ITM
VIX typically rises when market falls → all BF values drop (M2M loss)
Call BFs OTM but premium squeezed by VIX rise

Action:
Put side: Chain concept check. ITM rules apply.
Call side: Hold% check. VIX rise means FLP rose slightly → Hold% may be OK.
If call BFs are cheaper now → scale calls (VIX-spike opportunity).
M2M loss on puts is temporary if VIX spike is event-driven.
Wait for VIX to peak, then scale call side.


### SCENARIO 2A — Market Down 200+ pts

What happens:
VIX spike → all premiums elevated → all BF values drop
Call BFs very cheap (VIX elevated + deeper OTM)
Put BFs may be ITM → apply ITM rules

Action:
Put side: If middle exceeded → exit puts. Book whatever profit/loss remains.
Call side: AGGRESSIVELY SCALE calls at cheap rates.
Scale% will be very low (VIX elevated = high FLP = low value%)
This is the maximum opportunity scenario.
After VIX collapses: call BF values surge → large profit.
Patience is essential. M2M loss on puts + scaling calls looks bad briefly.
When VIX collapses (post-event or stabilization) → entire position profits.


### SCENARIO 3 — Market Stabilizing

What happens:
Theta extracting premium daily
Straddle shrinking
VIX stable or falling
All BF values gradually rising

Action:
OTM BFs: Book when they reach 80-90% of debit paid (zero value approaching,
free up margin, find better entry)
ATM BFs: Hold for more theta. These are the best performers now.
Book in stages: 1/3 at 1.5× debit, 1/3 at 2.0× debit, 1/3 at expiry
Look for fresh entries in next week or monthly if current week is near expiry.


### SCENARIO 3A — Market Continues Stable (Multiple Days)

Action:
Book all OTM positions immediately.
Hold ATM and near-ATM only.
As expiry approaches, move to next week's butterflies.
Daily homework: straddle decaying as expected?
If yes → normal operation
If no (straddle stable despite time) → event risk ahead?
Diversify: add Bank NIFTY butterflies, next week/monthly entries.
Number of trades must be HIGH to let probability work in your favor.


### SCENARIO 4 — Market Flat But VIX Rising

What happens:
Market not moving but VIX spiking (event ahead: RBI, FOMC, results)
Straddle not decaying (may be rising)
All BF values dropping (negative vega hurts)
M2M loss growing despite market not moving

Action:
DO NOT PANIC. This is the hardest scenario psychologically.
The position is NOT broken — VIX expansion is temporary.
Scale: Add BFs at the now-cheaper rates.
Your average cost drops.
Post-event VIX collapses → all positions profit simultaneously.
Sizing: Scale carefully — do not over-leverage before an event.
Stick to Scale% ≤ 20% rule even here.
Hold through the event if margin allows.


### SCENARIO 4A — VIX High + Market Moving

What happens:
Post-spike: VIX falling + market has moved (directional)
Best case for butterfly: VIX collapsing + market stabilizing

Action:
This is the payoff scenario.
Patiently hold all scaled positions.
Book in stages as VIX normalizes.
Do NOT exit all at once — let the full VIX collapse play out.
The further VIX fell, the more the BF values recovered.
Target: hold until straddle is back to pre-spike levels.


---

## PART 11 — EXPIRY SELECTION

### 11.1 Which Expiry to Trade

WEEKLY (current week):
Use when: Straddle is normal, DTE ≥ 3 days (Monday-Tuesday entry)
Tolerance: ±straddle move
Theta: Fastest decay — best for stable markets
Risk: Less time for recovery if market moves against

NEXT WEEK (one Thursday forward):
Use when: Current week DTE < 3, or want more buffer
Tolerance: ±300-400 pts (more time)
First two legs here, last leg in current weekly

MONTHLY (last Thursday of month):
Use when: VIX is elevated, want maximum buffer
Tolerance: ±500-600 pts
Best for scaling strategy (lots of time for VIX collapse)
First two legs here, last leg in nearest weekly

COMBINATION (multi-expiry):
Entry DTE 0-2 → skip current week, use next week + monthly
Entry DTE 3-4 → current week main, hedge with next weekly
Entry DTE > 5 → current week main legs


### 11.2 Expiry Day Rules by Entry Day

Monday entry → Current week expiry → 4 DTE → Maximum theta week
Tuesday entry → Current week expiry → 3 DTE → Good if conditions pass
Wednesday → Use next week expiry → current week too close
Thursday/Fri → Use next week expiry → current week expired/expiring


---

## PART 12 — DAILY HOMEWORK

### 12.1 Morning Routine (Pre-Market: 8:30-9:00 AM)
Check GIFT NIFTY / SGX NIFTY → gap up/down expectation
Check Global markets overnight → risk-on / risk-off
Check economic events today:
→ RBI policy? FOMC? GDP? Results?
→ If major event → DO NOT enter new positions
→ Existing positions: reduce size or hedge before event
Record expected straddle based on VIX level:
→ Formula: spot × VIX% × sqrt(DTE/252) × 0.4 ≈ straddle
Compare expected vs actual straddle at 9:15 → is premium normal or inflated?

### 12.2 At Market Open (9:15-9:30 AM)
Note ATM straddle at open
Note VIX at open
Calculate overshoot ratio if market gapped significantly
Check OI data in option chain → where is OI building?
Decide entry or wait based on opening conditions
If straddle is abnormally high at open → wait 10-15 mins for settlement
If straddle is normal → proceed with entry scoring

### 12.3 Weekly Straddle Calibration Database

Build and maintain this table over time:

VIX Level	DTE	Expected Straddle	Normal BF Rate (100 gap)
< 12	4	80-100	5-8 Rs
12-14	4	120-150	10-14 Rs
14-17	4	150-200	14-20 Rs
17-20	4	200-280	20-30 Rs
20-25	4	280-380	30-45 Rs

25 | 4 | 380+ | 45+ Rs

Use this to detect when current premiums are abnormally low
(back ratio signal) or abnormally high (aggressive BF entry).


---

## PART 13 — RISK MANAGEMENT

### 13.1 Position Limits

Daily loss limit : -2% of capital → stop all trading for the day
Weekly loss limit : -5% of capital → review strategy before continuing
Monthly loss limit : -10% of capital → pause and analyze
Single trade max : -0.5% of capital (= max debit per position)

These limits are NON-NEGOTIABLE.
When hit, no new positions, no scaling, no revenge trading.


### 13.2 Greeks Exposure Limits

Net Delta : ±50 (butterfly structure keeps this near zero naturally)
Net Gamma : ±20 (managed by value concept entry rule)
Net Theta : > +500 Rs/day (we want positive theta — always)
Net Vega : -500 to -3000 Rs/VIX point (we are short vega)
→ Track this carefully
→ When net vega breach → no new BF entries until within limit

Vega interpretation:
Net Vega = -1000 means:
Every 1 point VIX rises → portfolio loses ₹1,000
Every 1 point VIX falls → portfolio gains ₹1,000
This is why VIX is the ONLY greek you actively manage.


### 13.3 Correlation Risk

If holding:
Call BF at 22800 (short vega)
Put BF at 23000 (short vega)
Next week CE BF at 22700 (short vega)
Monthly CE BF at 22500 (short vega)

Net portfolio vega = sum of all positions' vega
This can be 3-4× a single position's vega.
When VIX spikes, ALL positions lose simultaneously.
Monitor PORTFOLIO vega, not individual position vega.

Rule: Net portfolio vega must not exceed -5000 Rs/VIX point
(adjust based on your capital)


### 13.4 Margin Rules

Before any new position:
Available margin must be > new position margin × 2
(keep 50% buffer at all times)

Before scaling:
Available margin must be > scale position margin × 3
(extra buffer needed for VIX spike scenarios)

Never use > 70% of total capital as margin at once.
The 30% buffer is for scaling opportunities during VIX spikes.
If you are 100% deployed → you cannot scale → you miss the
best opportunity. Keep powder dry.


---

## PART 14 — PSYCHOLOGY RULES

### 14.1 Patience Rules

Rule: Do NOT adjust for moves < 50 pts.
Market noise is ±30-50 pts constantly.
Adjusting to noise = death by a thousand cuts.
Wait for at minimum 50-100 pts adverse before any action.

Rule: M2M loss is NOT real loss until you exit.
A ₹10,000 M2M loss that becomes ₹5,000 profit next day
= no loss, just volatility. Do not react to M2M.

Rule: One losing week does not break the strategy.
The strategy wins on PROBABILITY across many trades.
Single trade outcomes mean nothing.
Monthly outcomes matter more.
Quarterly outcomes matter most.


### 14.2 Discipline Rules

Rule: NEVER break the 25% Hold% exit rule.
No matter how convinced you are of recovery.
The rule exists because the math says exit.

Rule: NEVER enter when Value% > 25%.
Not even once. The first exception creates a habit.
The habit creates a blowup.

Rule: NEVER scale when Margin buffer < 30%.
Scaling with no margin = forced exit at worst prices = maximum loss.

Rule: Review every trade, win or lose.
Every week, answer: What did I do right? What rule did I break?
The ones you broke are the lessons. Write them down.


### 14.3 Common Mistakes to Avoid

Mistake 1: Holding OTM positions hoping they recover
Correct: When Hold% > 25%, exit immediately.
Re-enter at ATM-closer strike. Don't hold dead trades.

Mistake 2: Not scaling during VIX spikes out of fear
Correct: VIX spike = premium gift. The scaling IS the strategy.
If you don't scale during VIX spikes, you miss 50% of returns.

Mistake 3: Exiting profitable positions too early
Correct: Use the tiered booking rule (1.5×, 2.0×).
Let winners run to their targets.

Mistake 4: Entering without checking OI concentration
Correct: Always check where OI is building before placing middle leg.
ATM is not always the best middle.

Mistake 5: Ignoring VIX direction
Correct: VIX level is secondary. VIX direction is primary.
VIX falling = enter. VIX rising sharply = wait.

Mistake 6: Trading expiry day without strict rules
Correct: On expiry day, tight stops, book by 2 PM.
Do not hold profitable positions to zero hoping for more.


---

## PART 15 — COMPLETE ENTRY CHECKLIST

Before EVERY trade, run through this checklist.
ALL items must be checked. No shortcuts.

PRE-TRADE CALIBRATION
□ ATM Straddle noted: _________ Rs
□ Gap calculated: _________ pts (straddle ÷ 2 → nearest 50)
□ Combination selected: 1:2 / 2:3 / 3:4
□ Adverse tolerance noted: _________ pts
□ VIX current: _________ (level + direction)
□ Overshoot ratio: _________ (today's move ÷ straddle)
□ OI concentration checked: Middle at _________ (OI peak strike)
□ PCR noted: _________ (neutral/bullish/bearish)
□ Events today: None / Event (if event → DO NOT ENTER)

VALUE CONCEPT CHECK
□ FLP of chosen strike: _________ Rs (must be > ₹8)
□ BF Rate calculated: _________ Rs (must be > ₹0.5)
□ Value%: _________% (must be < 25%)
□ Grade: Ideal / Good / Pass / Borderline / FAIL

CHAIN CONCEPT CHECK
□ Scenario UP (straddle up): BF = _________ Rs Survives? Y/N
□ Scenario FLAT (theta): BF = _________ Rs Survives? Y/N
□ Scenario DOWN (straddle dn):BF = _________ Rs Survives? Y/N
□ Scenarios survived: ___ / 3 (need ≥ 2)

ENTRY SCORE
□ Value% score: ____ / 60
□ VIX direction: ____ / 40
□ Overshoot ratio: ____ / 40
□ Double lock: ____ / 30
□ PCR: ____ / 10
□ Chain concept: ____ / 20
□ TOTAL SCORE: ____ / 200 (minimum 60 to enter)

SIZING
□ Lots based on Value% tier: ____ buy lots
□ Margin required: _________ Rs
□ Margin available: _________ Rs (must be > required × 2)
□ Post-trade margin buffer: _________ Rs (must be > 30% of total capital)

TRADE RECORD
□ Entry date: ___________
□ Strike: ___________ (K)
□ Gap: ___________ pts
□ Type: CE / PE
□ Combination: 1:2 / 2:3 / 3:4
□ Expiry: ___________
□ Entry debit: _________ Rs
□ Entry FLP: _________ Rs
□ Entry Value%: _________ %
□ Entry VIX: _________
□ Entry Straddle: _________ Rs
□ Stop loss level: _________ Rs (entry debit ÷ 3)
□ Partial booking level: _________ Rs (entry debit × 1.5)
□ Full booking level: _________ Rs (entry debit × 2.0)


---

## PART 16 — COMPLETE EXIT CHECKLIST

Run this checklist daily for every open position.

DAILY HOLD CHECK
□ Current BF Rate: _________ Rs
□ Current FLP: _________ Rs
□ Hold% = Entry Debit ÷ FLP: _________ %

□ Hold% ≤ 15% → HOLD (no action)
□ Hold% 15-20% → HOLD (monitor closely, check chain)
□ Hold% 20-25% → REVIEW (run chain check now)
□ Hold% > 25% → EXIT IMMEDIATELY

STOP LOSS CHECK
□ Stop level = Entry Debit ÷ 3 = _________ Rs
□ Current BF < Stop level? YES → EXIT / NO → Continue

PROFIT BOOKING CHECK
□ Partial booking level = Entry × 1.5 = _________ Rs
□ Current BF ≥ partial level? YES → Book 50% / NO → Continue
□ Full booking level = Entry × 2.0 = _________ Rs
□ Current BF ≥ full level? YES → Book 100% / NO → Continue

MIDDLE EXCEEDED CHECK
□ Market currently at: _________
□ Middle strike at: _________
□ Market > Middle (CE) or Market < Middle (PE)?
YES → EXIT FULL POSITION / NO → Continue

ITM CHECK (if applicable)
□ Is position ITM? YES / NO
□ If YES: Chain concept check for further adverse move
□ If BF at next adverse level < 40% of entry → EXIT
□ Hold% for ITM position threshold is 15% (tighter)

CHAIN CONCEPT CHECK (when Hold% 20-25%)
□ Adjacent strike BF (up move): _________ Rs
□ Adjacent strike BF (down move): _________ Rs
□ Both > 40% of entry debit?
YES → Hold one more session
NO → EXIT this leg

COMBINED PORTFOLIO CHECK (when running pairs)
□ Combined Hold% = Total Debit ÷ Total FLP = _________ %
□ Combined < 20% → hold both legs
□ Combined 20-25% → exit breached leg, hold compensating leg
□ Combined > 25% → exit both legs

DTE CHECK
□ DTE remaining: _____ days
□ DTE = 0 (expiry today): settle or exit by 2 PM
□ DTE = 1: book at 1.3× target, tighter stop
□ DTE = 2: book at 1.5× target
□ DTE ≥ 3: standard rules apply


---

## PART 17 — PERFORMANCE TRACKING

### 17.1 Per-Trade Record (Required Fields)

Trade ID | Date | Strike | Type | Gap | Combo | Expiry |
Entry Debit | Entry FLP | Value% | Score | VIX at Entry |
Straddle at Entry | Chain Survived |
Exit Date | Exit BF | Exit Reason | Hold% at Exit |
Days Held | P&L Points | P&L Rs | Lots


### 17.2 Weekly Review Questions
How many trades did I enter this week?
What was the average Value% at entry?
Did I follow every checklist item? If no, which did I skip?
Which exit reason was most common?
What was my VIX reading when I entered? Was it trending?
Did I scale during any VIX spike? If yes, what was the outcome?
What was my P&L this week vs expected (straddle × lots)?
Was there a trade I wish I had NOT taken? Why?
Was there a trade I wish I HAD taken but didn't? Why?
What one thing will I do differently next week?

### 17.3 Key Performance Benchmarks

Win Rate Target : > 60% of trades profitable
Avg P&L per trade : > 0.5× average entry debit
Profit Factor : > 1.5 (total wins / total losses)
Max Drawdown (monthly) : < 10% of capital
Sharpe Ratio (annual) : > 1.5
Avg Days Held : 2-4 days (weekly cycle)

If any metric consistently misses target for 4+ weeks:
→ Review entry rules (Value% threshold too loose?)
→ Review VIX timing (entering during spikes instead of after?)
→ Review exit rules (exiting too early / too late?)
→ Review sizing (too large for current capital?)


---

## SUMMARY — THE TEN COMMANDMENTS
Value% > 25% → NEVER ENTER. No exception, no discussion.
Hold% > 25% → EXIT IMMEDIATELY. Do not wait for recovery.
VIX direction matters more than VIX level.
Falling VIX = enter. Rising VIX = wait.
Middle strike = OI concentration peak (not blindly ATM).
M2M loss from VIX spike = scaling opportunity.
Do NOT exit because of M2M loss. Scale if Value% permits.
When market stabilizes → book OTM legs, hold ATM legs.
Keep minimum 30% margin buffer for scaling opportunities.
Run BOTH call and put butterflies. Evaluate combined Hold%.
One side can compensate the other.
After a big move (overshoot ratio > 1.0) → PRIME entry time.
The market has done the work. Theta now does the rest.
Number of trades must be HIGH.
Probability only works over many trades, not individual ones.
Small consistent profits accumulate into large monthly returns.

---

*Document Version: 3.0*
*Last Updated: Based on live market evidence from 22856 session*
*Next Review: After 50 completed trades or 3 months, whichever comes first*
*Owner: Strategy must be reviewed and updated as market conditions change.*

I want in .md file

Thought for 7s

Here is the complete file. Copy everything between the lines and save as butterfly_trade_rules.md.

# NIFTY 50 — Multi-Expiry 1:2:1 Butterfly
# Complete Trade Entry & Exit Rules
**Version 3.0 — Evidence-Based**

---

## PART 1 — PHILOSOPHY

### 1.1 What We Are

| Item | Value |
|---|---|
| Profit source | Theta decay + VIX contraction |
| Position type | Non-directional (delta-gamma neutral by structure) |
| Greek we manage | **VEGA only** (VIX = single active risk) |
| We are NOT | Naked sellers, directional traders, scalpers |

**Core equation**

PROFIT = Premium Erosion (Theta) + VIX Contraction (Vega gain)
LOSS = VIX Expansion (Vega loss) + Adverse move beyond tolerance


> M2M loss from VIX spike is **NOT** real loss. It is a **scaling opportunity**.

---

### 1.2 Structure

**Call Butterfly**

Buy 1 lot at K ← First Leg (FLP)
Sell 2 lots at K + gap ← Middle Leg (max profit here)
Buy 1 lot at K + 2×gap ← Last Leg (hedge)


**Put Butterfly**

Buy 1 lot at K ← First Leg (FLP)
Sell 2 lots at K − gap ← Middle Leg
Buy 1 lot at K − 2×gap ← Last Leg


- Net position: **Debit paid**
- Max profit: At expiry when **spot = middle leg strike**
- Max loss: **Entry debit** (limited, defined risk)

---

### 1.3 Multi-Expiry Version

| Leg | Expiry | Why |
|---|---|---|
| First Leg (buy) | Next Week or Monthly | More premium, wider range |
| Middle Leg (sell 2×) | Same as First Leg | Matches the buy |
| Last Leg (buy hedge) | Current Weekly | Cheap 1–3 Rs hedge |

---

### 1.4 Valid Combinations

| Combination | Ratio | Exit Shows | Margin/Lot | When |
|---|---|---|---|---|
| 1:2:1 | 1 buy : 2 sell | 2.0 | ~₹45,000 | Straddle ≤ 160 |
| 2:3:2 | 2 buy : 3 sell | 1.5 | ~₹35,000 | Straddle 160–280 |
| 3:4:3 | 3 buy : 4 sell | 1.33 | ~₹28,000 | Straddle > 280 |

> Last hedge leg is personal judgment. Not included in value% calculation.

---

## PART 2 — PRE-ENTRY CALIBRATION

### Step 1 — Read ATM Straddle

ATM Straddle = ATM Call LTP + ATM Put LTP


This single number drives **everything** — gap, combination, tolerance, and
whether to use debit BF or back ratio.

---

### Step 2 — Recommended Gap

Gap = ATM Straddle ÷ 2, rounded to nearest 50


| Straddle | Gap |
|---|---|
| ≤ 140 | 50 pts |
| 141–240 | 100 pts |
| 241–380 | 150 pts |
| > 380 | 200 pts |

**Example:** Straddle = 167.8 → 167.8 ÷ 2 = 83.9 → nearest 50 = **100 pts**

---

### Step 3 — Select Combination

| Straddle | Combination | Exit Ratio |
|---|---|---|
| ≤ 160 | 1:2 | 2.0 |
| 161–280 | 2:3 | 1.5 |
| > 280 | 3:4 | 1.33 |

---

### Step 4 — Adverse Move Tolerance

| Entry Timing | Tolerable Adverse Move |
|---|---|
| Expiry day | ±50 pts |
| Day before expiry | ±100 pts |
| Monday (weekly) | ±200–250 pts |
| Next week entry | ±300–400 pts |
| Monthly entry | ±500–600 pts |

> Tolerance ≈ ATM Straddle value. The premium tells you the market's own expectation.

---

### Step 5 — VIX Check

**VIX Level**

| VIX | Regime | Action |
|---|---|---|
| < 12 | Very low | Check back ratio instead |
| 12–14 | Normal | Ideal debit BF zone |
| 14–17 | Moderate | Good entry |
| 17–20 | Elevated | Excellent entry (more premium) |
| > 20 | High | Scale aggressively on dips |

**VIX Direction (MORE important than level)**

| VIX Direction Today | Score | Action |
|---|---|---|
| Falling > 1.0 | +40 pts | STRONG ENTRY — BF values rising |
| Falling 0.5–1.0 | +25 pts | Good entry |
| Falling 0–0.5 | +10 pts | Normal entry |
| Rising 0–1.0 | 0 pts | Caution |
| Rising > 1.0 | −20 pts | Wait for peak, enter on turn |

> **Best setup:** VIX > 17 AND falling today → Maximum priority entry

---

### Step 6 — Straddle Overshoot Signal *(New Rule)*

Overshoot Ratio = Today's market move ÷ ATM Straddle


| Ratio | Signal | Score |
|---|---|---|
| > 1.2 | PRIME ENTRY — market overshot weekly range | +40 pts |
| 0.9–1.2 | Good entry | +20 pts |
| 0.5–0.9 | Normal | +5 pts |
| < 0.5 | Market barely moved | −5 pts |

**Live data example:**
> Market moved 235 pts, straddle = 167.8 → Ratio = 1.40 → **PRIME ENTRY**

**Why:** A market that overshoots its weekly expected range in one session
almost always stabilizes for remaining days → best theta extraction setup.

---

### Step 7 — OI-Based Middle Strike Selection *(New Rule)*

**Do NOT always use ATM as middle strike.**

For CALL butterfly:
→ Find strike ABOVE ATM with highest positive OI change
→ Place MIDDLE leg there (market maker resistance / pin target)

For PUT butterfly:
→ Find strike BELOW ATM with highest positive OI change
→ Place MIDDLE leg there (market maker support / pin target)


**Live data example:**

| Strike | Call OI Change |
|---|---|
| 22850 (ATM) | +5,720,850 |
| **22900** | **+7,105,725** ← highest near ATM |
| 23000 | +9,207,900 ← resistance anchor |

→ Place **call BF middle at 22900**, not ATM 22850

**Why:** Market makers pin prices where they have maximum short interest.
Place your peak profit exactly where they want price to go.

---

### Step 8 — PCR Directional Bias

| PCR | OI Structure | Adjustment |
|---|---|---|
| < 0.7 | Bullish (more calls) | Shift call BF middles 50–100 pts higher |
| 0.7–1.2 | Neutral | Symmetric placement |
| 1.2–1.8 | Bearish (more puts) | Put BF middles 50 pts lower |
| > 1.8 | Extreme bearish | Reduce calls, increase puts |

---

### Step 9 — Double Middle Lock *(New Rule)*

Double Middle Lock = Call middle strike = Put middle strike


When both butterflies share the same middle strike:
- Call BF peaks at that level
- Put BF peaks at the **same** level
- Both profit **simultaneously** if market stabilizes there

> Add **+30 points** to entry score when double lock is achieved.
>
> **This is the highest-priority setup. Always try to create it.**

---

## PART 3 — VALUE CONCEPT (ENTRY GATE)

### 3.1 Core Formula

Value% = BF Rate ÷ FLP × 100

BF Rate = Leg1 LTP − (2 × Middle LTP) + Last Leg LTP
FLP = First Leg Premium (LTP of buying leg at strike K)


When Value% ≤ 20% → delta and gamma are automatically near zero.
**You do not need to manually hedge. The structure does it.**

---

### 3.2 Entry Thresholds

| Value% | Grade | Action | Lot Size |
|---|---|---|---|
| ≤ 10% | **IDEAL** | Enter immediately | 3× base lots |
| 10.1–15% | **GOOD** | Enter | 2× base lots |
| 15.1–20% | **PASS** | Enter | 1× base lots |
| 20.1–25% | **BORDERLINE** | Enter ONLY if chain ≥ 2/3 AND VIX falling AND overshoot > 1.0 | 0.5× base lots |
| > 25% | **FAIL** | **NEVER ENTER. Non-negotiable.** | 0 |

> **The FAIL rule has no exceptions.** Not for "good market conditions."
> Not for "strong VIX signal." Not for "this one time." Never.

---

### 3.3 Data Quality Gate (Check Before Value%)

| Check | Minimum | If Fails |
|---|---|---|
| FLP | > ₹8 | Skip — too deep OTM or stale data |
| BF Rate | > ₹0.5 | Skip — spread too tight or anomaly |
| ATM Straddle | > ₹50 | Consider back ratio instead |

---

### 3.4 Entry Scoring System

**Minimum score to enter: 60 points**

| Factor | Condition | Points |
|---|---|---|
| **Value%** | ≤ 10% | 60 |
| | ≤ 15% | 45 |
| | ≤ 20% | 25 |
| | ≤ 25% | 5 |
| | > 25% | REJECT |
| **VIX Direction** | Falling > 1.0 | +40 |
| | Falling 0.5–1.0 | +25 |
| | Falling 0–0.5 | +10 |
| | Rising 0–1.0 | 0 |
| | Rising > 1.0 | −20 |
| **Overshoot Ratio** | > 1.2 | +40 |
| | 0.9–1.2 | +20 |
| | 0.5–0.9 | +5 |
| | < 0.5 | −10 |
| **Double Lock** | Both middles same strike | +30 |
| | Different strikes | 0 |
| **PCR** | 0.7–1.3 (neutral) | +10 |
| | < 0.5 or > 1.8 | −10 |
| **Chain Concept** | 3/3 scenarios survive | +20 |
| | 2/3 survive | +10 |
| | 1/3 survive | 0 |
| | 0/3 survive | −20 |

**Lot sizing by score:**

| Score | Lots |
|---|---|
| ≥ 120 | 3× base |
| 90–119 | 2× base |
| 60–89 | 1× base |
| < 60 | Skip |

---

## PART 4 — CHAIN CONCEPT

### 4.1 Three-Scenario Validation

Before entering, simulate your butterfly value in 3 scenarios:

| Scenario | Market Level | Survival Condition |
|---|---|---|
| Adverse Up | Spot + straddle | BF ≥ 40% of entry debit |
| Flat | Spot unchanged | BF ≥ entry × 1.1 (theta works) |
| Adverse Down | Spot − straddle | BF ≥ 40% of entry debit |

> **Need minimum 2 of 3 to survive before entering.**

**Example — 22800 CE BF, entry debit = 18 Rs, straddle = 167:**

Threshold = 18 × 0.40 = 7.2 Rs

Scenario Up (22967): Check 23000 CE BF → 11 Rs > 7.2 → ✓ Survives
Scenario Flat (22800): BF ≈ 18×1.1 = 19.8 → > 7.2 → ✓ Survives
Scenario Down (22633): Check 22650 CE BF → 5 Rs < 7.2 → ✗ Fails

Result: 2/3 survive → Entry permitted


---

### 4.2 Chain Concept During Trade Management

> Always ask: **"If market moves X more pts from HERE, what does my BF become?"**

| Position State | Primary Tool | Action |
|---|---|---|
| Going ITM | Chain Concept (heavily) | Check adjacent BF rates |
| Going OTM | Value Concept (Hold%) | Check old debit ÷ new FLP |
| Stable | Both equally | Collect theta, monitor |

---

## PART 5 — HOLD RULES

### 5.1 Hold% Formula

Hold% = Entry Debit ÷ Current FLP × 100

Entry Debit = what you paid (fixed)
Current FLP = live price of first leg option NOW


**What it means:**

| Situation | Hold% Movement | Implication |
|---|---|---|
| FLP high (entry day) | Low | Good — structure intact |
| FLP falling (going OTM) | Rising | Warning — structure weakening |
| Hold% > 20% | High | Position becoming directional |
| Hold% > 25% | Very high | **Exit now** |

---

### 5.2 Hold% Decision Table

| Hold% | Grade | Action |
|---|---|---|
| ≤ 15% | Safe | **HOLD** — theta working, no action needed |
| 15–20% | Watch | **HOLD** — monitor closely, check chain |
| 20–25% | Review | **REVIEW** — run chain check now. Hold only if chain favorable AND VIX falling. |
| > 25% | Breach | **EXIT IMMEDIATELY** — non-negotiable |

---

### 5.3 Combined Portfolio Hold Rule *(New Rule)*

When running call + put butterfly pair together:

Combined Hold% = Total Entry Debit ÷ Total Current FLP × 100


| Combined Hold% | Action |
|---|---|
| < 20% | Hold both (one side compensating the other) |
| 20–25% | Exit breached leg, keep compensating leg |
| > 25% | Exit both |

**Example:**

Call BF: Hold% = 28% (bad), debit = 18, FLP = 64
Put BF: Hold% = 8% (good), debit = 16, FLP = 195

Combined = (18+16) / (64+195) × 100 = 34/259 × 100 = 13.1%

→ Combined is fine → Hold both
→ Do NOT exit call leg just because its individual Hold% is 28%


---

### 5.4 ITM Position Rules

| ITM Depth | Hold% Threshold | Action |
|---|---|---|
| Slightly ITM (0–50 pts) | Standard 25% | Normal hold/exit rules |
| Deeply ITM (50–100 pts) | Tighter 15% | Exit if Hold% > 15% |
| Very deeply ITM (> 100 pts) | Chain concept drives exit | Exit if chain shows BF → 0 on further move |

**ITM chain check:**

If BF at next adverse level < 40% of entry debit → EXIT NOW
One small loss now prevents a full debit loss later.


---

## PART 6 — EXIT RULES

### 6.1 Stop Loss

Exit when: Current BF < Entry Debit ÷ 3

Example: Entry debit = 18 Rs
Stop trigger = 18 ÷ 3 = 6 Rs
Exit when BF = 6 Rs


> Maximum actual loss = 12 Rs (67% of max possible).
> Without stop loss, loss = full 18 Rs debit.

---

### 6.2 Profit Booking

| Trigger | Action | Why |
|---|---|---|
| BF ≥ Entry × 1.5 | Book 50% of lots | Lock profit, let rest run |
| BF ≥ Entry × 2.0 | Book 100% | 100% return on risk — take it |
| Market crosses middle strike | Book 100% immediately | Max profit zone passed |

**DTE-adjusted booking targets:**

| DTE Remaining | Partial Booking | Full Booking |
|---|---|---|
| ≥ 3 days | Entry × 1.5 | Entry × 2.0 |
| 2 days | Entry × 1.3 | Entry × 1.7 |
| 1 day | Entry × 1.2 | Entry × 1.5 |
| 0 (expiry) | Any profit > 0 by 2 PM | — |

---

### 6.3 Chain Concept Exit

**When Hold% is 20–25%, run this chain exit check:**

Step 1: From current market level, look at adjacent BF rates
in the exit tool

Step 2: Ask "If market moves one more straddle-worth against me,
what does my BF rate become?"

Step 3:
Adjacent BF > 40% of entry debit → HOLD one more session
Adjacent BF 20–40% of entry debit → REVIEW (prepare to exit)
Adjacent BF < 20% of entry debit → EXIT NOW (chain says no hope)


---

### 6.4 Expiry Day Protocol

| Time | Action |
|---|---|
| 9:15–10:00 AM | Check ITM/OTM status. If profitable — hold. |
| 10:00–11:00 AM | Straddle decaying fast. Book partial if > 80% of max profit. |
| 11:00 AM–1:00 PM | If profitable → book 50–75% |
| 1:00–2:00 PM | Book all profitable positions |
| 2:00–3:00 PM | Book any remaining positions regardless |
| 3:15 PM | Settlement. Zero action needed. |

> **Never hold profitable positions to expiry hoping for more.**
> A 90% profit taken at 1 PM is better than 0% at 3:30 PM.

---

### 6.5 VIX Spike — Hold or Scale Decision

**When VIX rises > 1.0 in a session (all BF values drop):**

This is NOT an exit signal.
This IS a scaling opportunity.


**Decision process:**

Step 1: Is market move within 2× straddle?
YES → structure intact → proceed to step 2
NO → too extreme → reduce, do not scale

Step 2: Check available margin
Need > 3× new position margin → safe to scale

Step 3: Check Scale% (current BF ÷ current FLP × 100)
Scale% ≤ 20% → scale at this rate
Scale% > 20% → wait for more VIX expansion

Step 4: Scale in → average cost falls
Wait for VIX to collapse → ALL positions profit simultaneously


**What NOT to do:**

- ❌ Do NOT exit because of M2M loss
- ❌ Do NOT stop monitoring
- ❌ Do NOT scale without margin check
- ❌ Do NOT panic

---

## PART 7 — SCALING RULES

### 7.1 Scale% Formula

Scale% = Current BF Rate ÷ Current FLP × 100


**Entry conditions for scaling:**

| Condition | Must Be True |
|---|---|
| Scale% | ≤ 20% |
| Current BF vs entry | Current BF < entry debit (adding at cheaper price) |
| Margin available | > 3× new position margin |

**Ideal scaling conditions (scale aggressively when ALL true):**

✓ Scale% ≤ 10%
✓ VIX > 15 and elevated
✓ Overshoot ratio > 0.8 today
✓ VIX showing signs of topping (was higher earlier)
✓ Margin buffer > 40%


---

### 7.2 Scale Size Limits

| Scale% | Add | Total Max |
|---|---|---|
| ≤ 10% | Up to 2× original lots | 3× original |
| ≤ 20% | 1× original lots | 2× original |
| > 20% | 0 — wait | — |

> Never exceed 3× original lot size per expiry from a single entry.

---

### 7.3 Post-Scale Average Cost Check

New Average Debit = (Original lots × original debit + New lots × new BF)
/ (Original lots + New lots)

New Hold% = New Average Debit ÷ Current FLP × 100
Must remain < 20% after scaling.


**Scaling example:**

Original: 2 buy lots, debit 18 Rs, FLP 103 Rs, Hold% = 17.5%

VIX spikes, BF drops to 9 Rs, FLP rises to 130 Rs:
Scale% = 9/130 = 6.9% → IDEAL
Scale in: 2 more lots at 9 Rs

New avg debit = (2×18 + 2×9) / 4 = 13.5 Rs
New Hold% = 13.5 / 130 × 100 = 10.4% → safe ✓

When VIX normalizes, FLP returns to 103 Rs:
Hold% = 13.5 / 103 = 13.1% → comfortable ✓
BF rises to 18+ Rs
P&L on 4 lots vs original 2 lots → 2× profit


---

## PART 8 — BACK RATIO SPREAD

### 8.1 When to Use Back Ratio

**Use ONLY when ALL of these are true:**

✓ ATM Straddle < 80 Rs (premiums crushed)
✓ Current straddle < 55% of historical average for this DTE/VIX
✓ Expecting VIX expansion OR directional move
✓ You are NOT expecting range-bound theta play


> Back ratio = 10–20% of trades maximum.
> Debit butterfly = 80–90% of trades.

---

### 8.2 Back Ratio Structure

**Call Back Ratio** (expecting up move or VIX expansion):

Sell 1 lot at K (near ATM)
Buy 2 lots at K + gap (further OTM calls)
→ Net CREDIT received
→ Max profit: large up move
→ Max loss: at K + gap at expiry


**Put Back Ratio** (expecting down move or VIX expansion):

Sell 1 lot at K (near ATM)
Buy 2 lots at K − gap (further OTM puts)
→ Net CREDIT received
→ Max profit: large down move


**Stop loss:** Credit received × 2
**Time stop:** Exit by Wednesday regardless

---

## PART 9 — THE 8 SCENARIOS

### Case 1 — Market Up 100 pts

| What happens | Action |
|---|---|
| Call BFs approaching ITM | Chain check: if next 100 pt level still > 40% → hold |
| Put BFs moving OTM | Hold% check — if < 20% → hold, theta still working |
| VIX typically falls | Helps all BFs |

---

### Case 1A — Market Up 200+ pts

| What happens | Action |
|---|---|
| Call BFs ITM or near ITM | Apply ITM rules. If market > middle strike → **exit calls** |
| VIX at session low | Put premiums cheaper → **scale put BFs** |
| Put BFs OTM | Scale puts at better rates to compensate call loss |

---

### Case 2 — Market Down 100 pts

| What happens | Action |
|---|---|
| Put BFs toward ITM | Chain concept check, ITM rules apply |
| VIX rises → BF values drop | **Not a reason to exit** — M2M loss is temporary |
| Call BFs OTM | Hold% check. VIX rise = FLP rose → hold% may be ok |
| Call BFs cheaper | **Scale calls** — this is the VIX spike opportunity |

---

### Case 2A — Market Down 200+ pts

| What happens | Action |
|---|---|
| Put BFs may be ITM | Book puts if middle exceeded |
| VIX maximum spike | Call BFs **very cheap** → scale aggressively |
| All BF values dropping | This is the MAXIMUM scaling opportunity |

> After VIX collapses (post-event): all scaled call positions profit simultaneously.
> **Patience here creates the largest single-week P&L.**

---

### Case 3 — Market Stabilizing

| What happens | Action |
|---|---|
| Theta extracting daily | OTM BFs: book when near zero value |
| Straddle shrinking | ATM BFs: hold for more theta |
| All BF values rising | Book in stages: 1/3 at 1.5×, 1/3 at 2.0×, hold 1/3 |

---

### Case 3A — Market Continues Stable

→ Book all OTM positions immediately (free up margin)
→ Hold ATM and near-ATM only
→ Find fresh entries in next week or monthly
→ Diversify: Bank NIFTY, different expiries
→ Number of trades MUST be high


---

### Case 4 — Market Flat But VIX Rising

> Hardest scenario psychologically. Position intact — VIX expansion is temporary.

| What happens | Action |
|---|---|
| Straddle not decaying | M2M loss growing — **do not panic** |
| VIX spiking (event ahead) | Scale in at cheaper rates |
| BF values dropping | Average cost falls |

> Post-event VIX collapses → ALL positions profit simultaneously.

---

### Case 4A — VIX High + Market Moving

This is the PAYOFF scenario.

→ Patiently hold all scaled positions
→ Book in stages as VIX normalizes
→ Do NOT exit all at once — let full VIX collapse play out
→ Target: hold until straddle is back to pre-spike levels
→ The further VIX fell, the more BF values recovered


---

## PART 10 — EXPIRY SELECTION

### Which Expiry to Trade

| Entry Day | DTE | Use |
|---|---|---|
| Monday | 4 | Current week weekly |
| Tuesday | 3 | Current week weekly |
| Wednesday | 2 | **Next week** (current too close) |
| Thursday/Friday | 1/0 | Next week expiry |

### Multi-Expiry Configuration

| DTE at Entry | Leg 1+2 Expiry | Hedge Leg Expiry |
|---|---|---|
| 4–5 days | Current weekly | Current weekly (different strike) |
| 2–3 days | Next week | Current weekly |
| 8–12 days | Next week | Current weekly |
| > 15 days | Monthly | Current weekly |

---

## PART 11 — DAILY HOMEWORK

### Pre-Market (8:30–9:00 AM)

□ Check GIFT NIFTY / SGX NIFTY gap
□ Global markets overnight (risk-on / risk-off)
□ Economic events today? (RBI, FOMC, results)
→ Major event? → DO NOT enter new positions
□ Calculate expected straddle:
Formula: Spot × VIX% × sqrt(DTE/252) × 0.4 ≈ Straddle
□ Compare expected vs yesterday's actual → premium normal or inflated?


### At Market Open (9:15–9:30 AM)

□ Note ATM straddle at 9:15
□ Note VIX at open
□ Calculate overshoot ratio (if gap/move significant)
□ Check OI data → where is OI building?
□ Entry or wait?
□ If straddle abnormally high → wait 10–15 mins to settle


### Weekly Straddle Database

> Maintain this table. Update after every expiry.

| VIX | DTE | Expected Straddle | Normal BF (100 gap) |
|---|---|---|---|
| < 12 | 4 | 80–100 | 5–8 Rs |
| 12–14 | 4 | 120–150 | 10–14 Rs |
| 14–17 | 4 | 150–200 | 14–20 Rs |
| 17–20 | 4 | 200–280 | 20–30 Rs |
| 20–25 | 4 | 280–380 | 30–45 Rs |
| > 25 | 4 | 380+ | 45+ Rs |

Use this to detect **abnormally low premiums** (back ratio signal)
or **abnormally high premiums** (aggressive BF entry signal).

---

## PART 12 — RISK MANAGEMENT

### Position Limits

| Limit | Threshold | Action When Hit |
|---|---|---|
| Daily loss | 2% of capital | Stop all trading for the day |
| Weekly loss | 5% of capital | Review before continuing |
| Monthly loss | 10% of capital | Pause and full analysis |
| Single trade max loss | 0.5% of capital | Max debit per position |

### Greeks Exposure Limits

| Greek | Target | Max Allowed |
|---|---|---|
| Net Delta | ~0 (structure maintains) | ±50 |
| Net Gamma | ~0 (structure maintains) | ±20 |
| Net Theta | Positive (always) | > +500 Rs/day |
| Net Vega | Negative (we are short vol) | −500 to −3,000 Rs/VIX pt |

> **Portfolio Vega Rule:** Net portfolio vega must not exceed −5,000 Rs/VIX point.
> Track TOTAL portfolio vega, not per-position. Multiple positions compound each other.

### Margin Rules

Before any new position:
Available margin > new margin requirement × 2 (50% buffer)

Before scaling:
Available margin > scale margin requirement × 3 (67% buffer)

Maximum deployed: 70% of total capital
Minimum reserve: 30% of total capital (for scaling opportunities)


> **If 100% deployed → cannot scale → miss the best opportunity.**
> Keep powder dry for VIX spikes.

---

## PART 13 — PSYCHOLOGY & DISCIPLINE

### The Non-Negotiable Rules
Value% > 25% → NEVER ENTER. No exception.
Hold% > 25% → EXIT IMMEDIATELY. Do not wait.
Margin < 30% → Do NOT scale. Do not over-leverage.
Major event today → Do NOT enter new positions.
Daily loss limit hit → Stop trading for the day. Come back tomorrow.

### Common Mistakes to Avoid

| Mistake | Correct Behavior |
|---|---|
| Holding OTM positions hoping for recovery | Hold% > 25% → exit. Re-enter at ATM-closer strike. |
| Not scaling during VIX spikes (fear) | VIX spike = gift. Scale if Scale% ≤ 20%. |
| Exiting profitable positions too early | Use 1.5×, 2.0× booking rules. Let winners run. |
| Using ATM blindly as middle strike | Always check OI concentration first. |
| Ignoring VIX direction | VIX falling = enter. VIX rising fast = wait. |
| Panic-exiting during M2M loss | M2M loss from VIX = opportunity to scale, not exit. |
| Missing the expiry day booking window | Book by 2 PM on expiry, no exceptions. |
| Breaking rules "just this once" | The first exception creates a habit. Habits create blowups. |

### Patience Rules

Rule: Do NOT adjust for moves < 50 pts.
Market noise = ±30–50 pts constantly.
Reacting to noise = death by a thousand cuts.

Rule: One losing week does not break the strategy.
Probability works over MANY trades, not single ones.
Monthly outcomes matter. Quarterly outcomes matter most.

Rule: Review every trade, win or lose.
Every week answer:
→ What did I do right?
→ Which rule did I break?
→ What will I do differently next week?


---

## PART 14 — ENTRY CHECKLIST

> Run before EVERY trade. ALL items must be checked.

═══════════════════════════════════════════════════════
PRE-TRADE CALIBRATION
═══════════════════════════════════════════════════════
□ ATM Straddle: _________ Rs
□ Gap calculated: _________ pts
□ Combination: 1:2 / 2:3 / 3:4
□ Adverse tolerance: _________ pts
□ VIX level: _________
□ VIX direction today: Falling / Flat / Rising
□ Overshoot ratio: _________ (move ÷ straddle)
□ OI concentration checked: Middle strike = _________
□ PCR: _________ (bull/neutral/bear)
□ Events today: None / EVENT → STOP if EVENT

═══════════════════════════════════════════════════════
VALUE CONCEPT
═══════════════════════════════════════════════════════
□ Strike (K): _________
□ FLP: _________ Rs (must be > ₹8)
□ BF Rate: _________ Rs (must be > ₹0.5)
□ Value%: _________ % (must be < 25%)
□ Grade: Ideal / Good / Pass / Border / FAIL

═══════════════════════════════════════════════════════
CHAIN CONCEPT
═══════════════════════════════════════════════════════
□ Scenario UP — BF = _____ Rs (≥ 40% of entry?) Y/N
□ Scenario FLAT — BF = _____ Rs (≥ entry × 1.1?) Y/N
□ Scenario DOWN — BF = _____ Rs (≥ 40% of entry?) Y/N
□ Scenarios survived: ___ / 3 (need ≥ 2)

═══════════════════════════════════════════════════════
ENTRY SCORE
═══════════════════════════════════════════════════════
□ Value% score: ____ / 60
□ VIX direction score: ____ / 40
□ Overshoot score: ____ / 40
□ Double lock bonus: ____ / 30
□ PCR score: ____ / 10
□ Chain concept score: ____ / 20
□ TOTAL SCORE: ____ / 200 (minimum 60)

═══════════════════════════════════════════════════════
SIZING & MARGIN
═══════════════════════════════════════════════════════
□ Lot size (by value% tier): ____ buy lots
□ Margin required: _________ Rs
□ Margin available: _________ Rs (need > required × 2)
□ Post-trade buffer: _________ Rs (need > 30% of capital)

═══════════════════════════════════════════════════════
TRADE RECORD
═══════════════════════════════════════════════════════
□ Entry date: ___________
□ Strike (K): ___________
□ Gap: ___________ pts
□ Type: CE / PE
□ Combination: 1:2 / 2:3 / 3:4
□ Expiry: ___________
□ Entry debit: _________ Rs
□ Entry FLP: _________ Rs
□ Entry Value%: _________ %
□ Entry VIX: _________
□ ATM Straddle at entry: _________ Rs
□ Stop loss level: _________ Rs (debit ÷ 3)
□ Partial booking level: _________ Rs (debit × 1.5)
□ Full booking level: _________ Rs (debit × 2.0)
□ Double lock achieved: YES / NO
□ Chain scenarios survived: ___ / 3


---

## PART 15 — EXIT CHECKLIST

> Run daily for every open position.

═══════════════════════════════════════════════════════
DAILY HOLD CHECK
═══════════════════════════════════════════════════════
□ Current BF Rate: _________ Rs
□ Current FLP: _________ Rs
□ Hold% = Entry ÷ FLP: _________ %

□ ≤ 15% → HOLD (no action)
□ 15–20% → HOLD (check chain)
□ 20–25% → REVIEW (run chain check now)
□ > 25% → EXIT IMMEDIATELY

═══════════════════════════════════════════════════════
STOP LOSS CHECK
═══════════════════════════════════════════════════════
□ Stop level = Entry ÷ 3 = _________ Rs
□ Current BF < Stop? YES → EXIT / NO → Continue

═══════════════════════════════════════════════════════
PROFIT BOOKING CHECK
═══════════════════════════════════════════════════════
□ Partial level = Entry × 1.5 = _________ Rs
□ Current BF ≥ partial level? YES → Book 50% / NO → Continue
□ Full level = Entry × 2.0 = _________ Rs
□ Current BF ≥ full level? YES → Book 100% / NO → Continue

═══════════════════════════════════════════════════════
MIDDLE EXCEEDED CHECK
═══════════════════════════════════════════════════════
□ Market at: _________
□ Middle strike at: _________
□ Market crossed middle? YES → EXIT 100% / NO → Continue

═══════════════════════════════════════════════════════
ITM CHECK
═══════════════════════════════════════════════════════
□ Position ITM? YES / NO
□ If ITM: Hold% threshold tightened to 15% (not 25%)
□ Chain check — next adverse BF = _________ Rs
□ < 40% of entry? YES → EXIT / NO → Hold

═══════════════════════════════════════════════════════
COMBINED PORTFOLIO CHECK (when running pairs)
═══════════════════════════════════════════════════════
□ Combined Hold% = (all debits) ÷ (all FLPs) = _________ %
□ < 20% → Hold both
□ 20–25% → Exit breached leg, keep compensating leg
□ > 25% → Exit both

═══════════════════════════════════════════════════════
DTE CHECK
═══════════════════════════════════════════════════════
□ DTE remaining: _____ days
□ DTE 0 (expiry today): Book all by 2 PM
□ DTE 1: Book at 1.3× target, tight stop
□ DTE 2: Book at 1.5× target
□ DTE ≥ 3: Standard rules apply


---

## PART 16 — PERFORMANCE TRACKING

### Per-Trade Required Fields

Trade ID | Entry Date | Exit Date | Strike | Type | Gap |
Combo | Expiry | Entry Debit | Entry FLP | Value% | Score |
Grade | VIX at Entry | Straddle at Entry | Chain Survived |
Exit BF | Exit Reason | Hold% at Exit | Days Held |
P&L Points | P&L Rs | Lots | Double Lock | Partial Booked


### Weekly Review Questions
How many trades entered this week?
Average Value% at entry?
Did I follow every checklist item? Which did I skip?
Most common exit reason this week?
VIX trend during entries — was I timing correctly?
Did I scale during VIX spike? Outcome?
Did I maintain minimum 30% margin buffer?
One trade I wish I had NOT taken — why?
One opportunity I missed — why?
One thing I will do differently next week?

### Performance Benchmarks

| Metric | Target |
|---|---|
| Win Rate | > 60% |
| Avg P&L per trade | > 0.5× avg entry debit |
| Profit Factor | > 1.5 (total wins ÷ total losses) |
| Max Monthly Drawdown | < 10% of capital |
| Sharpe Ratio (annual) | > 1.5 |
| Avg Days Held | 2–4 days |

> If any metric misses target for **4+ consecutive weeks** →
> full rule review before continuing.

---

## THE TEN COMMANDMENTS

I. Value% > 25% → NEVER ENTER. No exception. No discussion.

II. Hold% > 25% → EXIT IMMEDIATELY. Do not wait for recovery.

III. VIX direction matters more than VIX level.
Falling VIX = enter. Rising VIX fast = wait.

IV. Middle strike = OI concentration peak.
Do not blindly use ATM.

V. M2M loss from VIX spike = scaling opportunity.
Do NOT exit because of M2M loss.

VI. When market stabilizes → book OTM legs, hold ATM legs.

VII. Keep minimum 30% margin buffer at all times.
This 30% IS the scaling weapon. Never use it for new entries.

VIII.Run call AND put butterflies together.
Evaluate combined Hold%. One side can compensate the other.

IX. Overshoot ratio > 1.0 = PRIME entry time.
Market has done the work. Theta does the rest.

X. Number of trades MUST be high.
Probability only works over many trades.
Small consistent profits accumulate into large monthly returns.


---

*Document: butterfly_trade_rules.md*
*Version: 3.0*
*Strategy: NIFTY 50 Multi-Expiry 1:2:1 Butterfly*
*Based on: Course rules + Live market evidence (22856 session) + Backtesting*
*Review cycle: Every 50 trades or 3 months, whichever comes first*
*Next review: After 50 completed trades*
