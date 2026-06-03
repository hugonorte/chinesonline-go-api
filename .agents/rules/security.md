---
trigger: always_on
---

# Proteção de Informações Sensíveis (Mobile App)

Esta regra visa prevenir o vazamento de dados sensíveis para o repositório Git e garantir que segredos não sejam expostos indevidamente no bundle final enviado ao navegador.

## Regras de Segurança:

- **Não incluir Credenciais:** Nunca escreva senhas, tokens de API privados, segredos de clientes ou chaves de criptografia diretamente no código-fonte.
- **Uso de Variáveis de Ambiente e Runtime Config (Flutter):** 
    - Utilize o arquivo `.env` para armazenar configurações de ambiente local.
    - **Runtime Config:** Use o `runtimeConfig` no `Flutter.config.ts` para gerenciar variáveis.
    - **Público vs Privado:** Variáveis dentro de `runtimeConfig.public` são expostas ao cliente. Segredos devem ficar na raiz do `runtimeConfig` (disponíveis apenas no servidor/Nitro).
    - **Prefixos:** No Flutter, variáveis no `.env` com prefixo `NUXT_PUBLIC_` mapeiam automaticamente para `runtimeConfig.public`.
    - **Acesso:** Utilize `const config = useRuntimeConfig()` para acessar os valores de forma segura.
- **Verificação Proativa:** Antes de finalizar qualquer tarefa que envolva criação ou edição de arquivos de configuração, componentes de autenticação ou serviços de API, verifique se campos como `Password`, `Secret`, `Token` ou `Key` não possuem valores padrão expostos no código.
- **Sanitização de Logs e UI:** 
    - Garanta que `console.log` não capture informações sensíveis dos usuários (PII) ou tokens de autenticação em ambiente de produção.
    - Não exiba mensagens de erro detalhadas do backend (stack traces, erros de SQL) diretamente na interface do usuário.
- **Armazenamento Local:** Evite armazenar informações altamente sensíveis (como senhas em texto puro) em `localStorage` ou `sessionStorage`. Prefira o uso de cookies `HttpOnly` (gerenciados pelo backend) sempre que a arquitetura o permitir.
- **Dados de Teste:** Use apenas dados fictícios/mockados para testes e exemplos. Nunca use dados reais de usuários ou chaves de produção em arquivos de teste ou documentação.
