# Cronograma Geral do Projeto ChinesOnline (MVP)
- **Fase 1:** Infraestrutura & Fundação (Semana 1)
- **Fase 2:** Backend API - CRUD e Auth (Semanas 2-3)
- **Fase 3:** Frontend Administrativo (Semana 4)
- **Fase 4:** Backend API - Game Engine (Semana 5)
- **Fase 5:** O Aplicativo Flutter - MVP (Semanas 6-8)
- **Fase 6:** Anti-Cheat & Deploy (Semanas 9-10)

---

# Tarefas Detalhadas: Backend em Go

## Fase 2: Backend API - CRUD e Auth
- [ ] **Setup Inicial do Projeto**
  - [ ] Inicializar o módulo (`go mod init`).
  - [ ] Instalar framework web (Gin, Fiber ou Echo).
  - [ ] Instalar biblioteca do banco de dados (GORM ou sqlc).
  - [ ] Configurar leitura de variáveis de ambiente (`.env`).
- [ ] **Conexão com Banco de Dados**
  - [ ] Criar arquivo de conexão para o Neon (PostgreSQL).
  - [ ] Configurar Connection Pooling.
- [ ] **Integração Firebase Auth (Middleware)**
  - [ ] Instalar Firebase Admin SDK para Go.
  - [ ] Carregar a Service Account Key do Firebase via env vars.
  - [ ] Criar o middleware `VerifyJWT` para proteger rotas.
  - [ ] Criar o middleware `VerifyAdmin` que decodifica o JWT e valida se `admin: true`.
- [ ] **Rotas do Módulo Administrativo (Protegidas por VerifyAdmin)**
  - [ ] `POST /api/v1/admin/ideograms`: Cadastro de novos ideogramas.
  - [ ] `PUT /api/v1/admin/ideograms/:id`: Atualização de Pinyin, traduções e nível.
  - [ ] `GET /api/v1/admin/ideograms`: Listagem com paginação e filtros.
  - [ ] `DELETE /api/v1/admin/ideograms/:id`: Remoção lógica (soft delete).
  - [ ] `GET /api/v1/admin/users`: Listagem de usuários cadastrados e seus `max_score`.

## Fase 4: Backend API - Game Engine
- [ ] **Lógica de Geração de Sessões (Batches)**
  - [ ] Criar rota `GET /api/v1/quiz/sessions/new?level={nivel}`.
  - [ ] Escrever query SQL que busca de 10 a 20 ideogramas aleatórios do nível solicitado.
  - [ ] Criar registro na tabela `quiz_sessions` armazenando o `user_id` e o `timestamp_start`.
- [ ] **Implementação do Anti-Cheat (Hashing & Salting)**
  - [ ] Para cada ideograma sorteado, gerar um `salt` aleatório forte.
  - [ ] Implementar a função `sha256(resposta_correta + salt)`.
  - [ ] Montar o payload JSON contendo o ideograma, o `salt` e o `correct_hash` (NUNCA enviar a resposta em texto plano).
- [ ] **Lógica de Validação e Submissão**
  - [ ] Criar rota `POST /api/v1/quiz/sessions/{id}/submit`.
  - [ ] **Validação de Time-Spoofing**: Calcular `time.Now() - timestamp_start`. Se for menor que um limite mínimo humano (ex: 10 questões em 1 segundo), rejeitar com erro 403.
  - [ ] **Validação de Respostas**: Iterar sobre o array de respostas enviadas pelo usuário.
  - [ ] Buscar as respostas corretas originais no banco e compará-las para calcular os pontos reais.
  - [ ] **Atualização de Score**: Buscar o `max_score` atual do usuário na tabela `users`. Se o score da sessão for maior, dar UPDATE.
- [ ] **Implementação do Middleware Firebase App Check**
  - [ ] Integrar a validação do token do App Check nas rotas do Game Engine para bloquear scripts de terceiros.

## Fase 6: Deploy e Otimização
- [ ] Escrever o `Dockerfile` otimizado (multi-stage build) para compilar o binário Go minúsculo.
- [ ] Configurar logs estruturados (ex: zerolog ou slog) para o Google Cloud Logging.
- [ ] Preparar servidor web para rodar na porta exigida pelo Cloud Run (leitura da variável `$PORT`).
