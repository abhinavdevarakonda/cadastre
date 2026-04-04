# Maplet Feature: Structure
## The Goal: Mapping the City

The **Structure** feature is Maplet’s "Reconnaissance Mission." Before we can trace anything, we must have a complete map of the city (the project).

---

### 1. The Analogy: The City Map
Imagine the project is a city. 
- **Directories** are the Boroughs.
- **Files** are the Buildings.
- **Functions** are the individual Apartments inside those buildings.

Maplet starts by flying a drone over the city to identify every building and then walking through the hallways of every building to find the apartments.

### 2. The Path of Data (Step-by-Step)

#### Step A: The Foot Patrol (`analyzer.Scan`)
Maplet (the Go binary) recursive walks through your project directory. It ignores "high-noise" areas (like `.git` or `__pycache__`) and builds a list of every file and folder it finds.

#### Step B: The Building Inspection (`lang.Extractor`)
For every building (file), Maplet checks the extension. If it's a `.py` building, it brings in the **Python Inspector.**
- The Inspector reads the source code.
- It looks for the "Function Definition" pattern (e.g., `def my_func(x):`).
- It records exactly where that apartment is located (File + Line Number).

#### Step C: The Central Registry (`graph.Graph`)
Every apartment (function) found is registered in a central database called the **Graph**. 
- **Node ID:** A unique ID for the function (e.g., `package.module.function_name`).
- **Metadata:** The file path, the name, and the line numbers.

#### Step D: The Display (TUI)
The TUI (the "Handler's Dashboard") reads the Central Registry. It uses a **Tree Layout** to show you the hierarchy. When you expand a folder, it simply asks the Registry: "Show me all buildings and apartments that are contained in this Borough."

---

### 3. Why this is "Agnostic" (The Future)
In the future, the "Inspector" won't be a human (a hardcoded Go script); it will be **Tree-sitter.** 
Instead of the Inspector needing to "know" Python, they will just follow a set of **Rules** (the `.scm` file) that tells them: "Look for anything that has this shape." 
