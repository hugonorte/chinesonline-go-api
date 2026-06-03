---
description: Executa apenas o papel de Product Manager (@pm) para definir especificações técnicas.
---

Quando o usuário digitar `/pm`, orquestre o papel de **Product Manager (@pm)** seguindo estritamente as definições em `.agents/agents.md`.

### Sequência de Execução:
1. **Leitura de Objetivos**: Leia o arquivo `task.md`. Caso o arquivo `task.md` esteja vazio, apenas faça uma revisão na documentação contida em `production_artifacts/Technical_Specification.md` e certifique-se que tudo está de acordo com as especificações gerais do projeto. 
2. **Definição Técnica**: Atue como o **Product Manager** e execute a skill `write_specs.md` para criar ou atualizar o arquivo `production_artifacts/Technical_Specification.md`.
3. **Pausa para Aprovação**: Você **DEVE** pausar para aprovação explícita do usuário. Se o usuário fornecer feedback ou adicionar comentários diretamente no arquivo Markdown, revise o documento. Repita este passo até que o usuário aprove.