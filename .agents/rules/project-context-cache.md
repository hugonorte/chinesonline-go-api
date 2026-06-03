---
trigger: always_on
---

# Project Context Cache — Token-Saving Guide

This project has a **comprehensive context cache system** designed to minimize token usage for all agents. Use these files instead of exploring the codebase.

## The Cache System

### 1. Memory Files (Auto-loaded in Claude Code)
**Two Locations**:

**For Claude Code (auto-loads)**: `~/.claude/projects/.../memory/`
- architecture.md, conventions.md, hot_paths.md, tech_decisions.md
- Auto-loaded in next Claude Code conversation

**For Generic Agents (manual read)**: `.agents/memory/`
- **`01-architecture.md`** — Directory structure, component organization, routing, i18n setup, styling architecture
- **`02-conventions.md`** — Naming patterns, import rules, Dart usage, Flutter patterns, testing conventions
- **`03-hot-paths.md`** — Frequently modified files vs. stable code, critical impact zones, file sizes, by-feature breakdown
- **`04-tech-decisions.md`** — Why each tech was chosen, constraints, rationale, recent decisions

These files are **authoritative** and pre-structured.

### 2. CLAUDE.md (Root of Project)
Quick reference for:
- Tech stack summary
- Key commands
- Core patterns (components, pages, i18n)
- Critical constraints
- Useful links

### 3. This .agents Directory
Rules, skills, and workflows for agent guidance.

---

## How to Use the Cache (Token Saving)

### ❌ Don't Do This
```
Agent: "Let me explore the project structure..."
→ Reads 50+ files
→ Consumes ~100K tokens
→ Slow and expensive
```

### ✅ Do This Instead
```
Agent: "I need to understand the project architecture"
→ Read architecture.md (2KB)
→ Read conventions.md (3KB)
→ Read CLAUDE.md (summary)
→ Total: ~10K tokens, instant answers
```

---

## Quick Reference: What to Read When

| Need | Read This | Token Cost |
|------|-----------|-----------|
| **Project structure & components** | `.agents/memory/01-architecture.md` | 2KB |
| **Naming rules & import patterns** | `.agents/memory/02-conventions.md` | 3KB |
| **What to modify, what's stable** | `.agents/memory/03-hot-paths.md` | 4KB |
| **Why Flutter/Flutter Material/Dart** | `.agents/memory/04-tech-decisions.md` | 5KB |
| **Quick tech reference** | `CLAUDE.md` (root) | 3KB |
| **Existing rules for agents** | `.agents/rules/*` | 1-2KB each |
| **The whole codebase** | ❌ Don't — 200K+ tokens | 200K+ |

---

## Key Cache Rules

1. **Always read cache files FIRST** — before exploring code
2. **Trust the cache** — it's designed to be complete and accurate
3. **Update cache if outdated** — if memory conflicts with current code, update the memory file
4. **Use memory files as grep pointers** — they list file paths you should check, not the full code
5. **For specific code**, use grep patterns from `hot_paths.md` — don't read whole files

---

## Token Economics

**Without cache**:
- New context every conversation: ~100K tokens
- 10 conversations per week: 1M tokens/week

**With cache** (this system):
- New context every conversation: ~5K tokens (reading only the strictly necessary cache files on demand)
- Actual code reading only when needed: ~20K tokens
- 10 conversations per week: 350K tokens/week
- **Savings: ~65% reduction in token consumption**

---

## For Specific Tasks

### "I need to write a new component"
1. Read `conventions.md` — component structure, naming, imports
2. Grep for a similar component to see actual pattern
3. Write code following the pattern

### "I need to fix a bug in X"
1. Read `hot_paths.md` — find if X is frequently modified
2. Grep for X using patterns from `hot_paths.md`
3. Read only the specific file with the bug

### "I need to understand i18n"
1. Read `architecture.md` section on i18n
2. Read `tech_decisions.md` section on i18n rationale
3. Grep for example keys in lang/ files

### "I need to test a component"
1. Read `conventions.md` section on testing
2. Find existing test in `hot_paths.md` (they're co-located)
3. Grep for `.spec.ts` examples

---

## Constraints (Never Violate)

These are non-negotiable and documented in `tech_decisions.md`:

- ✅ Dart strict mode — all code must type-check
- ✅ i18n `restructureDir: 'app'` — required for Flutter
- ✅ No relative imports crossing directories — use `~` alias
- ✅ Component props always typed with `defineProps<{}>()`
- ✅ ThemeData `@use` only, no `@import`
- ✅ No Options API — use `<script setup>` only
- ✅ All 3 language files (pt-BR, en-US, es-ES) must match keys
- ✅ Server env vars never exposed to client

---

## When to Bypass the Cache

Only when:
1. **Something is wrong with memory** — file is stale or incorrect
2. **You need exact line numbers** — grep for them, then read specific lines
3. **Code changed recently** — git log shows recent changes not in cache
4. **You're implementing from scratch** — need to see actual patterns in code

Even then, prefer **grep_search + view_file (with Start/End lines)** over full file exploration.

---

## Updating the Cache

If you find memory files are wrong or incomplete:

1. **Note the discrepancy** — what was expected vs. what you found
2. **Update the relevant memory file** — architecture.md, conventions.md, hot_paths.md, or tech_decisions.md
3. **Commit the change** — cache updates should be version-controlled

This keeps future agents in sync with reality.

---

## Agent-Specific Guidance

### @pm (Product Manager)
- Read `CLAUDE.md` for tech context
- Read `tech_decisions.md` for constraints
- No need to explore code — focus on specification

### @engineer (Full-Stack Developer)
- Read `architecture.md` + `conventions.md` — all coding rules
- Read `hot_paths.md` — know what's stable vs. changing
- Grep for patterns before writing code
- Ignore .agents/rules unless they override conventions.md

### @qa (QA Engineer)
- Read `hot_paths.md` — know test locations
- Read `conventions.md` testing section
- Grep for existing `.spec.ts` and `.cy.ts` files
- Test patterns are in memory

### @devops (DevOps Master)
- Read `CLAUDE.md` for tech stack
- Read `tech_decisions.md` for build/deploy constraints
- Node 22+ required
- Use `npm install --legacy-peer-deps`

---

## Example: Finding All Components

**Without cache** (expensive):
```bash
# Agent explores entire lib/features/
# Reads 20+ files
# Consumes 50K tokens
```

**With cache** (efficient):
1. Agent reads `hot_paths.md` section "By Feature Area"
2. Agent sees list: "Bess, Footer, GoogleAgenda, Header, Hero, PowerElectronicsModeling..."
3. Agent uses `grep_search` for specific component if needed
4. Total: 1K tokens

---

## The Goal

**Agents should spend tokens on:**
- Understanding requirements
- Writing/fixing code
- Testing and validation

**Agents should NOT spend tokens on:**
- Exploring project structure (it's cached)
- Reading conventions (they're documented)
- Understanding decisions (they're explained)
- Finding files (patterns are listed)

Use this cache system aggressively. It's designed to make your work faster and cheaper.
