# Maplet Flow: Debugger Architecture Deep-Dive

Maplet Flow doesn't just "watch" your program; it **instruments** it. It behaves like a lightweight debugger (similar to `pdb` or `gdb`) but focuses on visualizing the high-level execution path rather than just inspecting variables.

Here is the step-by-step breakdown of how the machine works.

---

## 1. The Sensor: [py_trace.py](file:///home/abee/code/projects/maplet/internal/tracer/py_trace.py) (The Hook)
The magic starts inside your application's process. When you run `maplet run "python app.py"`, Maplet wraps your script with a specialized tracer.

### The `sys.settrace` Hook
Python provides a powerful primitive called `sys.settrace(callback)`. Our tracer uses this to register a "Global Trace Function." 
- Every time Python enters a new function, it pauses for a microsecond and calls our [trace_calls](file:///home/abee/code/projects/maplet/internal/tracer/py_trace.py#35-68) function.
- We receive a `frame` object, which contains the function name, the file path, and the current line number.

### Intelligent Filtering
To prevent "Noise Bloat" (e.g., tracing internal Python library calls or 3rd party modules), we apply strict filters:
- **Exclude Lists**: We ignore anything inside `lib/`, `site-packages`, or `<frozen>` modules.
- **Root Sync**: We automatically detect your project's root directory and only record hits that happen inside your code.

---

## 2. The Intercom: Async TCP Communication
Once a hit is detected, it needs to get to the Maplet TUI. We use a **Low-Latency Intercom System** over local TCP (Port 9876).

### The Sender Thread
Inside your app, we spawn a dedicated **Background Python Thread**. 
1. **The Queue**: The tracer (the "Sensor") drops JSON events into a fast internal queue. This ensures your app doesn't slow down waiting for the network.
2. **The Reconnector**: If the Maplet Monitor isn't open yet, the background thread silently retries once every second. This makes the system "Accident-Prone"—you can close and reopen the TUI at any time without crashing your app.
3. **The Payload**: Each message is a tiny JSON blob:
   ```json
   {"event": "call", "name": "process_request", "file": "app.py", "line": 42}
   ```

---

## 3. The Ear: `tracer.Listen` (The Server)
Inside the Maplet Go backend ([internal/tracer/tracer.go](file:///home/abee/code/projects/maplet/internal/tracer/tracer.go)), we run a TCP server.

- **Concurrent Monitoring**: Using Go's **Goroutines**, the server can listen to multiple processes at once (e.g., if you have a Frontend and a Backend running simultaneously, both can stream hits to a single Maplet window).
- **Scanner Pipeline**: We use a `bufio.Scanner` to read the incoming JSON stream. Each valid line is parsed into a Go `tracer.Event` struct and passed to the TUI.

---

## 4. The Nerve Center: The TUI Loop
The Maplet TUI ([internal/tui/tui.go](file:///home/abee/code/projects/maplet/internal/tui/tui.go)) is built using the **Bubble Tea** (Elm Architecture) framework.

### The `TraceEventMsg`
When the background Go listener hears a hit, it "pokes" the TUI by sending a `TraceEventMsg`.
- **History Storage**: The model maintains a slice called `m.history`. This is our **Master Tape**. It stores every function hit in chronological order.
- **Auto-Following**: If you are in "Live" mode, the `playhead` (the current view pointer) is pinned to `len(m.history) - 1`.

---

## 5. The Lens: Heatmaps & Visualization
The final step is turning that history into the "Wired" visual experience.

### The Heatmap Calculation
We maintain a high-speed map: `hitCounts map[string]int`. 
- Every time an event arrives, we resolve the file/function name to a node in our Static Graph.
- We increment the count for that node ID. 
- The UI then uses a **Thermal Scale** (Cool Green -> Warm Orange -> Hot Red) to color the tree icons based on their intensity.

### The High-Contrast Jump
When you "Scrub" (move the playhead with `H/L`), the TUI:
1. Looks up the function name at the current playhead.
2. Automatically navigates the tree to that node.
3. **Expands all parent folders** if they are closed.
4. Highlights it with the **Yellow Glow** (High-Intensity Background).

---

## Summary of Execution Flow

1. **Invoke**: `maplet run "python app.py"`
2. **Hook**: `sys.settrace` identifies a call to `save_to_db()`.
3. **Queue**: Hit is queued in a background thread.
4. **Broadcast**: Hit is sent via TCP to `localhost:9876`.
5. **Receive**: Go backend parses JSON and pings the TUI.
6. **Update**: TUI adds to history and colors the node "Orange".
7. **View**: User sees an instant yellow flash on the left and a new entry on the right.
