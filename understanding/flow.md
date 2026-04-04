# Maplet Feature: Flow
## The Goal: Real-Time Shadowing

The **Flow** feature is the most advanced part of Maplet. It is the only "Dynamic" feature—where we watch the code execute live.

---

### 1. The Analogy: The Tail (Shadowing)
Instead of just reading a map of the city, we are now **Shadowing** the suspect. We are following them in a car as they drive from building to building. We record every single stop they make, the order they make them in, and a live "Heatmap" of which buildings they visit the most.

### 2. The Path of Data (Step-by-Step)

#### Step A: Planting the Spy (Injection)
When you run `maplet flow --target "python app.py"`, Maplet starts by "planting a spy" inside the Python process. 
- It sets a `PYTHONPATH` variable that forces Python to load our **Agent** (`sitecustomize.py` and `py_trace.py`).
- The spy is now "inside" the Gang (the process).

#### Step B: The Secret Handover (The Socket)
The Spy has a direct radio link (a **Unix/TCP Socket**) back to the Handler (Maplet Go binary).
- Every time a function is called, the Spy whispers: `"Call: handle_order, File: main.py, Line: 10"`.
- This happens hundreds of times a second.

#### Step C: The Live Feed (`TraceEventMsg`)
Maplet (the Go process) receives these whispers. It does two things instantly:
1.  **The History Log:** It adds the call to a "Sequence" list (the Flow sequence).
2.  **The Heatmap:** It increments a "Hit Counter" for that function.

#### Step D: The Dashboard (TUI)
The TUI updates in real-time:
- **Glowing Selection:** The most recent call glows yellow.
- **Heatmap Colors:** Functions that are called many times turn **Red/Blazing**.
- **Scrubbing:** You can pause the execution and "Scrub" through time using `H/L` to see exactly what happened 30 calls ago.

---

### 3. The Power of the Socket
Because the Spy and the Handler talk over a **Socket**, the Suspect can even be inside a Docker container or on a different server. As long as the Spy can "radio back" the JSON data, Maplet can draw the live flow.
