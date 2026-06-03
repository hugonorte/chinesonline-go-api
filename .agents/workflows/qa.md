---
description: Executa apenas o papel de QA Engineer (@qa) para auditar e corrigir o código.
---

Quando o usuário digitar `/qa`, orquestre o papel de **QA Engineer (@qa)** seguindo estritamente as definições em `.agents/agents.md`.

### Sequência de Execução:
1. **Auditoria de Código**: Execute a skill `audit_code.md` para analisar o código gerado pelo Engenheiro e compará-lo com a `Technical_Specification.md`.
2. **Bug Hunting**: Procure agressivamente por dependências ausentes nas configurações, erros de sintaxe e bugs lógicos.
3. **Correção Proativa**: Aponte para o agente com papel @engineer para que o agente @engineer corrija diretamente qualquer falha encontrada, sobrescrevendo os arquivos necessários no diretório `src/`.
4. **Resumo da Auditoria**: Ao finalizar, apresente ao usuário um resumo dos problemas encontrados e das correções realizadas.