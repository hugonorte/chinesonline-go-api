# Project Architecture - ChinesOnline

## 1. Overview
This project is a mobile application built with **Flutter** and powered by **Google Firebase**. It functions as an interactive quiz game for learning Chinese characters and words.

## 2. Core Technologies
- **Mobile App**: Flutter (Dart)
- **Backend/BaaS**: Firebase (Auth, Firestore, Analytics)
- **State Management**: Riverpod (Recommended) or BLoC
- **Routing**: GoRouter

## 3. Directory Structure
We follow a **Feature-First Architecture** to keep the codebase modular and scalable.

```
lib/
├── core/                   # Componentes globais e configurações
│   ├── theme/              # Temas, cores, tipografia
│   ├── utils/              # Funções auxiliares e helpers
│   ├── constants/          # Strings, assets paths
│   └── firebase/           # Inicialização e configs do Firebase
├── features/               # Módulos principais do app
│   ├── auth/               # Autenticação (Login, Registro)
│   ├── quiz/               # Lógica do jogo (Múltipla escolha, Input de Pinyin)
│   ├── profile/            # Perfil e histórico de pontuação
└── main.dart               # Entrypoint da aplicação
```

## 4. Key Application Layers (Inside Features)
Each feature directory contains:
- `presentation/`: Widgets, Pages, and ViewModels/StateNotifier.
- `domain/`: Business logic, Entities, Models.
- `data/`: Repositories and APIs.

## 5. Domain Modules
### 5.1 Quiz Engine
- Manages the difficulty progression (Levels 1 to 8).
- Switches UI from Multiple Choice (Lv 1-2) to Text Input (Lv 3-8).
- Calculates points locally for feedback, but relies on Server validation.

### 5.2 User & Scoring
- Users are authenticated via Firebase Auth.
- Communicates with Go API to validate scores.
