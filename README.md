# Maplet

Maplet is a terminal tool for exploring the structure of a codebase.

When working in an unfamiliar project, it can be hard to answer questions like:
    - Can I safely modify this function?
    - What does this actually trigger?
    - How is this function connected to the rest of the codebase?

Maplet builds a graph of directories, files, functions, and their call relationships. It provides commands and a terminal UI for navigating that structure, inspecting callers, tracing callees, and jumping directly to call sites.

# How it works
Maplet performs static analysis on a project and constructs a graph containing:
    - directories
    - files
    - functions
    - call edges between functions

This graph is used to support operations such as:
    - navigating functions within the project
    - finding functions that call a given function (impact)
    - tracing the functions a given function calls
    - locating call sites in source files

and the TUI exposes this information interactively so relationships can be explored without leaving the terminal.

It gives you the experience of an IDE in your terminal and helps you avoid jumping in and out, and up and down trying to make changes to fragile sections.


# Current features

**Project Navigation**
    `maplet nav`

opens an interactive terminal UI for exploring the structure of the project
You can jump to specific functions or even where they are called; all viewable with a glance.

**Impact Analysis**
    `maplet impact <symbol>`
Lists the functions that depend on the given function
This helps determine what parts of the code may be affected before modifying it.

**Project Analysis**
    `maplet analyze <path>`
Analyzes a project and prints detected functions and call relationships

**Graph Export**
    `maplet export <path>`
Exports the internal graph of the project as a json.

**MCP server**
    `maplet mcp <path>`
Starts an MCP server exposing the code graph for external tools or agents to query.
