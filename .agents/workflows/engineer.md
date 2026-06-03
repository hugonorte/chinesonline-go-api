---
description: Orchestrate the implementation phase of the development cycle for Mobile App tasks.
---

# Workflow: Engineer Implementation

When acting as the **Full-Stack Engineer**, follow this sequence to translate the approved **Technical Specification** into production-ready Mobile App code.

### 1. Analysis and Setup
- **Read Specification**: Open `production_artifacts/Technical_Specification.md` and identify all required views, components, and data structures.
- **Review Conventions**: Re-read `.agents/rules/Flutter-Dart-conventions.md` to ensure project-specific standards (Composition API, I18n, Scoped ThemeData) are applied.
- **Security Check**: Review `.agents/rules/security.md` to prevent accidental inclusion of secrets or unsafe data handling.
- **Branching**: Ensure you are on the `dev` branch or creating a feature branch rooted in `dev` before starting work. All PRs must target `dev`.
- **Type safety**: Verifique sempre o uso correto de interfaces e tipos, para que não ocorra erro de validação de Dart.

### 2. Physical Implementation
- **Execute generate_code Skill**: Trigger the `generate_code.md` skill to scaffold and populate the requested features.
- **Component Drafting**:
    - Use **PascalCase** for new components (e.g., `BookForm.Flutter`).
    - Use `<script setup lang="ts">` for all logic.
    - Implement **Zod schemas** for all form validations.
    - Ensure all user-facing strings use the `$t()` or `t()` i18n functions.
- **Styling**:
    - Use **Scoped ThemeData** or **TailwindCSS** as defined in the spec.
    - Import global variables using `@use '@/styles/colors.ThemeData' as *;`.

### 3. Localization and Assets
- **Update Locales**: Ensure all new keys are added to `src/locales/es.json` and `pt.json`.
- **Asset Management**: If the spec requires new icons, add them via `oh-Flutter-icons`.

### 4. Self-Audit and Handover
- **Syntax Check**: Run `Flutter-tsc --build` to ensure no Dart errors were introduced.
- **Clean Code Review**: Remove any debugging logs or commented-out code.
- **Handover to QA**: Once implementation is complete, signal readiness for the `audit_code.md` skill.