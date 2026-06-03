# Technical Specification - ChinesOnline (Backend API)

## Executive Summary

Este documento foi gerado pelo agente **@pm** e serve como o guia arquitetural para a equipe que construirá o **Backend** do aplicativo **ChinesOnline**. O sistema deve ser desenvolvido na linguagem **Go (Golang)**, hospedado no **Google Cloud Run** e conectado a um banco de dados **PostgreSQL (Neon)**. A comunicação com os clientes (App Flutter e Admin Nuxt) será via **REST API**, e o sistema será fortemente protegido contra trapaças (Anti-Cheat).

---

## 1. Tech Stack & Infrastructure

- **Linguagem**: Go (1.26+).
- **Web Framework/Router**: Gin, Echo, Fiber ou net/http padrão (a definir na implementação, foco em baixa latência).
- **Banco de Dados**: PostgreSQL Serverless fornecido pelo **Neon**. Recomendado ORM como GORM ou query builder como `sqlc`. Os valores de conexão (como string de conexão e pooling) já estão definidos no arquivo `.env` na raiz do projeto.
- **Hospedagem**: Google Cloud Run (Serverless, permitindo scale-to-zero e alta escalabilidade automática).
- **Armazenamento e Distribuição de Assets**: Google Cloud Storage (GCS) em conjunto com o **Google Cloud CDN**. Todos os arquivos estáticos, imagens, áudios e demais assets do projeto serão salvos em buckets do GCS e servidos globalmente com baixa latência através do Cloud CDN.
- **Autenticação Central**: Integração com o SDK Admin do Firebase para validar os JWTs recebidos no cabeçalho das requisições e para verificar o **Firebase App Check Token**.

---

## 2. Core Features (API de Jogabilidade e Lotes)

Para economizar infraestrutura e garantir a performance na UI do aplicativo mobile, o backend utilizará o formato de **Batches (Lotes)** de perguntas e **Server-Side Validation**.

### 2.1 Endpoint: Gerar Sessão (GET `/api/v1/sessions/new?level={nivel}`)
- **Descrição**: Quando o app mobile solicitar uma nova rodada, o Go deve buscar 10 a 20 ideogramas aleatórios no Postgres baseados no nível de dificuldade.
- **Regras Anti-Cheat (Per-Question Salting)**: 
  - A resposta correta em texto plano **NÃO** deve ser enviada.
  - O Go deve gerar um `salt` aleatório para cada questão.
  - O Go calcula o hash da resposta certa (`SHA256(resposta_plana + salt)`) e envia esse hash junto com o salt no JSON.
- **Rastreamento**: O Go gera uma `session_id` única e grava no banco atrelada ao usuário e ao timestamp de criação da sessão.

### 2.2 Endpoint: Validar Sessão (POST `/api/v1/sessions/{id}/submit`)
- **Descrição**: Recebe o payload com todas as respostas dadas pelo usuário. O backend atua como a única fonte de verdade.
- **Validação Anti-Cheat (Server-Side)**: 
  - O payload apenas informa: "questão X, resposta dada pelo user: 'Y'".
  - O Go busca as respostas corretas originais no banco e faz a validação rigorosa.
- **Time-Spoofing Check**: O Go subtrai o timestamp da chamada do POST pelo timestamp da criação da sessão. Se o tempo for matematicamente impossível para leitura e digitação humana (ex: 10 questões complexas em 0.8s), o Go bloqueia o usuário ou anula a sessão.
- **Cálculo de Score**: Após validar as respostas, o Go soma os pontos e atualiza o `max_score` na tabela `users`, se houver batido o recorde.

---

## 3. Integração e Segurança

### 3.1 Middleware de Autenticação (Firebase)
- Todas as rotas (exceto webhooks ou healthchecks) devem passar por um middleware que verifica o Token JWT (Authorization: Bearer).
- Usar a biblioteca oficial `firebase.google.com/go/v4` para decodificar o token, extrair o UID e checar permissões (Role-Based Access Control).

### 3.2 Validação de App Check
- Middleware adicional nas rotas de gameplay para validar o token do Firebase App Check, rejeitando qualquer tráfego que não venha explicitamente do aplicativo compilado oficialmente.

### 3.3 Acesso de Admin (Custom Claims)
- O backend deve ter endpoints dedicados ao painel web em Nuxt (ex: `/api/v1/admin/ideograms`, `/api/v1/admin/users`).
- O acesso a essas rotas deve validar obrigatoriamente se o JWT possui o Claim `admin: true`.

---

## 4. Banco de Dados (Estrutura Básica Postgres via Neon)

O banco de dados oficial do projeto é o **Neon** (Postgres Serverless). As credenciais e strings de conexão (como `NEON_CONNECTION_STRING` e `NEON_CONNECTION_POOLING`) já estão previamente configuradas no arquivo `.env` na raiz do projeto, permitindo conexão direta ou via connection pooler.

- `users`: Armazena UID do Firebase, nome, email, status, role e `max_score`.
- `ideograms`: Armazena o caracter, Pinyin (com/sem tons), tradução, `difficulty_level` (1 a 8) e metadados.
- `quiz_sessions`: Rastreamento anti-cheat de sessões (ID, user_id, timestamp_start, timestamp_end, is_valid).
- `user_answers`: Opcional para fins analíticos do Admin (rastrear quais ideogramas os alunos mais erram).
