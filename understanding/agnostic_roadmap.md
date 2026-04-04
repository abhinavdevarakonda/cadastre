# Maplet: Universal Language Roadmap

To scale Maplet beyond Python into a universal code-navigation engine, we must transition from **Hardcoded Language Drivers** to a **Pattern-Driven Engine**.

---

## 1. The Strategy: Pattern-Driven Extraction

The core of Maplet's future is **Tree-sitter**, a high-performance incremental parsing library. Instead of writing Go code to find functions for every language, we use **Language Queries (.scm files)**.

### What is an `.scm` file?
It stands for **Scheme** (a Lisp-like syntax). In Tree-sitter, it’s used for **Structural Pattern Matching**. 

Think of it as **Regex for Trees**:
1.  **Direct Pattern Matching:** You define the *structure* of a function declaration, not just the text.
2.  **Captures:** You specify which parts to "capture" (e.g., the function name or the body).
3.  **Speed:** Tree-sitter scans these queries in a single pass of the AST (Abstract Syntax Tree).

### The `python.scm` Blueprint
Instead of 100 lines of Go code, Maplet will find every Python function with just this:
```scheme
;; Pattern: Capture any function definition and its name
(function_definition
  name: (identifier) @name) @function

;; Pattern: Capture any class and its name
(class_definition
  name: (identifier) @name) @class
```

### The Universal Approach
To support a new language like **Rust** or **Mojo**, you don't touch the Go code. You simply:
1.  Add the Rust grammar to the binary.
2.  Drop a `rust.scm` file into the `queries/` folder.
3.  The engine automatically "sees" the new patterns and builds the Maplet graph.

---

## 2. Universal Dynamic Tracing (Socket Protocol)

Tracing execution is runtime-dependent. To make this agnostic, Maplet becomes a **Receiver** over a local socket.

### How it works:
1.  **The Socket:** Maplet opens a raw TCP/Unix socket (e.g., `/tmp/maplet.sock`).
2.  **The Agents:** Each language (Python, Node.js, Ruby) gets a tiny, one-file "Agent."
    - The Agent connects to the socket and streams execution events in a standard JSON format.
3.  **Agnosticism:** Maplet doesn't need to know *how* the execution is being traced—it just draws the JSON events it receives.

---

## 3. Initial Implementation Steps (The Branch Strategy)

To start the **Agnostic Branch**, we build the foundation without the complex binary dependencies:

1.  **`queries/`**: The root-level store for all `.scm` patterns.
2.  **`internal/lang/universal/`**: The Go interface that handles loading queries and parsing files universally.
3.  **`maplet.yaml`**: The project manifest where users can define their own custom grammars and queries.

---

## 4. Phase-by-Phase Roadmap

### Phase 1: Logic-as-Data (Static Analysis)
- [ ] Integrate `github.com/smacker/go-tree-sitter`.
- [ ] Move Python extraction to the `.scm` query engine.
- [ ] Add Go/JavaScript support by simply adding `.scm` files.

### Phase 2: The Socket Receiver (Dynamic Analysis)
- [ ] Decouple the Python tracer from the TUI.
- [ ] Define a standard JSON schema for Execution Events.
- [ ] Build a tiny Node.js Agent as a proof of concept.

### Phase 3: Manifest-Driven Intelligence
- [ ] Implement `maplet.yaml` to allow zero-config project analysis.
- [ ] Allow users to override `.scm` queries for their specific framework styles.

