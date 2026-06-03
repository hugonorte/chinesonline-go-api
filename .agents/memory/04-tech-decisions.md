# Technical Decisions

This document outlines the rationale behind the primary technologies chosen for the ChinesOnline app.

## 1. Flutter
**Why**: Allows building high-quality, natively compiled applications for iOS, Android, and Web from a single codebase. Given the interactive nature of a gamified quiz app, Flutter's rich widget system and smooth animations are ideal.
**Constraint**: Ensure no platform-specific code (e.g., native Android/iOS channels) is written unless absolutely necessary.

## 2. Google Firebase
**Why**: Provides a complete Backend-as-a-Service (BaaS).
- **Firestore**: Perfect for real-time leaderboards, storing ideograms, and syncing user data. Document-based structure fits the quiz data nicely.
- **Firebase Auth**: Ready-to-use secure authentication flow.
**Constraint**: Data reading is high-frequency during gameplay; optimize queries and cache locally when possible to reduce Firestore read costs.

## 3. Game Logic Strategy
**Why Difficulty Levels 1-8**: Progressive difficulty enhances user engagement.
- Level 1-2: Multiple Choice builds confidence and initial recognition.
- Level 3-8: Pinyin text input reinforces active recall and pronunciation spelling.
**Implementation Detail**: The Quiz Engine must validate string inputs dynamically (e.g., ignoring case, trimming spaces).

## 4. State Management (Riverpod / BLoC)
**Why**: Passing state manually in a quiz game with timers, scores, and changing UI formats is error-prone. A robust state management solution ensures the UI simply reacts to the current game state (e.g., `GameState.playing`, `GameState.levelComplete`).
