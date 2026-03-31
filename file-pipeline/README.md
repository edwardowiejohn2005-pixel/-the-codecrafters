# File Pipeline — Operation Gopher Protocol

## How to Run

```bash
go run . input.txt output.txt
```

---

## The 5 Transformation Rules (in order)

| # | Rule |
|---|------|
| 1 | **Trim whitespace** — remove all leading and trailing spaces from every line |
| 2 | **Remove blank/dash lines** — drop lines that are empty or made of only dashes |
| 3 | **Replace TODO:** — `TODO:` (any case) becomes `✦ ACTION:` |
| 4 | **Redact CLASSIFIED:** — `CLASSIFIED:` (any case) becomes `[REDACTED]:` |
| 5 | **Title-case ALL CAPS** — lines where every letter is uppercase get converted to Title Case |

---

## Output Format

The output file contains:

1. A header block:
   ```
   SENTINEL FIELD REPORT — PROCESSED
   ════════════════════════════════════════
   Processed: 2026-03-31 20:00:00
   ════════════════════════════════════════
   ```

2. Processed lines, each prefixed with a 3-digit number:
   ```
   001. First line here
   002. Second line here
   ```

3. A summary block at the end:
   ```
   ════════════════════════════════════════
   SUMMARY
     Lines read    : 19
     Lines written : 15
     Lines removed : 4
     Rules applied : ...
   ════════════════════════════════════════
   ```

---

## What Happens When Rules Overlap

Rules fire **in order** on each line. Here's what happens in overlap cases:

- **ALL CAPS + starts with TODO:**
  Rule 3 fires first and changes `TODO:` → `✦ ACTION:`. Then Rule 5 checks if the line is still all caps — because the `✦ ACTION:` prefix now contains lowercase letters, the line will **not** be title-cased. The TODO replacement wins.

- **CLASSIFIED: + long line:**
  Rule 4 replaces the prefix. The resulting line may be shorter or longer — no truncation rule is applied (not one of our 5), so it passes through as-is.

- **Blank line after trimming (Rule 1 → Rule 2):**
  A line of only spaces becomes `""` after Rule 1. Rule 2 then catches it as blank and removes it. Both rules fire in sequence — this is intentional.

- **ALL CAPS + starts with CLASSIFIED:**
  Rule 4 fires first and adds `[REDACTED]:` (mixed case). Rule 5 then checks — the line is no longer all caps, so title-case conversion does **not** apply.

---

## Example Run 1

**Input (excerpt):**
```
ALERT LEVEL FIVE DETECTED IN NORTHERN SECTOR
todo: confirm coordinates with Agent Bulus
classified: target has been relocated
-----------------------------------------------
```

**Output (excerpt):**
```
001. Alert Level Five Detected In Northern Sector
002. ✦ ACTION:  confirm coordinates with Agent Bulus
003. [REDACTED]:  target has been relocated
```
(The dash line is removed — not written.)

---

## Example Run 2

**Input:**
```
   extra spaces here   
```

**Output:**
```
001. extra spaces here
```
(Spaces stripped by Rule 1.)

---

## Edge Cases Handled

| Case | Behaviour |
|------|-----------|
| Input file missing | Prints `✗ File not found:` and exits |
| Empty input file | Writes empty output, prints warning |
| All lines blank/dashes | Output has only header + summary |
| Same file as input and output | Rejected before processing begins |
| Output path is a directory | Rejected with clear error |
| Windows line endings (`\r\n`) | `\r` stripped before processing |
| Line that is only whitespace | Treated as blank — removed |
| Multiple rules on one line | Applied strictly in order 1→5 |