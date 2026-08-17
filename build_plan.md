# Build Plan

## v0.1.0 — Bare Minimum Functionality (COMPLETE)

**Design Goals:**

- Two windows: note selector on the left and display on the right.
- Note selector displays file tree for a hard-coded directory.
- Display window shows raw text content of `.md` files (no rendering yet).
- Keys allow navigation between `.md` files in left pane.

## v0.2.0 — Proper Markdown Parsing and Rendering

**Design Goals:**

- Update `mdparse` package to properly process raw markdown from files.
- Display properly rendered markdown in the *detail* view when focusing a note.
