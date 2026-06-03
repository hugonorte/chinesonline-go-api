---
description: Executa a suíte de testes E2E/Componentes com o Flutter Integration Test no Mobile App.
---

# Workflow: Executar Testes do Sistema (Flutter Integration Test)

Este workflow automatiza a execução de testes E2E (End-to-End) e/ou componentes para validar a integridade das interfaces e fluxos do usuário após alterações no projeto Mobile App.

## Passos:

// turbo

1. Executa o comando de testes em modo headless no diretório raiz:

   ```bash
   npx Flutter Integration Test run
   ```
   *(Nota: Caso existam testes unitários Flutter Test que devam rodar em conjunto, execute `flutter test:unit` primeiro).*

2. Exibe o resultado da execução dos testes (passes, fails, pending) para o usuário.
3. Se houver falhas, sugere a análise dos logs de erro no terminal e vídeos/screenshots (gerados pelo Flutter Integration Test na pasta `Flutter Integration Test/screenshots` ou `Flutter Integration Test/videos`) para correção do componente.
