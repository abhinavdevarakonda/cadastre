# Maplet Feature: Impact
## The Goal: Finding the Incoming Phone Calls

The **Impact** feature is Maplet’s way of figuring out who communicates *with* our target. It answers the question: "If I change this function, who is going to be affected?"

---

### 1. The Analogy: The Phone Line Tap
If a suspect (a function) is interesting, we want to know every single person who has called their phone number. We aren't looking at who the suspect is calling—we are looking at who is **initiating** a conversation with them.

### 2. The Path of Data (Step-by-Step)

#### Step A: Identifying the Caller (`Fact Extraction`)
During the "Structure" phase, the Inspector also looks for **Call Expressions** (e.g., `other_function()`).
- When the Inspector sees a call, it records: "This Building/Apartment is reaching out to *FunctionName*."

#### Step B: Establishing the Edge (`graph.Edge`)
In the Central Registry, we create an **Edge** (a connection).
- **From:** The Caller.
- **To:** The Target.
- **Type:** `calls`.

#### Step C: The Reverse Search (`analyzer.ImpactAnalysis`)
When you select a function in the TUI, Maplet does a **Reverse Lookup** in the Registry. It asks: "Give me every Edge where the **'To'** field matches my current function."

#### Step D: The Impact Preview
For every caller found, Maplet goes to that caller's building (file) and reads the exact line where the call happens (the "Call Site"). This is what you see in the TUI’s right pane:
- The name of the Caller.
- The file they are in.
- A "preview" of the call site itself.

---

### 3. Why this matters
Impact analysis is the key to safe refactoring. By knowing all the callers, Maplet gives you a "blast radius" for any change you make. If a function has 10 callers, you have 10 places in the city you must check before making a change.
