# 🧠 Agent Memory System

This directory contains **cached project context** specifically for generic AI agents (not just Claude).

## What's Here

### `memory.md` (Main File)
A comprehensive, single-file knowledge base covering:
- Project structure
- Component patterns (copy-paste templates)
- Page patterns
- Composable patterns
- i18n setup (critical Flutter config)
- Testing patterns
- Styling rules
- Dart requirements
- Development commands
- Constraints (must not break)
- File locations
- Grep patterns for efficient searching
- Decision history
- Token-saving tips
- FAQ

**Size**: ~15KB  
**Read time**: 10-15 minutes (one-time per conversation)  
**Token cost**: ~15K tokens (vs. 100K+ exploring codebase)

---

## How Agents Should Use This

### First Time
1. Read `memory.md` completely (15 min)
2. Bookmark key sections (patterns, constraints, grep patterns)
3. Start working

### Every Task
1. Open `memory.md`
2. Search for relevant section (Ctrl+F)
3. Copy exact patterns
4. Follow constraints
5. Use grep patterns to find code
6. Never explore the full codebase

### When Something's Not Here
1. Check CLAUDE.md (quick reference at project root)
2. Check .agents/rules/ for agent-specific guidance
3. Check .agents/skills/ for searching techniques
4. Check .agents/workflows/ for process guidance
5. If still missing, update this memory.md file

---

## Token Economics

### Without This Memory
- Agent explores codebase: 50+ files read
- Token cost: 100K+ per conversation
- Time: 30-45 minutes understanding project
- Cost: $1.50-3.00 per conversation

### With This Memory
- Agent reads memory.md: 1 file (15KB)
- Agent greps for examples: 1-5 files
- Agent writes code: knows pattern
- Token cost: 20-30K per conversation
- Time: 5-10 minutes for understanding
- Cost: $0.30-0.60 per conversation

**Savings: 70-80% tokens, 83% faster**

---

## Structure of memory.md

The file is organized into 15 sections, each self-contained:

```
SECTION 1:  Project Structure
SECTION 2:  Component Patterns
SECTION 3:  Page Patterns
SECTION 4:  Composables
SECTION 5:  i18n (Internationalization)
SECTION 6:  Testing Patterns
SECTION 7:  Styling
SECTION 8:  Dart Requirements
SECTION 9:  Development Commands
SECTION 10: Constraints (CRITICAL)
SECTION 11: File Locations (Hot Paths)
SECTION 12: Grep Patterns
SECTION 13: Decision History
SECTION 14: Token-Saving Tips
SECTION 15: Related Resources
```

### Quick Search Guide

| Need | Find This | Search For |
|------|-----------|-----------|
| How do I make a component? | SECTION 2 | "Component Template" |
| What's the project structure? | SECTION 1 | "Directory Layout" |
| How do I add translations? | SECTION 5 | "Key Structure" |
| How do I test? | SECTION 6 | "Unit Tests" or "E2E Tests" |
| What can't I do? | SECTION 10 | "Never Violate These" |
| Where are the files? | SECTION 11 | "By Feature Area" |
| How do I find code? | SECTION 12 | "GREP PATTERNS" |

---

## Critical Constraints (From SECTION 10)

These **cannot be broken**:

1. ✅ **Dart Strict** — No `any` type
2. ✅ **i18n restructureDir** — Must be `restructureDir: 'app'`
3. ✅ **No relative imports** — Use `~` alias only
4. ✅ **Props always typed** — `defineProps<{}>()`
5. ✅ **ThemeData @use only** — No `@import`
6. ✅ **Script setup only** — No Options API
7. ✅ **All 3 language files** — Keep synced (pt-BR, en-US, es-ES)
8. ✅ **Server vars private** — Never expose to client

If you break any of these, tests will fail.

---

## Copy-Paste Templates

memory.md includes ready-to-use templates for:
- Component (with props, emits, reactive state, i18n)
- Page (with SEO meta)
- Composable (with state and methods)
- Test (unit with @Flutter/test-utils)
- Test (E2E with Flutter Integration Test)
- Dart types (props, emits)

**Just search for "Template"** in memory.md and copy.

---

## Grep Patterns (SECTION 12)

memory.md includes grep patterns for:
- Finding components
- Finding pages
- Finding composables
- Finding usage of specific code
- Finding i18n usage
- Finding test files
- Seeing recent changes
- Checking git diffs

**Copy these exactly** and paste in terminal. No exploration needed.

---

## Decision History (SECTION 13)

Recent changes and the reasoning:
- Why Flutter? (SSR/SSG + auto-routing)
- Why Dart strict? (catch bugs early)
- Why Flutter Test? (faster than Jest)
- What changed recently? (last 4 weeks)

Understand *why* decisions were made, not just *what* was decided.

---

## Updating This Memory

If you find:
- Something is wrong or outdated
- Something is missing
- A pattern changed
- New constraints were added

**See `UPDATE_TRIGGERS.md` for SPECIFIC conditions when to update.**

**General update process**:
1. Find the relevant section
2. Fix or add the information
3. Update the "UPDATE HISTORY" table at the end
4. Commit to git with message: `docs(memory): update [SECTION] for [change]`
5. (Optional) Share with team

This keeps future agents in sync with reality.

**Important**: Read `UPDATE_TRIGGERS.md` first to understand WHEN to update (don't update prematurely).

---

## For Different Agent Roles

### @pm (Product Manager)
Read: SECTION 1, 13, 10 (constraints)

### @engineer (Full-Stack Developer)
Read: SECTION 2-10 (all patterns + constraints)

### @qa (QA Engineer)
Read: SECTION 6, 10, 12 (testing + constraints + grep patterns)

### @devops (DevOps Master)
Read: SECTION 1, 9, 10 (structure + commands + constraints)

### Generic Agent (Any LLM)
Read: SECTION 1-15 (everything)

---

## Example: Add a Component

Without memory:
1. Agent explores: 50+ files
2. Agent searches for examples: 10+ more reads
3. Agent reads component code: another 30+ lines
4. Agent finally understands pattern
5. Agent writes code
6. Token cost: 80K+

With memory:
1. Agent reads memory.md SECTION 2 (copy-paste template)
2. Agent greps for similar component (1 read)
3. Agent writes code
4. Token cost: 5K

**Savings: 75K tokens** 🎉

---

## Example: Fix a Bug

Without memory:
1. Agent searches codebase for bug location
2. Agent reads 10+ files to understand context
3. Agent reads solution examples
4. Token cost: 50K+

With memory:
1. Agent reads SECTION 11 (file locations)
2. Agent greps for exact bug location
3. Agent reads 1-2 relevant files
4. Token cost: 10K

**Savings: 40K tokens** 🎉

---

## Example: Add i18n

Without memory:
1. Agent reads i18n config in code
2. Agent explores lang/ folder
3. Agent reads multiple examples
4. Agent finally understands the pattern
5. Token cost: 60K+

With memory:
1. Agent reads SECTION 5 (i18n setup + critical fix note)
2. Agent reads "Key Structure" example
3. Agent reads SECTION 5 "Rules (Critical)"
4. Agent writes translation keys
5. Token cost: 2K

**Savings: 58K tokens** 🎉

---

## Integration with Other Files

### In Claude Code (auto-loaded)
See: `~/.claude/projects/.../memory/` for auto-loading copies

### In .agents/ (agent guidance)
- `rules/project-context-cache.md` — References this memory
- `skills/search-code-efficiently.md` — Uses grep patterns from here
- `workflows/token-efficient-work.md` — Instructs reading this memory first

### At Project Root
- `CLAUDE.md` — Quick reference (links to this memory)

All files work together to save tokens.

---

## Verification Checklist

Before starting ANY work:

- [ ] You've read memory.md SECTION 1 (structure)
- [ ] You understand SECTION 10 (constraints)
- [ ] You know the pattern for your task
- [ ] You know where to find files (SECTION 11)
- [ ] You have a grep pattern ready (SECTION 12)
- [ ] You're not exploring the codebase

If all checked: **You're ready to work efficiently!**

---

## FAQ

**Q: Is this memory always up to date?**  
A: It should be. Check the "UPDATE HISTORY" at the bottom of memory.md. If outdated, update it.

**Q: Can I modify this memory?**  
A: Yes, please do. It's a living document. Update it when reality changes.

**Q: Should I read the entire file?**  
A: First time: yes. Then use Ctrl+F to search sections. No need to re-read everything.

**Q: What if memory.md doesn't have the answer?**  
A: Check CLAUDE.md (root), then .agents/ directory. If still missing, add it to memory.md.

**Q: How do I share this with another agent?**  
A: Just point them to `.agents/memory/memory.md`. It's self-contained and readable.

**Q: Is this replacing the actual code?**  
A: No. This is a **reference guide**, not the source of truth. Code is the source of truth. Memory summarizes it.

---

## Token Budget Examples

### Task: Add Component
```
Read memory.md SECTION 2:  2K
Grep for example:          1K
Write code:                0K
Total:                     3K tokens
Without memory:           80K tokens
Savings:                  77K tokens (96%)
```

### Task: Fix Bug
```
Read memory.md SECTION 11: 2K
Grep for location:         1K
Read buggy file:           5K
Write fix:                 0K
Total:                    8K tokens
Without memory:           50K tokens
Savings:                  42K tokens (84%)
```

### Task: Write Test
```
Read memory.md SECTION 6:  3K
Grep for test example:     1K
Write test:                0K
Total:                     4K tokens
Without memory:           30K tokens
Savings:                  26K tokens (87%)
```

---

## Support

**How do I use this memory?** → See "How Agents Should Use This" section above

**What's missing?** → Update memory.md and commit

**Is this secure?** → Yes, it's documentation, not secrets. No sensitive data stored.

**Can non-Claude agents use this?** → Yes, it's plain English Markdown. Any LLM can read it.

---

## Version History

| Date | Status |
|------|--------|
| 2026-05-04 | Created: memory.md + README.md |
| Next | Update as code evolves |

---

**Remember**: This memory exists to save tokens. Read it instead of the codebase. Update it when it becomes wrong. Share it with teammates.

**Goal**: Help all agents work faster, cheaper, and more consistently.

🚀 **Now stop reading docs and start building!**
