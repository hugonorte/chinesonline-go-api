# Hot Paths & Modification Impact - Flutter App

This file highlights the most frequently modified areas of the codebase and the impact of changing them.

## 🔴 High Frequency & High Impact (Core Feature)
**`lib/features/quiz/`**
- **Impact**: Critical. Modifying the quiz logic (levels 1-8 rules) directly impacts gameplay and scoring.
- **When modifying**: Ensure state updates correctly transition the UI between Multiple Choice (Level 1-2) and Text Input (Level 3-8). Verify the score tally works correctly.

## 🟠 Medium Frequency
**`lib/features/admin/`**
- **Impact**: High for admins, low for users.
- **When modifying**: Changes here involve Firebase CRUD for ideograms. Make sure Firestore Security Rules are respected and data structures match the Quiz engine's expectations.

**`lib/features/profile/`**
- **Impact**: Medium. Saving max score and user data.
- **When modifying**: Verify that updates to `maxScore` only occur if the new score is strictly greater than the historical max.

## 🟡 Low Frequency (Stable)
**`lib/core/firebase/`**
- **Impact**: Extreme. Changes here break the entire app's connectivity.
- **When modifying**: Only touch this when adding new Firebase services (e.g., configuring Crashlytics or Analytics).

**`lib/core/theme/`**
- **Impact**: Global visual change.
- **When modifying**: Ensure all widgets using `Theme.of(context)` are checked for visual regressions.
