# LazyNotes Reference

*LazyNotes* is a simple terminal note management program written in Go. It's
effectively a vim/neovim wrapper that allows for navigation and editing of notes
in a directory. Notes are formatted in markdown (stored as `.md` files) and may
be grouped by folder (no recursive folders).

## Comment Rules

Follow three rules for comments:

- **Rule 1 — Names explain *what***: Choose clear, descriptive names for
classes, methods, and variables. If the name is good enough, no comment is
needed to explain what it does.
- **Rule 2 — Code explains *how***: The code itself should be readable enough
to show how things work. Don't write comments that restate the code.
- **Rule 3 — Comments explain *why***: Only add comments when the reason
behind a decision isn't obvious from the code. Explain *why* this approach was
chosen, *why* a workaround exists, or *why* a non-obvious value is used.

## Log Format

When appending to logs, use this format:

```markdown
## YYYY-MM-DD — Short Title

- What was done (bullet points)
- Files created or modified
- Tests added or updated
- Any issues encountered
```

> This project uses Git commits as a log. All commit descriptions should be
> formatted in this manner. Commit titles use *conventional commits*.

## Versioning

The project is versioned via the `Major.Minor.Patch` convention formatted as `vX.Y.Z`:

- **Major (`vX.0.0`)** — multiple large features, large breaking changes, or
  major refactoring.
- **Minor (`v0.Y.0`)** — a single large feature, moderate refactoring,
  additions to existing features, or a large group of bug fixes.
- **Patch (`v0.0.Z`)** — targeted bug fixes only; no new features, no sweeping patches.

> This project uses GitHub Releases for version management.

## Build Plan

### v0.0.1 — Bare Minimum Functionality

- Two windows: note selector on the left and display on the right.
- Note selector displays file tree for a hard-coded directory.
- Display window shows raw text content of `.md` files (no rendering yet).
- Keys allow navigation between `.md` files in left pane.
