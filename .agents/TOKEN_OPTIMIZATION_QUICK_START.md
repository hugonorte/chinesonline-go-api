# ⚡ Token Optimization — Quick Start

**TL;DR**: Use cache, grep instead of read, follow patterns. This saves **70% of tokens**.

---

## 5-Minute Setup

### 1. Understand the Cache
```
.agents/memory/01-architecture.md    ← Directory structure & patterns
.agents/memory/02-conventions.md     ← Naming & Dart rules
.agents/memory/03-hot-paths.md      ← Which files to modify
.agents/memory/04-tech-decisions.md ← Why each tool, constraints
CLAUDE.md                             ← Quick reference (read this first!)
```

**Total read time**: 15 minutes  
**Token cost**: 15K (one time)

### 2. Know Your Native Search Tools
```
# Find component
Tool: grep_search (Pattern: "[NAME]")

# Find page
Tool: list_dir (Path: "lib/features/")

# Find composable
Tool: grep_search (Pattern: "use[NAME]", Path: "app/")

# Find usage
Tool: grep_search (Pattern: "MyComponent", Path: "app/")

# Find test
Tool: grep_search (Pattern: "[NAME]", Path: "app/", Includes: ["*.spec.ts"])
```

**Each search**: <1K tokens  
**Read only what's needed**: 1-5K tokens

### 3. Write Code (No Exploration)
```
Know the pattern from conventions.md? → Write it.
Need example? → Grep for 1 similar file → Read 50 lines → Write.
Not sure about something? → Check cache, not codebase.
```

**Token cost**: 0 (you're writing, not reading)

---

## The Matrix: What to Read When

| Need | Cache File | Grep Pattern | Token Cost |
|------|-----------|--------------|-----------|
| **"How do components work?"** | 02-conventions.md | - | 3K |
| **"Where is ProposalForm?"** | 03-hot-paths.md | `grep_search: Proposal` | 1K |
| **"How do I write Dart?"** | 02-conventions.md | - | 3K |
| **"What's the i18n setup?"** | 01-architecture.md | `grep_search: i18n in Flutter.config.ts` | 3K |
| **"What changed recently?"** | - | `git log --oneline -10` | 1K |
| **"Should I use this pattern?"** | 02-conventions.md | - | 3K |
| **"Is this file stable?"** | 03-hot-paths.md | - | 4K |
| **"I need to see an example"** | 04-tech-decisions.md | `grep_search: MyPattern in app/` | 2K |

---

## Role-Specific Quick Paths

### @pm (Product Manager)
```
1. Read: CLAUDE.md (3KB)
2. Read: tech_decisions.md constraints (2KB)
3. Write specification
4. Ask for approval
Total: 5K tokens
```

### @engineer (Full-Stack Developer)
```
1. Read: .agents/memory/02-conventions.md (3KB)
2. Read: .agents/memory/03-hot-paths.md (4KB)
3. Grep: 1 pattern example (1KB)
4. Write: code using pattern (0KB)
5. Run: tests (0KB)
Total: 8K tokens per task
```

### @qa (QA Engineer)
```
1. Read: .agents/memory/02-conventions.md testing (1KB)
2. Read: .agents/memory/03-hot-paths.md test paths (4KB)
3. Grep: existing tests (1KB)
4. Write: new tests (0KB)
5. Run: tests (0KB)
Total: 6K tokens per task
```

### @devops (DevOps Master)
```
1. Read: CLAUDE.md (3KB)
2. Read: tech_decisions.md build (2KB)
3. Run: npm commands (0KB)
4. Check: GitHub Actions (1KB)
Total: 6K tokens per deployment
```

---

## Common Mistakes = Token Waste

| ❌ Don't | 💸 Cost | ✅ Do Instead |
|---------|---------|--------------|
| Read entire codebase | 200K | Read cache (15K) |
| Explore similar files | 50K | Grep + read 1 (5K) |
| Read full component | 30K | Head/tail or grep (3K) |
| Ask without trying | - | Try cache first |
| Start without cache | 50K | Load cache (10K) |
| Search without grep | 100K | Grep first (1K) |

---

## Token Budget Example

**Task**: "Add i18n to PowerElectronicsModeling component"

```
Phase 1: Setup (5 min)
├─ Read .agents/memory/02-conventions.md: 3KB
├─ Read .agents/memory/03-hot-paths.md: 4KB
└─ Total: 7K tokens

Phase 2: Analysis (10 min)
├─ Grep for example: 1KB
├─ Check existing tests: 1KB
└─ Total: 2K tokens

Phase 3: Code (30 min)
├─ Write component: 0KB
├─ Write tests: 0KB
└─ Total: 0K tokens

Phase 4: Validation (10 min)
├─ Run tests: 0KB
├─ Git diff: 1KB
└─ Total: 1K tokens

TOTAL: 10K tokens
BUDGET: 30K tokens
SUCCESS: ✅ Well under budget
```

---

## The Rules (Golden)

1. **Cache first** — Always read cache before exploring code
2. **Grep second** — Search before reading large files
3. **Pattern match** — Know the pattern, write it, don't explore for examples
4. **Selective read** — Only read the file/section you're modifying
5. **Git history** — Use git log/blame instead of code analysis
6. **Validate fast** — Run tests, don't read test code to understand

---

## Checklist: Before Starting Any Task

- [ ] Have I read CLAUDE.md? (yes = 3K tokens saved)
- [ ] Do I know my role's cache files? (yes = 5K tokens saved)
- [ ] Can I grep instead of explore? (yes = 40K tokens saved)
- [ ] Do I know the pattern from conventions.md? (yes = 5K tokens saved)
- [ ] Have I read only the file I'm modifying? (yes = 50K tokens saved)
- [ ] Are tests passing? (yes = 0 tokens on debugging)

If you checked all boxes: **You're doing it right. ~20K tokens per task.**

---

## When You're Stuck

**Stuck?** Check this order (stop at first match):

1. ✅ Is the answer in cache? → Read cache (1K tokens)
2. ✅ Is there an example in the codebase? → Grep for it (1K tokens)
3. ✅ Does git history help? → `git log` the file (1K tokens)
4. ✅ Can I infer from test? → Read test file (2K tokens)
5. ❌ Read entire codebase (only if ALL above fail)

---

## Files You Should Know Exist

**Cache System** (read these):
```
CLAUDE.md                                    ← START HERE
.agents/memory/01-architecture.md     ← Structure
.agents/memory/02-conventions.md      ← Patterns
.agents/memory/03-hot-paths.md       ← File locations
.agents/memory/04-tech-decisions.md  ← Constraints
```

**Token Optimization Rules** (read before starting):
```
.agents/rules/project-context-cache.md       ← How cache works
.agents/rules/token-optimization-strategies.md ← Techniques
.agents/skills/search-code-efficiently.md    ← How to grep
.agents/workflows/token-efficient-work.md    ← Step-by-step
```

**Project Structure** (know, don't read):
```
lib/features/                              ← Components
lib/features/                                   ← Pages
lib/core/                             ← Logic
app/lang/                                    ← Translations
Flutter.config.ts                               ← Config
pubspec.yaml                                 ← Dependencies
```

---

## Success Looks Like

✅ **Tokens under 30K per task**  
✅ **Task completed correctly**  
✅ **Tests passing**  
✅ **No code exploration**  
✅ **Pattern-based implementation**

❌ **Tokens over 50K** = You're exploring too much  
❌ **Reading entire files** = Pattern not understood  
❌ **Tests failing** = Didn't follow conventions  

---

## Start Now

1. **Right now**: Read `CLAUDE.md` (3 minutes)
2. **For your role**: Read relevant cache file from .agents/memory/ (5 minutes)
3. **Pick a task**: Understand what needs doing
4. **Search pattern**: Grep for similar code
5. **Write code**: Use pattern, don't explore
6. **Run tests**: Validate
7. **Commit**: Done!

**Total time**: 1-2 hours  
**Total tokens**: ~20K  
**Result**: Task complete, budget happy

---

**Questions?**

- How do caches work? → `.agents/rules/project-context-cache.md`
- How to search efficiently? → `.agents/skills/search-code-efficiently.md`
- What's the exact workflow? → `.agents/workflows/token-efficient-work.md`
- What are constraints? → `.agents/memory/04-tech-decisions.md`

---

**Remember**: Every token you save is a token you can spend on actual work.

**Now go build something! 🚀**
