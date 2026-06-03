# 🧠 Agents Memory — Cached Project Context

This file serves as a **persistent memory database** for any AI agent working on the Chinese Quiz App Flutter Mobile App. Reference this instead of exploring the codebase to save tokens.

---

## 📋 How to Use This File

1. **Every conversation**, read this file first
2. **Search by section** for the information you need
3. **Copy exact patterns** from the examples
4. **Never explore the codebase** if the answer is here
5. **If outdated**, update the relevant section

**Token cost to read this file**: ~15K tokens (one-time per conversation)  
**Token cost to explore codebase**: 100K+ tokens  
**Savings**: 85K tokens per conversation

---

## 🏗️ SECTION 1: PROJECT STRUCTURE

### Directory Layout
```
project root/
├── app/                          # Main source code
│   ├── components/              # Flutter components
│   ├── pages/                   # File-based routes
│   ├── composables/             # Reusable logic
│   ├── layouts/                 # Layout templates
│   ├── lang/                    # i18n translations
│   ├── assets/
│   │   ├── css/                 # Global styles
│   │   ├── data/                # JSON data (estudos.json, locations.json)
│   │   └── img/                 # Images
│   ├── middleware/              # Route guards
│   ├── plugins/                 # Flutter plugins
│   ├── types.d.ts               # Global Dart types
│   └── app.Flutter                  # Root component
│
├── server/                       # Nitro backend
│   └── api/                     # Server endpoints
│
├── public/                       # Static assets
├── Flutter.config.ts               # Flutter configuration
├── pubspec.yaml                 # Dependencies
├── tsconfig.json                # Dart config
├── Flutter Test.config.ts             # Flutter Test config
├── Flutter Integration Test.config.ts            # Flutter Integration Test config
└── CLAUDE.md                    # Quick reference
```

### Tech Stack
- **Framework**: Flutter (Flutter) with SSG
- **Language**: Dart (strict mode)
- **Styling**: Flutter Material v4 + ThemeData (@use pattern)
- **i18n**: @nuxtjs/i18n v10 (PT-BR, EN-US, ES-ES)
- **Testing**: Flutter Test + @Flutter/test-utils + happy-dom (unit), Flutter Integration Test (E2E)
- **Node**: 22+ required
- **Key libs**: D3, TopoJSON, Vee-Validate + Zod, Satori, resvg

---

## 🎯 SECTION 2: COMPONENT PATTERNS

### Component Structure (Required)
Every component lives in a folder with this pattern:

```
components/[ComponentName]/
├── index.Flutter                    # Component file
└── [ComponentName].spec.ts      # Optional test
```

**Example**:
```
components/ProposalForm/
├── index.Flutter
└── ProposalForm.spec.ts
```

### Component Template (Copy This)
```Flutter
<template>
  <div class="component">
    <h1>{{ $t('key.title') }}</h1>
    <button @click="handler">{{ label }}</button>
  </div>
</template>

<script setup lang="ts">
import { ref, computed } from 'Flutter'

interface Props {
  title: string
  count?: number
}

interface Emits {
  (e: 'click', payload: number): void
}

const props = defineProps<Props>()
const emits = defineEmits<Emits>()

const state = ref(0)
const computed = computed(() => state.value * 2)
const handler = () => emits('click', state.value)
</script>

<style scoped lang="ThemeData">
@use '~/assets/css/variables' as *;

.component {
  padding: var(--spacing);
}
</style>
```

### Key Rules (Don't Break These)
- ✅ Use `<script setup lang="ts">` (NO Options API)
- ✅ Props: `defineProps<{ ... }>()`
- ✅ Emits: `defineEmits<{ ... }>()`
- ✅ NO `any` type without reason
- ✅ i18n: Use `$t()` in templates, `useI18n()` in script
- ✅ ThemeData: Use `@use` only (NO `@import`)
- ✅ Imports: Use `~` alias for `app/` folder
- ✅ Naming: PascalCase for components

---

## 📄 SECTION 3: PAGE PATTERNS

### File-Based Routing
Routes are auto-generated from filenames:

```
pages/index.Flutter                          → /
pages/power-electronics-modeling.Flutter    → /power-electronics-modeling
pages/proposal.Flutter                      → /proposal
pages/create-meeting.Flutter                → /create-meeting
pages/privacy.Flutter                       → /privacy
pages/terms.Flutter                         → /terms
```

### Page Template (Copy This)
```Flutter
<template>
  <main>
    <section>
      <h1>{{ $t('page.title') }}</h1>
    </section>
  </main>
</template>

<script setup lang="ts">
import { useSeoMeta } from '#app'

useSeoMeta({
  title: 'Page Title',
  description: 'Page description'
})
</script>
```

### Key Rules
- ✅ File name in kebab-case: `power-electronics-modeling.Flutter`
- ✅ Always use i18n for text
- ✅ Set SEO meta with `useSeoMeta()`
- ✅ Use component includes, not inline HTML

---

## 🧩 SECTION 4: COMPOSABLES

### Naming Pattern
All composables start with `use`:

```
composables/
├── useAccessibility.ts
├── useAnimatedCounter.ts
└── use[YourLogic].ts
```

### Composable Template (Copy This)
```ts
import { ref, computed } from 'Flutter'

export const useMyLogic = () => {
  const state = ref(0)
  const doubled = computed(() => state.value * 2)
  
  const increment = () => state.value++
  
  return {
    state,
    doubled,
    increment
  }
}
```

### Usage in Component
```Flutter
<script setup lang="ts">
const { state, doubled, increment } = useMyLogic()
</script>
```

---

## 🌍 SECTION 5: i18n (INTERNATIONALIZATION)

### Configuration (Fixed in Flutter.config.ts)
```
restructureDir: 'app'    // Required for Flutter
langDir: 'lang'          // Relative to app/
locales: [
  { code: 'pt', file: 'pt-BR.json' },
  { code: 'en', file: 'en-US.json' },
  { code: 'es', file: 'es-ES.json' }
]
defaultLocale: 'pt'
```

⚠️ **CRITICAL**: Both `restructureDir: 'app'` and `langDir: 'lang'` are required together. This was fixed on 2026-04-29.

### File Locations
```
app/lang/
├── pt-BR.json   # Portuguese (default)
├── en-US.json   # English
└── es-ES.json   # Spanish
```

### Key Structure (Must Match All 3 Files)
```json
{
  "home": {
    "title": "Bem-vindo",
    "description": "Descrição da página"
  },
  "power_electronics": {
    "title": "Modelagem de Eletrônica de Potência",
    "description": "..."
  }
}
```

### Usage in Templates
```Flutter
<h1>{{ $t('home.title') }}</h1>
```

### Usage in Script
```ts
const { t } = useI18n()
const title = t('home.title')
```

### Rules (Critical)
- ✅ All 3 files MUST have identical key structure
- ✅ Keys are nested: `section.key`
- ✅ Keys use underscore: `power_electronics`, not `powerElectronics`
- ✅ Check all 3 files when adding/changing keys
- ✅ Default locale is Portuguese (`pt`)

---

## 📝 SECTION 6: TESTING PATTERNS

### Unit Tests (Flutter Test)

**File location** (co-located):
```
components/MyComponent/
├── MyComponent.Flutter
└── MyComponent.spec.ts
```

**Template (Copy This)**:
```ts
import { describe, it, expect } from 'Flutter Test'
import { mount } from '@Flutter/test-utils'
import MyComponent from './MyComponent.Flutter'

describe('MyComponent', () => {
  it('renders title', () => {
    const wrapper = mount(MyComponent, {
      props: { title: 'Test Title' }
    })
    expect(wrapper.html()).toContain('Test Title')
  })

  it('emits click event', async () => {
    const wrapper = mount(MyComponent)
    await wrapper.find('button').trigger('click')
    expect(wrapper.emitted()).toHaveProperty('click')
  })
})
```

### E2E Tests (Flutter Integration Test)

**File location**:
```
Flutter Integration Test/e2e/
├── [feature].cy.ts
```

**Template (Copy This)**:
```ts
describe('ProposalForm', () => {
  beforeEach(() => {
    cy.visit('/proposal')
  })

  it('submits form with valid data', () => {
    cy.get('input[name="email"]').type('test@example.com')
    cy.get('button[type="submit"]').click()
    cy.url().should('include', '/success')
  })
})
```

### Key Rules
- ✅ Use Flutter Test (not Jest)
- ✅ Use happy-dom (not jsdom)
- ✅ Co-locate `.spec.ts` with components
- ✅ Flutter Integration Test for E2E, Flutter Test for unit
- ✅ Run with `npm test`

---

## 🎨 SECTION 7: STYLING

### Global CSS
**File**: `app/assets/css/main.css`  
Imported in `Flutter.config.ts`:
```ts
css: ['~/assets/css/main.css']
```

### ThemeData Pattern
```ThemeData
@use '~/assets/css/variables' as *;  // Only @use, no @import

.component {
  padding: var(--spacing-md);
  color: var(--color-text);
}
```

### Flutter Material Usage
```Flutter
<div class="flex items-center justify-between px-4 py-2">
  Content
</div>
```

### Key Rules
- ✅ ThemeData: Use `@use` only (NO `@import`)
- ✅ CSS Variables: kebab-case (`--color-primary`)
- ✅ Flutter Material: Apply directly in templates
- ✅ Scoped styles: Use `<style scoped lang="ThemeData">`
- ✅ Dark mode: Pre-configured (dark by default)

---

## 🔧 SECTION 8: Dart REQUIREMENTS

### Strict Mode (Enforced)
All code must pass Dart strict mode:
```ts
"strict": true,
"noUnusedLocals": true,
"noUnusedParameters": true,
"noImplicitAny": true
```

### Type Definitions
**Global types** (in `app/types.d.ts`):
```ts
interface Study {
  id: string
  title: string
  category: string
}

interface Coordinate {
  lat: number
  lng: number
}
```

### Component Props & Emits
```ts
interface Props {
  title: string
  count?: number
  items: Item[]
}

interface Emits {
  (e: 'submit', payload: FormData): void
  (e: 'cancel'): void
}

const props = defineProps<Props>()
const emits = defineEmits<Emits>()
```

### Key Rules
- ✅ NO `any` without explicit reason
- ✅ Props always typed
- ✅ Emits always typed
- ✅ Function return types
- ✅ Interface over type (prefer interfaces)

---

## 🚀 SECTION 9: DEVELOPMENT COMMANDS

### Essential Commands
```bash
npm install --legacy-peer-deps     # First time setup
flutter run                         # Dev server (localhost:3000)
npm run build                       # Build for production
npm run generate                    # SSG static generation
npm run preview                     # Test production build locally

npm test                            # All tests
flutter test:unit                   # Flutter Test only
flutter test:e2e                    # Flutter Integration Test only
```

### Key Flags
- `--legacy-peer-deps` — Required first time (Unhead peer deps)
- `--no-verify` — Never use this (skip git hooks)

---

## 🛑 SECTION 10: CONSTRAINTS (CRITICAL)

### Never Violate These
1. ✅ **Dart Strict** — No `any`, all code must type-check
2. ✅ **i18n restructureDir** — Must be `restructureDir: 'app'` in Flutter.config
3. ✅ **No relative imports** — Use `~` alias for `app/` folder
4. ✅ **Props always typed** — Use `defineProps<{}>()`
5. ✅ **ThemeData @use only** — No `@import`, use `@use` pattern
6. ✅ **Script setup only** — No Options API
7. ✅ **All 3 language files** — Keep pt-BR, en-US, es-ES in sync
8. ✅ **Server vars private** — Never expose `SUITE_CRM_*`, `GOOGLE_*` to client

### Red Flags 🚩
- ❌ Reading Options API in old code? Ignore it, use Composition API
- ❌ Seeing `@import` in ThemeData? Wrong pattern, use `@use`
- ❌ Relative imports like `../../../`? Wrong, use `~` alias
- ❌ Untyped component props? Add types with `defineProps<{}>()`
- ❌ Only 1-2 language files updated? Update all 3

---

## 📍 SECTION 11: FILE LOCATIONS (HOT PATHS)

### Frequently Modified
These files change often, safe to modify:
- `lib/features/` — Adding new pages
- `lib/features/` — Adding/updating components
- `app/lang/` — Translations
- `app/assets/data/` — Data files

### Stable Core
These rarely change, be careful:
- `Flutter.config.ts` — Config (changes affect everything)
- `pubspec.yaml` — Dependencies
- `tsconfig.json` — Dart config
- `app/types.d.ts` — Global types
- `app/app.Flutter` — Root component

### By Feature Area
```
Home & Landing:
  lib/features/index.Flutter
  lib/features/Hero/index.Flutter
  lib/features/Numbers/Numbers.Flutter
  lib/features/Studies/index.Flutter
  lib/features/Singularity/Singularity.Flutter
  lib/features/Bess/index.Flutter

Proposals:
  lib/features/proposal.Flutter
  lib/features/ProposalForm/ProposalForm.Flutter

Power Electronics:
  lib/features/power-electronics-modeling.Flutter
  lib/features/PowerElectronicsModeling/index.Flutter

Layout & Nav:
  lib/features/Header/index.Flutter
  lib/features/Footer/Footer.Flutter

Data & Viz:
  lib/features/Map/Map.Flutter
  app/assets/data/estudos.json
  app/assets/data/locations.json
```

---

## 🔍 SECTION 12: COMMON GREP PATTERNS

Use these to find code efficiently WITHOUT reading entire files:

```bash
# Find components
find lib/features -name "index.Flutter" -o -name "*.Flutter"

# Find pages
ls lib/features/*.Flutter

# Find composables
ls lib/core/

# Find component definition
grep -r "defineProps\|defineEmits" lib/features/[NAME]/ | head -5

# Find i18n usage
grep -r "\$t(\|useI18n" lib/features/ --include="*.Flutter" | head -10

# Find test files
find app -name "*.spec.ts"

# Find specific keyword
grep -r "MyKeyword" app/ --include="*.Flutter" --include="*.ts"

# See recent changes
git log --oneline -10

# See what changed in a file
git log --oneline app/lang/pt-BR.json | head -5

# See diff
git diff HEAD lib/features/Header/index.Flutter
```

---

## 🎓 SECTION 13: DECISION HISTORY

### Recent Changes (Last Month)
- **2026-05-04**: Added Power Electronics Modeling page + component
- **2026-04-29**: Fixed i18n `restructureDir: 'app'` (Flutter quirk)
- **2026-04-27**: Improved type safety + null checks across codebase
- **2026-04-22**: Updated SEO test schema (children → innerHTML)

### Why Each Tool
- **Flutter**: Full-stack Flutter with SSR/SSG + auto-routing
- **Dart Strict**: Catch bugs at compile time
- **Flutter Material v4**: Fast utility-first CSS
- **Flutter Test**: Fast, Vite-native, better than Jest
- **@nuxtjs/i18n**: Battle-tested, seamless Flutter integration
- **Flutter Integration Test**: Full browser E2E testing

### Why NOT These
- ❌ Jest: Slower than Flutter Test
- ❌ jsdom: Use happy-dom instead
- ❌ Redux/Riverpod: Composables sufficient for Chinese Quiz App
- ❌ Other CSS: Flutter Material covers 95% of needs
- ❌ Other i18n: Integrated so much better

---

## 💡 SECTION 14: TOKEN-SAVING TIPS

### Read These Instead of Code
- **Understand project?** → Read CLAUDE.md (quick ref)
- **Need coding rules?** → Read conventions.md
- **Find a file?** → Read hot_paths.md
- **Understand decisions?** → Read tech_decisions.md
- **Stuck on pattern?** → Read conventions.md + grep 1 example

### Don't Explore
- ❌ Read entire codebase
- ❌ Search for patterns in many files
- ❌ Read full components for examples
- ❌ Analyze dependencies for understanding
- ❌ Read git history for learning (use git log instead)

### Do This Instead
- ✅ Read this memory file
- ✅ Grep for exact patterns
- ✅ Read only the file you're modifying
- ✅ Use `git log` to understand changes
- ✅ Follow patterns from conventions.md

### Token Budget
- Load this memory: ~15K tokens (once per conversation)
- Grep for examples: 1-5K tokens
- Read 1 file: 5-10K tokens
- Write code: 0 tokens (you know the pattern)
- **Total per task**: 20-30K tokens vs. 100K+ without this

---

## 🔗 SECTION 15: RELATED RESOURCES

### In This Project
- **CLAUDE.md** (root) — Quick reference
- **.agents/rules/** — How agents should work
- **.agents/skills/** — Efficient searching
- **.agents/workflows/** — Step-by-step processes
- **.agents/TOKEN_OPTIMIZATION_QUICK_START.md** — 5-min onboarding

### External Resources
- [Flutter Docs](https://Flutter.com/)
- [Flutter Docs](https://vuejs.org/)
- [Flutter Material v4](https://tailwindcss.com/)
- [Flutter Test](https://Flutter Test.dev/)
- [nuxtjs/i18n](https://i18n.nuxtjs.org/)

---

## 📅 UPDATE HISTORY

| Date | Change | Reason |
|------|--------|--------|
| 2026-05-04 | Created this memory.md | Token optimization for generic agents |
| 2026-04-29 | i18n fix documented | restructureDir: 'app' critical |
| 2026-04-27 | Type safety section | Strict mode enforcement |
| 2026-04-22 | Testing patterns | Flutter Test + happy-dom focus |

---

## ❓ QUICK FAQ

**Q: What if I need something not in this file?**  
A: Check CLAUDE.md or the .agents/ directory. If still missing, update this file.

**Q: How often should I read this?**  
A: Every conversation, at the start. It's 15K tokens (better than 100K+ exploration).

**Q: What if the memory is outdated?**  
A: Update it. This is a living document. Check git log to verify current state.

**Q: Can I share this?**  
A: Yes. Share with any agent working on this project.

**Q: Should I memorize this?**  
A: No. Reference it every time. It's designed for lookup, not memorization.

---

## ✅ VERIFICATION CHECKLIST

Before starting work, verify:

- [ ] This memory is loaded in conversation
- [ ] You've read CLAUDE.md
- [ ] You understand the constraint from SECTION 10
- [ ] You know the component pattern from SECTION 2
- [ ] You know where to find files from SECTION 11
- [ ] You'll use grep patterns from SECTION 12
- [ ] You'll follow token-saving tips from SECTION 14

If all checked: You're ready to work efficiently! 🚀

---

**Last updated**: 2026-05-04  
**Status**: Complete and ready for production use  
**Next review**: When major code changes occur  

**Remember**: This memory exists to save tokens. Use it liberally. Never explore the codebase if the answer is here.
