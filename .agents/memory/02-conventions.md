# Coding Conventions - Flutter & Dart

## 1. Naming Conventions
- **Classes, Enums, Typedefs**: `PascalCase`
- **Variables, Functions, Methods**: `camelCase`
- **Files and Directories**: `snake_case` (e.g., `quiz_screen.dart`, `user_model.dart`)
- **Constants**: `lowerCamelCase` or `UPPER_SNAKE_CASE` (follow Dart style guide, generally `lowerCamelCase` is preferred for constants).

## 2. Dart & Flutter Patterns
- **Stateless/Stateful**: Prefer `StatelessWidget` when using state management (like Riverpod), keeping business logic out of the UI.
- **Null Safety**: Strict null safety must be maintained. Use `?` and `!` responsibly. Avoid `!` unless absolutely certain. Use `?.` and `??` for safe fallbacks.
- **Async/Await**: Always use `async`/`await` for Firebase calls. Handle exceptions using try/catch blocks.

## 3. Firebase Interactions
- **Data Models**: Use `fromJson` and `toJson` methods for serialization/deserialization of Firestore documents.
- **Collections**: Hardcode collection names as constants in a dedicated `FirebaseConstants` file to avoid typos.

## 4. UI/UX Rules
- **Themes**: Do not hardcode colors or text styles in widgets. Use `Theme.of(context)` to ensure consistency.
- **Responsive**: Design should be flexible, accommodating different phone screen sizes.
- **Localization**: All user-facing strings must use the `intl` package or Flutter's localizations. No hardcoded strings in the UI.

## 5. Testing
- Use `flutter test` for unit testing logic and state controllers.
- Use `flutter test --machine` for widget testing key UI components.
