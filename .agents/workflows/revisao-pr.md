---
description: Este workflow automatiza a revisão de código em repositórios GitHub seguindo os padrões sênior de .NET.
---

# Workflow: Review Automático de Pull Request .NET
Este workflow automatiza a revisão de código em repositórios GitHub seguindo os padrões sênior de .NET.

## 🏁 Trigger
Comando: "Iniciar revisão de PR" ou "Simular CodeRabbit"

## 📝 Passos do Workflow

### Passo 1: Coleta de Contexto
- **Ação:** Pergunte ao usuário: "Qual é o ID do Pull Request que você deseja revisar?"
- **Variável:** Guarde a resposta como `{{PR_ID}}`.
- **Nota:** Certifique-se de que o Pull Request tem como alvo a branch `dev`, conforme a nova política de Git Flow.

### Passo 2: Listagem de Arquivos
- **Ação:** Use a ferramenta `github-mcp.list_pull_request_files` informando o `{{PR_ID}}`.
- **Filtro:** Identifique apenas os arquivos com extensão `.cs`.

### Passo 3: Análise e Crítica
- **Ação:** Para cada arquivo `.cs` identificado:
    1. Leia o conteúdo do diff/arquivo.
    2. Aplique rigorosamente a regra `@dotnet-reviewer`.
    3. Identifique violações de: Async/Await, Performance de LINQ, Injeção de Dependência e Padrões de Clean Code.

### Passo 4: Publicação de Comentários
- **Ação:** Se encontrar problemas:
    - Use `github-mcp.create_inline_comment` para postar cada sugestão diretamente na linha correspondente no GitHub.
    - Se o código estiver perfeito: Poste um comentário geral no PR parabenizando o desenvolvedor.

### Passo 5: Verificar necessidade de ajustes 
- **Ação:** Havendo necessidade de ajustes:
   - Acionar o agente @engineer conforme o arquivo agents.md para implementar os ajustes
   - Após o @engineer efetuar os ajustes, acionar o agente @qa 
- **Ação:** Não havendo necessidade de ajustes:
   - Prosseguir para o passo 6 (finalização)

### Passo 6: Finalização
- **Ação:** Informe ao usuário: "Revisão concluída no PR #{{PR_ID}}. Você pode conferir os comentários diretamente no GitHub."