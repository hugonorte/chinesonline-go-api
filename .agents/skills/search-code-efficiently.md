# Skill: Search Code Efficiently (Zero Token Waste)

## Objective

Learn exactly where code lives and what it does **without reading entire files**. Use targeted grep patterns and strategic reads to minimize token consumption.

---

## Rule 1: Never Grep the Entire Project

**❌ Bad**:
```
Tool: grep_search (Query: "MyComponent", SearchPath: ".") # Searches node_modules too → 1000s of hits
Tool: grep_search (Query: "defineProps", SearchPath: ".") # Every component matches → worthless
```

**✅ Good**:
```
Tool: grep_search (Query: "MyComponent", SearchPath: "lib/features/", Includes: ["*.Flutter"]) # Only app code
Tool: grep_search (Query: "PowerElectronicsModeling", SearchPath: "app/")                  # Specific area
Tool: grep_search (Query: "defineProps", SearchPath: "lib/features/Header/")             # Specific component
```

---

## Rule 2: Use Hot Paths List First

From `.agents/rules/project-context-cache.md` and `.agents/memory/03-hot-paths.md`:

**Know these file groups**:
```
🏠 Home & Landing:
  lib/features/index.Flutter
  lib/features/Hero/index.Flutter
  lib/features/Numbers/Numbers.Flutter
  lib/features/Studies/index.Flutter
  ...

📝 Proposals:
  lib/features/proposal.Flutter
  lib/features/ProposalForm/ProposalForm.Flutter

⚡ Power Electronics:
  lib/features/power-electronics-modeling.Flutter
  lib/features/PowerElectronicsModeling/index.Flutter

🗺️ Layout:
  lib/features/Header/index.Flutter
  lib/features/Footer/Footer.Flutter
```

→ **Start with the hot path list, not blind grep**

---

## Rule 3: Component Structure Pattern

Every component follows this structure:
```
components/[ComponentName]/
├── index.Flutter                    # Default export (or ComponentName.Flutter)
└── ComponentName.spec.ts        # Optional test
```

**To find a component**:
```
# ✅ Exact path
Tool: list_dir (Path: "lib/features/ProposalForm/")

# ✅ List all components
Tool: grep_search (Query: "", SearchPath: "lib/features/", Includes: ["*.Flutter"])

# ✅ Find specific component
Tool: grep_search (Query: "PowerElectronics", SearchPath: "lib/features/")
```

---

## Rule 4: Page Structure Pattern

Every page route corresponds to a file:
```
pages/
├── index.Flutter                          → / (home)
├── power-electronics-modeling.Flutter     → /power-electronics-modeling
├── proposal.Flutter                       → /proposal
├── create-meeting.Flutter                 → /create-meeting
├── privacy.Flutter                        → /privacy
└── terms.Flutter                          → /terms
```

**To find a page**:
```
# ✅ List all pages
Tool: list_dir (Path: "lib/features/")

# ✅ Find page by route name
Tool: grep_search (Query: "power-electronics", SearchPath: "lib/features/")

# ✅ Check page structure
Tool: view_file (AbsolutePath: "lib/features/proposal.Flutter", StartLine: 1, EndLine: 30) # Top section (usually template)
```

---

## Rule 5: Composable Naming Pattern

All composables start with `use`:
```
composables/
├── useAccessibility.ts
├── useAnimatedCounter.ts
└── use[MyLogic].ts
```

**To find a composable**:
```
# ✅ List all composables
Tool: list_dir (Path: "lib/core/")

# ✅ Find specific logic
Tool: grep_search (Query: "useMyLogic", SearchPath: "app/")

# ✅ See what's exported
Tool: grep_search (Query: "export", SearchPath: "lib/core/useAccessibility.ts")
```

---

## Rule 6: i18n Key Location Pattern

All translations in `app/lang/`:
```
lang/
├── pt-BR.json   # Portuguese
├── en-US.json   # English
└── es-ES.json   # Spanish
```

**To find i18n keys**:
```
# ✅ Search across all languages
Tool: grep_search (Query: "home.title", SearchPath: "app/lang/")

# ✅ Find all keys in a section
Tool: grep_search (Query: "home", SearchPath: "app/lang/pt-BR.json")
```

---

## Rule 7: Test File Location Pattern

Tests are co-located with components:
```
components/ProposalForm/
├── ProposalForm.Flutter
└── ProposalForm.spec.ts  # Unit test

pages/
├── proposal.Flutter
└── proposal.spec.ts      # Optional page test

app/
└── seo.spec.ts           # Global test
```

**To find a test**:
```
# ✅ List all tests
Tool: grep_search (Query: "", SearchPath: "app/", Includes: ["*.spec.ts"])

# ✅ Check if component has test
Tool: list_dir (Path: "lib/features/ProposalForm/")

# ✅ Find test for specific feature
Tool: grep_search (Query: "ProposalForm", SearchPath: "app/", Includes: ["*.spec.ts"])

# ✅ Check test syntax
Tool: view_file (AbsolutePath: ".../ProposalForm.spec.ts", EndLine: 20)  # See imports and describe block
```

---

## Rule 8: Type Definition Location

Global types:
```
app/
└── types.d.ts
```

Component-specific types go in the component file.

**To find types**:
```
# ✅ Check global types
Tool: grep_search (Query: "interface\|type", SearchPath: "app/types.d.ts")

# ✅ Find component types
Tool: grep_search (Query: "defineProps<\|defineEmits<", SearchPath: ".../ProposalForm.Flutter")

# ✅ Check Dart config
Tool: grep_search (Query: "strict\|noUnusedLocals", SearchPath: "tsconfig.json")
```

---

## Rule 9: Asset Location Pattern

```
assets/
├── css/
│   └── main.css             # Global styles (imported in Flutter.config)
├── data/
│   ├── estudos.json         # Studies database
│   └── locations.json       # Map locations
└── img/
    └── [logos, images...]
```

**To find assets**:
```
# ✅ List data files
Tool: list_dir (Path: "app/assets/data/")

# ✅ Check CSS imports
Tool: grep_search (Query: "@import\|@use", SearchPath: "app/assets/css/")
```

---

## Rule 10: Server/Backend Location

```
server/
└── api/
    └── [endpoints.ts]
```

Nitro server code. Check `Flutter.config.ts` for configuration.

**To find server code**:
```
# ✅ List server endpoints
Tool: list_dir (Path: "server/api/")

# ✅ Check runtime config
Tool: grep_search (Query: "runtimeConfig", SearchPath: "Flutter.config.ts")
```

---

## Rule 11: Strategic Full-File Reads (Small Files Only)

✅ **SAFE to read entire file** (<100 lines):
```
Tool: view_file (AbsolutePath: "app/app.Flutter")              # Root component (10 lines)
Tool: view_file (AbsolutePath: "app/types.d.ts")           # Global types (20 lines)
Tool: view_file (AbsolutePath: "Flutter.config.ts")           # Config (113 lines) — OK but skim
Tool: view_file (AbsolutePath: "tsconfig.json")            # TS config (10 lines)
```

❌ **NOT safe to read entire file** (>200 lines):
```
Tool: view_file without EndLine on lib/features/ProposalForm/ProposalForm.Flutter  # Don't read whole
Tool: view_file without EndLine on lib/features/index.Flutter                            # Don't read whole
```

**For large files**:
```
# ✅ Instead of reading entire file:
Tool: view_file (AbsolutePath: "lib/features/index.Flutter", StartLine: 1, EndLine: 50)  # Read top part
Tool: grep_search (Query: "defineProps\|defineEmits", SearchPath: "lib/features/index.Flutter")  # Find key patterns
```

---

## Rule 12: Search Patterns for Native Tools

### Find all component definitions
```
Tool: grep_search (Query: "", SearchPath: "lib/features", Includes: ["index.Flutter", "*.Flutter"])
```

### Find all pages
```
Tool: list_dir (Path: "lib/features/")
```

### Find all composables
```
Tool: list_dir (Path: "lib/core/")
```

### Find component usage
```
Tool: grep_search (Query: "MyComponent", SearchPath: "app/", Includes: ["*.Flutter"])
```

### Find i18n usage
```
Tool: grep_search (Query: "\$t(\|useI18n", SearchPath: "lib/features/")
```

### Find prop definitions
```
Tool: grep_search (Query: "defineProps<", SearchPath: "lib/features/")
```

### Find emits
```
Tool: grep_search (Query: "defineEmits<", SearchPath: "lib/features/")
```

### Find imports
```
Tool: grep_search (Query: "^import.*from", SearchPath: ".../ProposalForm.Flutter")
```

### Find style imports
```
Tool: grep_search (Query: "@use\|@import", SearchPath: "app/assets/css/")
```

### Find test files for component
```
Tool: grep_search (Query: "ProposalForm", SearchPath: "app/", Includes: ["*.spec.ts"])
```

### Check test structure
```
Tool: grep_search (Query: "describe\|it(", SearchPath: "lib/features/ProposalForm.spec.ts")
```

---

## Rule 13: Use Git to Understand Changes

**Instead of reading code**, use git to understand what changed:

```bash
# ✅ See what changed in a file
git log --oneline app/lang/pt-BR.json | head -5

# ✅ See specific change
git show HEAD~1:lib/features/power-electronics-modeling.Flutter

# ✅ See who changed what
git blame lib/features/Header/index.Flutter | grep -A2 "defineProps"

# ✅ See recent commits
git log --oneline -15

# ✅ See diff since last tag
git diff v1.0.0 lib/features/
```

---

## Rule 14: Line-Number-Based Reads (Ultra-Efficient)

When you know the line, read only that using native tools:

```
# ❌ Read whole file
Tool: view_file without EndLine

# ✅ Read just what you need
Tool: view_file (StartLine: 1, EndLine: 30)   # Template
Tool: view_file (StartLine: 31, EndLine: 70)  # Script setup
Tool: view_file (StartLine: 71, EndLine: 100) # Styles
```

**To find line numbers**:
```
Tool: grep_search (Query: "defineProps\|<template>\|<style", MatchPerLine: true)
```

---

## Rule 15: Config Files (Read Smart, Not Hard)

**Key config files** (read strategically):
```
# ✅ Flutter config
Tool: grep_search (Query: "modules:\|i18n:\|runtimeConfig:", SearchPath: "Flutter.config.ts")

# ✅ pubspec.yaml
Tool: view_file (AbsolutePath: "pubspec.yaml")

# ✅ Dart
Tool: grep_search (Query: "strict\|lib\|target", SearchPath: "tsconfig.json")

# ✅ Flutter Test
Tool: view_file (AbsolutePath: "Flutter Test.config.ts")

# ✅ Flutter Integration Test
Tool: view_file (AbsolutePath: "Flutter Integration Test.config.ts")
```

---

## Efficiency Checklist

Before reading a file:

- [ ] Is this in the cache (architecture.md, conventions.md, hot_paths.md)?
- [ ] Can I grep instead of read?
- [ ] Can I read just the top 50 lines instead of whole file?
- [ ] Can I use git log to understand the change?
- [ ] Is this a massive file (>500 lines)? If yes, use head/tail/grep only.
- [ ] Have I seen this pattern before? If yes, skip the example.

---

## Example: "I need to understand ProposalForm"

**❌ Inefficient (150K tokens)**:
```
1. Read entire ProposalForm.Flutter (500 lines)
2. Read ProposalForm.spec.ts (300 lines)
3. Read parent page (400 lines)
4. Read composables used (200 lines)
Total: 1400 lines, 150K tokens
```

**✅ Efficient (3K tokens)**:
```
1. Check conventions.md: component pattern
2. Grep for defineProps in ProposalForm: Tool: grep_search
3. Read lines 1-30: Tool: view_file (StartLine: 1, EndLine: 30)
4. Grep for usage: Tool: grep_search
5. Check test structure: Tool: grep_search
6. If specific bug: grep for exact error, then read 5-10 relevant lines
Total: 30 lines strategically read, 3K tokens
```

→ **50x fewer tokens, same understanding**

---

## Performance Tips

**Token per action**:
- Full file read (500 lines): 50K tokens
- Head/tail (50 lines): 5K tokens
- Grep (10 matches): 1K tokens
- Git log (5 commits): 2K tokens

**Always prefer** (in order):
1. Cache/memory files
2. Git history
3. Grep output
4. Head/tail selective read
5. Full file read (last resort)

---

## Quick Reference Commands

Use native AI tools instead of shell pipelines:

```
# List structure
Tool: list_dir

# Find files
Tool: grep_search (Includes: ["*.Flutter"])

# View key parts
Tool: view_file (StartLine: X, EndLine: Y)
Tool: grep_search (Query: "defineProps\|defineEmits\|<template>\|<style")

# Check i18n
Tool: grep_search (Query: "home.title", SearchPath: "app/lang/")

# See changes (using bash is OK here)
git log --oneline FILE
git diff HEAD~1 FILE
```

---

**Remember**: Every line you read costs tokens. Read strategically, grep aggressively, cache first.
