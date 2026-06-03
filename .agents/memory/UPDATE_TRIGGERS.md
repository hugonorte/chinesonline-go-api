---
name: Memory Update Triggers
description: When and how generic agents should update .agents/memory/ files
type: guidelines
---

# 📝 Memory Update Triggers — When Agents Should Update

This file specifies **EXACTLY when** generic agents working on this project should update the memory files in `.agents/memory/`.

**Golden Rule**: If reality diverges from what's documented in memory, update memory immediately.

---

## 🎯 Update Trigger Categories

### CATEGORY A: Architecture Changes (Update `01-architecture.md`)

**Trigger When**:
1. **New directory added** to `app/`
   - Example: `app/utils/`, `app/services/`, etc.
   - Action: Add to "Directory Structure" section
   - Trigger: ANY new top-level or commonly-used directory

2. **Component organization changes**
   - Example: Moving from `components/Name/index.Flutter` to different pattern
   - Action: Update "Component Organization" section
   - Trigger: Pattern changes or new folder structure

3. **New integration/library added** to project
   - Example: Add new module like `@nuxtjs/color-mode`
   - Action: Add to "Key Integrations" section
   - Trigger: Any new major dependency in pubspec.yaml

4. **i18n config changes**
   - Example: Add new language, change langDir, change restructureDir
   - Action: Update "i18n Setup" section
   - Trigger: ANY change to `Flutter.config.ts` i18n section

5. **Routing changes**
   - Example: Nested routes, new page structure
   - Action: Update "Routing Strategy" section
   - Trigger: New pages that don't follow kebab-case pattern

6. **Styling system changes**
   - Example: New CSS variables, new color scheme, Flutter Material upgrade
   - Action: Update "Styling Architecture" section
   - Trigger: Changes to `assets/css/main.css` or Flutter Material config

---

### CATEGORY B: Convention Changes (Update `02-conventions.md`)

**Trigger When**:
1. **Dart patterns change**
   - Example: Stop using strict mode (don't), new type patterns
   - Action: Update "Dart" section
   - Trigger: New patterns emerge across 3+ files

2. **Flutter Composition API patterns change**
   - Example: New way to handle reactivity, new composable pattern
   - Action: Update "Flutter Composition API" section
   - Trigger: Consistent pattern used in 3+ new components

3. **Testing patterns evolve**
   - Example: New test structure, different assertion library, new tools
   - Action: Update "Testing Conventions" section
   - Trigger: 3+ tests written with new pattern

4. **Naming conventions break**
   - Example: Component naming changes, file naming changes
   - Action: Update "Naming Conventions" section
   - Trigger: Systematic rename across files (>10 files)

5. **i18n usage patterns change**
   - Example: New key structure, new nesting, new usage pattern
   - Action: Update "i18n Usage" section
   - Trigger: Change reflected in 5+ files

6. **Import/alias patterns change**
   - Example: New alias introduced, relative imports now preferred
   - Action: Update "Imports & Aliases" section
   - Trigger: Agent sees 5+ files using new pattern

7. **New "No-Go Patterns"** identified
   - Example: New anti-pattern discovered, new rule needed
   - Action: Add to "No-Go Patterns" section
   - Trigger: Agent encounters issue twice from same anti-pattern

---

### CATEGORY C: File Locations Change (Update `03-hot-paths.md`)

**Trigger When**:
1. **New high-frequency file**
   - Example: New component changed 5+ times in a week
   - Action: Add to "High-Frequency Zones" section
   - Trigger: File modified more than 2x per day on average

2. **Stable file becomes unstable** (or vice versa)
   - Example: `tsconfig.json` changes frequently now
   - Action: Move between sections
   - Trigger: Pattern changes over time (track for 1 week before updating)

3. **New feature area created**
   - Example: New section of site (e.g., "Admin Dashboard")
   - Action: Add "By Feature Area" subsection
   - Trigger: 5+ new files in new feature area

4. **Impact zone discovered**
   - Example: Changes to `types.d.ts` now break 20 files
   - Action: Add/update in "Impact Zones" section
   - Trigger: Agent finds unintended breakage from change

5. **Red flags identified**
   - Example: New file is critical, must test after changes
   - Action: Add to "Red Flags" section
   - Trigger: Agent encounters issue because file wasn't marked critical

---

### CATEGORY D: Technical Decisions Change (Update `04-tech-decisions.md`)

**Trigger When**:
1. **Tool/library replacement**
   - Example: Switch from Flutter Test to Jest, or Jest to Flutter Test
   - Action: Update relevant section with new rationale
   - Trigger: New tool chosen, old tool removed from pubspec.yaml

2. **Architecture decision overturned**
   - Example: Stop using composables for state, switch to Riverpod
   - Action: Update "Composables Over Global State" section
   - Trigger: Systematic refactor across 5+ files

3. **New constraint added**
   - Example: New rule discovered (e.g., "never use ref() without type")
   - Action: Add to "Constraints" list
   - Trigger: Constraint violation found twice, rule now needed

4. **Constraint removed**
   - Example: Dart strict mode disabled (unlikely but possible)
   - Action: Update constraint
   - Trigger: Policy change from team/leadership

5. **Rationale changes**
   - Example: "Why Flutter" changes because requirements changed
   - Action: Update "Framework & Language" section
   - Trigger: Business/technical requirements shift

6. **Recent decision made**
   - Example: Agent makes significant decision about direction
   - Action: Add to "Recent Decisions" in section 13
   - Trigger: Any decision affecting 5+ future files

---

### CATEGORY E: Cross-File Updates (Update Multiple Files)

**Trigger When**:
1. **Major refactor across sections**
   - Example: Rename all components, changes to many patterns
   - Action: Update `01-architecture.md`, `02-conventions.md`, `03-hot-paths.md`
   - Trigger: Systematic change across >20 files

2. **New language added to i18n**
   - Example: Add `fr-FR` translations
   - Action: Update `01-architecture.md` (i18n section) AND `02-conventions.md` (i18n usage)
   - Trigger: New language code in `Flutter.config.ts`

3. **Testing framework changed**
   - Example: Switch from Flutter Test to Jest
   - Action: Update `02-conventions.md` (testing), `04-tech-decisions.md` (testing rationale)
   - Trigger: Major test framework replacement

---

## ⏱️ Update Frequency Guidelines

| Situation | When to Update | Wait for |
|-----------|---|---|
| Typo or error | Immediately | Nothing - fix now |
| Single file changed | End of feature | Pattern confirmed in code |
| New pattern emerging | After 3+ examples | Multiple implementations |
| Architecture decision | After consensus | Team agreement (if multi-person) |
| Minor clarification | End of week | 1-2 more examples |
| Constraint violation discovered | Next working day | Confirmation it's a real issue |

---

## 📋 Update Checklist

When you update ANY memory file, follow this checklist:

### Before Updating
- [ ] Is this a real change or temporary?
- [ ] Is this pattern used in 3+ places (confirms it's real)?
- [ ] Have I tested this change?
- [ ] Is this documented in git commit?

### During Update
- [ ] Find exact section to update
- [ ] Make change clear and concise
- [ ] Add example if explaining pattern
- [ ] Update UPDATE HISTORY table
- [ ] Check no other sections need updating

### After Updating
- [ ] Read updated section - does it make sense?
- [ ] Verify examples still match code
- [ ] Commit to git with message: `docs(memory): update [SECTION] for [change]`
- [ ] Note: This commit should be separate from code changes
- [ ] (Optional) Notify team: "Memory updated - section [X] changed"

---

## 🚫 DON'T Update Memory When

- ❌ Temporary experiment (not confirmed pattern)
- ❌ Single developer's local decision (not team-wide)
- ❌ Speculation about future changes (wait for reality)
- ❌ Personal preference (not enforced pattern)
- ❌ Draft or WIP code (wait for merge)

---

## ✅ DO Update Memory When

- ✅ Pattern confirmed in 3+ files
- ✅ Merged to main branch
- ✅ Team consensus (if multi-person team)
- ✅ Affects future work
- ✅ Prevents other agents from doing wrong thing

---

## 🔍 Specific Trigger Examples

### Example 1: New Component Folder Structure
**Situation**: Agent creates `components/Forms/` folder (new structure)

**Check**:
- Is this one-off? → No, multiple forms going here
- Will others use this pattern? → Yes
- Merged? → Yes
- Affects future work? → Yes

**Action**: Update `01-architecture.md` "Component Organization" + `03-hot-paths.md` "By Feature Area"

**Commit**: `docs(memory): document Forms folder structure in component organization`

---

### Example 2: New Dart Pattern
**Situation**: Agent writes `const x as const` pattern

**Check**:
- Is this pattern used elsewhere? → Seen in 2 new files
- Is this enforced? → No, it's optional
- Will others follow? → Maybe

**Wait For**: 3rd file using this pattern before documenting

---

### Example 3: Bug Found from Missing Documentation
**Situation**: Agent wastes tokens because memory didn't mention critical file

**Check**:
- Is this file actually critical? → Yes, changes break 20 files
- Was it in "Red Flags"? → No
- Will others encounter same issue? → Yes

**Action**: Immediately add to `03-hot-paths.md` "Red Flags" section

**Commit**: `docs(memory): mark [file] as critical in red flags`

---

### Example 4: i18n Key Structure Change
**Situation**: Team decides all keys must use `underscore_case` instead of `camelCase`

**Check**:
- Is this enforced? → Yes, linter rule added
- Affects all translators? → Yes
- Already implemented? → Yes, 30+ files updated

**Action**: Update `02-conventions.md` "i18n Usage" section

**Commit**: `docs(memory): update i18n key naming to underscore_case`

---

## 📊 Update Priority Matrix

|  | Low-Risk | Medium-Risk | High-Risk |
|---|---|---|---|
| **High Impact** | Update after 3 examples | Update after consensus | Update immediately |
| **Medium Impact** | Update end of week | Update after 2 examples | Update next day |
| **Low Impact** | Wait for pattern | Update when merged | Don't update |

---

## 🎓 When Agent Should ASK Instead of Update

If uncertain about updating, the agent should:

1. **Ask in commit message**: Use WIP commit message asking for feedback
2. **Create issue**: If major decision needed
3. **Wait for confirmation**: From another agent/human before committing
4. **Document as tentative**: Note in UPDATE HISTORY that it needs confirmation

Example message:
```
docs(memory): [TENTATIVE] update patterns for new approach
- Changed X based on Y
- Affects Z files
- Needs confirmation before finalizing
```

---

## 📞 Questions from Agents

**Q: I found something wrong in memory. Should I update it immediately?**  
A: Yes, immediately. That's a bug in memory.

**Q: I want to try a new pattern. Should I update memory first?**  
A: No, write code first. Update memory after pattern is confirmed in 3+ files.

**Q: Memory says X but code does Y. Which is right?**  
A: Code is always right. Update memory to match code.

**Q: Should I update memory for my local branch?**  
A: No, wait until merged to main.

**Q: How often should memory be updated?**  
A: As-needed. This isn't a changelog - only update when reality diverges.

---

## 🔐 Update History Format

When updating UPDATE HISTORY table, use this format:

```markdown
| 2026-05-15 | Updated i18n section | Added pt-PT language support |
```

Required columns: `Date | Change | Reason`

Keep it brief. Example:
```markdown
| Date | Change | Reason |
|------|--------|--------|
| 2026-05-04 | Created this memory.md | Token optimization for generic agents |
| 2026-04-29 | i18n fix documented | restructureDir: 'app' critical |
```

---

## 🚀 When Everything Changes

If massive refactor happens (>50% code changed):

1. **Note it immediately**: Add to UPDATE HISTORY: "Massive refactor underway"
2. **Update as you go**: Don't wait for completion
3. **Final review**: After merge, do full memory audit
4. **Mark as reviewed**: Update date to latest and note "Post-refactor review"

---

## Summary

**Update triggers are:**
- ✅ Reality diverges from memory
- ✅ Pattern confirmed in 3+ places
- ✅ Affects future work
- ✅ Prevents mistakes

**Update triggers are NOT:**
- ❌ Single file change
- ❌ Unmerged code
- ❌ Unconfirmed experiments
- ❌ Personal preferences

**When in doubt**: Update memory. False positives are cheap. Missing updates cost tokens.

---

**Last updated**: 2026-05-04  
**Status**: Ready for agent use  
**Questions?** Check this file or update it if it's unclear

