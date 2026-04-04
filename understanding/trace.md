# Maplet Feature: Trace
## The Goal: Mapping the Suspect's Contact List

The **Trace** feature (static) is the opposite of Impact. It answers the question: "Who does this function depend on? What functions does it call during its execution?"

---

### 1. The Analogy: The Contact List
Every suspect (function) has a "Contact List" in their phone. These are the people they reach out to when they need to get a job done. The **Trace** analysis opens that contact list and shows you every outgoing call.

### 2. The Path of Data (Step-by-Step)

#### Step A: Outgoing Search (`Fact Extraction`)
Just like in Impact mode, the Inspector has already identified every call expression. In this mode, we look at the calls **originating** from our target.

#### Step B: The Forward Lookup (`analyzer.TraceAnalysis`)
When you select a function and switch to Trace mode (`t`), Maplet asks the Registry: "Give me every Edge where the **'From'** field matches my current function."

#### Step C: Resolving the Target
This is the hardest part. If the code says `utils.save()`, Maplet must figure out *which* `save` function in the entire city it’s talking about.
- It looks at imports.
- It looks at package names.
- It finds the unique "Apartment ID" for that callee.

#### Step D: The Callee Preview
The TUI shows you the **Callees** (the people being called). 
- **Stacked View:** You see the callee's name, their file, and their function definition signature.
- **Interaction:** If you hit `enter` here, Maplet teleports you directly to that callee's building (file) so you can see how they work.

---

### 3. Difference between Static Trace and Dynamic Flow
- **Trace (Static):** Shows you who the function *could* call (based on reading the code).
- **Flow (Dynamic):** Shows you who the function *actually* called (during a live run).
