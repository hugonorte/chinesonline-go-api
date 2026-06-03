# 📑 Memory Directory Index

Quick navigation for the agent memory system.

---

## Files in This Directory

### 1. `memory.md` (675 lines, 17KB)
**The main knowledge base for agents**

Comprehensive, single-file reference covering:
- 15 sections of project context
- Copy-paste templates
- Grep patterns
- Constraints
- File locations
- Decision history

**Use when**: You need detailed information about any aspect of the project

**Search patterns** (use Ctrl+F):
- "Component Template" → How to write a component
- "Page Template" → How to write a page
- "Testing Patterns" → How to test
- "GREP PATTERNS" → How to find code
- "Never Violate These" → Critical constraints
- "By Feature Area" → Where files are
- "Key Structure" → How i18n works

**Read time**: 15 minutes first time, then use Ctrl+F for lookups

**Token cost**: ~15K (one-time per conversation)

---

### 2. `README.md` (380 lines, 9.6KB)
**How to use the memory system**

Explains:
- What's in this directory
- How agents should use it
- Token economics (70-80% savings)
- Quick search guide
- Critical constraints
- Copy-paste templates location
- Update instructions
- Role-specific reading paths
- Examples (how it saves tokens)
- FAQ

**Use when**: You're new to this memory system or need guidance

**Read time**: 10 minutes

**Token cost**: ~10K

---

### 3. `INDEX.md` (This File)
**Navigation guide**

Points to:
- What each file contains
- When to read each file
- Search strategies
- Example workflows
- Quick reference matrix

**Use when**: You're lost and need to find something

**Read time**: 3 minutes

**Token cost**: ~5K

---

### 4. `UPDATE_TRIGGERS.md` (NEW)
**When agents should update memory files**

Specifies:
- EXACTLY when to update (.agents/memory/*.md files)
- 5 categories of changes (Architecture, Conventions, Hot Paths, Tech Decisions, Cross-File)
- Update frequency guidelines
- Specific trigger examples
- What NOT to update
- Priority matrix for decisions

**Use when**: You find something wrong/outdated in memory, or code changes

**Read time**: 10 minutes (reference as needed)

**Token cost**: ~10K

---

## Quick Decision Matrix

| I Need To... | Read This | Ctrl+F Search |
|---|---|---|
| Understand the project | memory.md | "Directory Layout" |
| Write a component | memory.md | "Component Template" |
| Write a page | memory.md | "Page Template" |
| Write a test | memory.md | "Unit Tests" |
| Add translations | memory.md | "Key Structure" |
| Know the constraints | memory.md | "Never Violate" |
| Find a file | memory.md | "By Feature Area" |
| Find code | memory.md | "GREP PATTERNS" |
| Learn how to use memory | README.md | Start from top |
| Navigate the system | INDEX.md (this file) | Any section |

---

## Reading Paths by Role

### For @pm (Product Manager)
1. README.md (understand system)
2. memory.md SECTION 1 (project structure)
3. memory.md SECTION 10 (constraints)
4. memory.md SECTION 13 (decisions)

**Total**: ~20 minutes, ~5K tokens

### For @engineer (Full-Stack Developer)
1. README.md (understand system)
2. memory.md SECTION 1 (structure)
3. memory.md SECTION 2 (components)
4. memory.md SECTION 10 (constraints)
5. memory.md SECTION 12 (grep patterns)

**Total**: ~30 minutes, ~8K tokens
**Then per task**: ~5K tokens (grep + code)

### For @qa (QA Engineer)
1. README.md (understand system)
2. memory.md SECTION 6 (testing)
3. memory.md SECTION 10 (constraints)
4. memory.md SECTION 12 (grep patterns)

**Total**: ~20 minutes, ~5K tokens

### For @devops (DevOps Master)
1. README.md (understand system)
2. memory.md SECTION 1 (structure)
3. memory.md SECTION 9 (commands)
4. memory.md SECTION 10 (constraints)

**Total**: ~15 minutes, ~3K tokens

### For Generic Agent (First Time)
1. README.md (start here)
2. memory.md (read everything)
3. Bookmark SECTION 12 (grep patterns)
4. Bookmark SECTION 2, 6, 10 (your role's sections)

**Total**: ~45 minutes, ~15K tokens (one-time)

---

## Section Descriptions (memory.md)

**SECTION 1: PROJECT STRUCTURE** (Grep: "Directory Layout")
- Directory tree
- Tech stack
- Key dependencies

**SECTION 2: COMPONENT PATTERNS** (Grep: "Component Structure")
- Folder structure
- Template (copy this!)
- Key rules

**SECTION 3: PAGE PATTERNS** (Grep: "File-Based Routing")
- Route mapping
- Template (copy this!)
- Rules

**SECTION 4: COMPOSABLES** (Grep: "Naming Pattern")
- Naming convention
- Template (copy this!)
- Usage example

**SECTION 5: i18n** (Grep: "restructureDir")
- Config (critical!)
- File locations
- Key structure (copy this!)
- Rules (CRITICAL!)

**SECTION 6: TESTING** (Grep: "Unit Tests")
- Unit test template (copy this!)
- E2E test template (copy this!)
- Rules

**SECTION 7: STYLING** (Grep: "Global CSS")
- Global styles
- ThemeData pattern
- Flutter Material usage
- Rules

**SECTION 8: Dart** (Grep: "Strict Mode")
- Types in components
- Props & emits
- Rules

**SECTION 9: COMMANDS** (Grep: "Essential Commands")
- npm commands
- Flags
- First-time setup

**SECTION 10: CONSTRAINTS** (Grep: "Never Violate")
- 8 critical rules
- Red flags ⚠️
- What breaks when violated

**SECTION 11: FILE LOCATIONS** (Grep: "By Feature Area")
- Frequently modified
- Stable core
- By feature area

**SECTION 12: GREP PATTERNS** (Grep: "GREP PATTERNS")
- Copy-paste grep commands
- No interpretation needed

**SECTION 13: DECISION HISTORY** (Grep: "Recent Changes")
- What changed
- Why each tool
- Why NOT these

**SECTION 14: TOKEN TIPS** (Grep: "Token-Saving Tips")
- What to read
- What NOT to explore
- Token budgets

**SECTION 15: RESOURCES** (Grep: "CLAUDE.md")
- Related files
- External links
- Update history

---

## Workflow Example: Add Component

```
Agent needs to add a component

Step 1: Orient self (1 min)
├─ Read: This INDEX.md (5 min)
└─ Understand: memory.md organization

Step 2: Get pattern (2 min)
├─ Open: memory.md
├─ Search: "Component Template"
└─ Copy: Template code

Step 3: Understand rules (2 min)
├─ Search: "Key Rules" (in SECTION 2)
└─ Verify: Will follow these

Step 4: Find example (1 min)
├─ Memory says: Components in lib/features/[NAME]/
├─ Use grep: find lib/features -name "index.Flutter"
└─ Read: 1 existing component (50 lines)

Step 5: Write code (0 min)
├─ Use template from memory
└─ Follow rules from memory

Step 6: Verify (1 min)
├─ Check: All rules from SECTION 10 followed
└─ Test: flutter test:unit

Total tokens: ~5K (grep + read 1 example)
Without memory: ~80K (explore 30+ files)
Savings: 75K tokens
```

---

## Workflow Example: Fix Bug

```
Agent needs to fix a bug

Step 1: Locate bug (2 min)
├─ memory.md SECTION 11: By Feature Area
├─ Know where bug likely is
└─ Grep for: grep -r "bugKeyword" app/

Step 2: Understand constraints (1 min)
├─ memory.md SECTION 10: Never violate these
└─ Plan fix within constraints

Step 3: Check pattern (1 min)
├─ memory.md relevant section
└─ Know how to implement fix

Step 4: Implement fix (0 min)
├─ Write code following pattern
└─ No need to explore

Step 5: Validate (1 min)
├─ npm test
└─ git diff

Total tokens: ~8K (read sections + grep + read buggy file)
Without memory: ~50K (explore, understand bug context, find solution)
Savings: 42K tokens
```

---

## Search Strategy

### If you know the keyword:
```
Ctrl+F in memory.md → Find section → Read section
```

### If you know the task:
```
1. Check README.md "Quick Search Guide"
2. Find matching task
3. Go to suggested section
4. Search with suggested keyword
```

### If you don't know where to look:
```
1. Read this INDEX.md (3 min)
2. Pick your role's path
3. Follow suggested sections
4. Use Ctrl+F within those sections
```

---

## Update Flow

When something changes in the code:

```
1. Note what changed
2. Update memory.md relevant section
3. Update "UPDATE HISTORY" in memory.md
4. (Optional) Commit to git with message:
   "docs(memory): update [SECTION] for [change]"
```

---

## Token Budget Summary

| Task | Read Time | Token Cost | Savings vs. Explore |
|------|-----------|-----------|-------------------|
| First-time orientation | 45 min | 15K | 85K (85%) |
| Add component | 10 min | 5K | 75K (94%) |
| Fix bug | 5 min | 8K | 42K (84%) |
| Write test | 5 min | 4K | 26K (87%) |
| Add feature | 15 min | 10K | 90K (90%) |

**Average savings per task**: 80% of tokens

---

## Verification: Are You Using Memory Correctly?

✅ **Good Signs**:
- You started by reading README.md
- You searched memory.md with Ctrl+F
- You copied templates from memory.md
- You used grep patterns from SECTION 12
- Your code follows constraints from SECTION 10
- You haven't opened 10+ files

❌ **Bad Signs**:
- You're opening files one by one
- You're searching for patterns in codebase
- You don't know what the constraint section says
- You've read 30+ files already
- You're exploring instead of referencing memory

If you see bad signs: **Stop and go back to memory.md**

---

## Pro Tips

1. **Bookmark SECTION 12** (grep patterns) in memory.md
2. **Keep README.md open** (for reference)
3. **Use Ctrl+F** (it's your best friend)
4. **Copy templates exactly** (don't modify until you understand)
5. **Check constraints** (before writing any code)
6. **Update memory** (when you find it's wrong)

---

## Support

**Where do I find X?** → This INDEX.md → Quick Decision Matrix

**How do I use memory?** → README.md

**What's the pattern for Y?** → memory.md + Ctrl+F

**Is something wrong?** → Update memory.md

**Can I modify memory?** → Yes, please do if it's outdated

---

## File Locations in Repo

```
.agents/memory/
├── memory.md          (Main knowledge base)
├── README.md          (How to use)
└── INDEX.md          (This file)

Also related:
├── .agents/rules/      (Agent guidance)
├── .agents/skills/     (Searching techniques)
├── .agents/workflows/  (Processes)
└── CLAUDE.md          (Quick ref at root)
```

---

## Status

✅ **memory.md**: Complete, 675 lines, ready to use
✅ **README.md**: Complete, 380 lines, explains everything
✅ **INDEX.md**: Complete, navigation guide

**Last updated**: 2026-05-04
**Status**: Production-ready
**Next review**: When major code changes occur

---

## Quick Links (in memory.md, use Ctrl+F)

- "Directory Layout" → Project structure
- "Component Template" → How to write component
- "Page Template" → How to write page
- "Key Structure" → i18n setup
- "Unit Tests" → How to test
- "Never Violate These" → Constraints
- "By Feature Area" → File locations
- "GREP PATTERNS" → Find code efficiently
- "Token-Saving Tips" → Optimize for cost

---

🚀 **Ready to work efficiently?** Pick your role's path above and start reading!
