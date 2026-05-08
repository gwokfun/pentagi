# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Core Interaction Rules

1. **Always use English** for all interactions, responses, explanations, and questions with users.
2. **Password Complexity Requirements**: For all password-related development (registration, password reset, API token generation, etc.), the following rules must be enforced:
   - Minimum 12 characters
   - Must contain at least 1 uppercase letter, 1 lowercase letter, 1 number, and 1 special character
   - Common weak passwords (e.g., `password`, `123456`) are prohibited
   - Both backend and frontend validation must be implemented; do not rely on frontend validation alone

## Project Overview

**PentAGI** is an automated security testing platform powered by AI agents. This is a **simplified version** configured to run with only **DeepSeek** as the LLM provider, making it easy to get started quickly.

The application is a monorepo with:
- **`backend/`** — Go REST + GraphQL API server
- **`frontend/`** — React + TypeScript web UI

## Quick Start

### Minimum Requirements

```bash
# .env file - only these are required:
DEEPSEEK_API_KEY=your-deepseek-api-key
PENTAGI_POSTGRES_PASSWORD=change-this-password
COOKIE_SIGNING_SALT=change-this-salt
```

Run with:
```bash
docker compose up -d
```

Access at: https://localhost:8443

## Build & Development Commands

### Backend (run from `backend/`)

```bash
go mod download                              # Install dependencies
go build -trimpath -o pentagi ./cmd/pentagi  # Build main binary
go test ./...                                # Run all tests
golangci-lint run --timeout=5m               # Lint

# Code generation (run after schema changes)
go run github.com/99designs/gqlgen --config ./gqlgen/gqlgen.yml  # GraphQL resolvers
swag init -g ../../pkg/server/router.go -o pkg/server/docs/ --parseDependency --parseInternal --parseDepth 2 -d cmd/pentagi  # Swagger docs
```

### Frontend (run from `frontend/`)

```bash
npm ci                    # Install dependencies
npm run dev               # Dev server on http://localhost:8000
npm run build             # Production build
npm run lint              # ESLint check
npm run test              # Vitest
npm run graphql:generate  # Regenerate GraphQL types from schema
```

### Docker (run from repo root)

```bash
docker compose up -d                          # Start core services
docker build -t local/pentagi:latest .        # Build image
```

## Architecture (Simplified Version)

### What's Included

- **DeepSeek LLM Provider** - The only supported LLM provider
- **PostgreSQL + pgvector** - Required database with vector storage
- **DuckDuckGo Search** - Default search engine (no API key needed)
- **Web Scraper** - For gathering web information
- **Docker Isolation** - Safe sandboxed execution
- **Multi-Agent System** - Researcher, Developer, Executor agents

### What's Removed

To keep the setup simple, these features are **not included**:
- ❌ Multiple LLM providers (OpenAI, Anthropic, Gemini, Bedrock, Ollama, etc.)
- ❌ OAuth authentication (Google, GitHub)
- ❌ Advanced search engines (Tavily, Perplexity, Google, Searxng, etc.)
- ❌ Observability stack (Langfuse, OpenTelemetry, Grafana)
- ❌ Knowledge graph (Graphiti + Neo4j)
- ❌ PentAGI Cloud/License integration

### Backend Package Structure

| Package | Role |
|---|---|
| `cmd/pentagi/` | Main entry point; initializes config, DB, server |
| `pkg/config/` | Environment-based config parsing (simplified) |
| `pkg/server/` | Gin router, middleware, auth (JWT/API tokens), Swagger |
| `pkg/controller/` | Business logic for REST endpoints |
| `pkg/graph/` | gqlgen GraphQL schema and resolvers |
| `pkg/database/` | GORM models, SQLC queries, goose migrations |
| `pkg/providers/` | DeepSeek LLM provider integration |
| `pkg/tools/` | Penetration testing tool integrations |
| `pkg/docker/` | Docker SDK wrapper for sandboxed execution |
| `pkg/terminal/` | Terminal session management |
| `pkg/queue/` | Async task queue |
| `pkg/csum/` | Chain summarization for context management |

### Frontend Structure

```
frontend/src/
├── app.tsx / main.tsx     # Entry points and router setup
├── pages/                 # Route-level page components
│   ├── flows/             # Flow management UI
│   └── settings/          # Provider and prompt settings
├── components/
│   ├── layouts/           # App shell layouts
│   └── ui/                # Base Radix UI components
├── graphql/               # Auto-generated Apollo types
├── hooks/                 # Custom React hooks
├── lib/                   # Apollo client, HTTP utilities
└── schemas/               # Zod validation schemas
```

### Data Flow

1. User creates a "flow" (penetration test) via the UI or REST API
2. Backend queues the flow and spawns agent goroutines
3. Researcher agent gathers information; Developer plans strategies; Executor runs tools
4. Results stored in PostgreSQL with pgvector for semantic search
5. Real-time progress pushed to frontend via GraphQL subscriptions

### Authentication

- **Session cookies** for browser login (secure, httpOnly)
- **Bearer tokens** (API tokens table) for programmatic access
- OAuth has been removed in this simplified version

### Configuration

All configuration is in `.env`:

```bash
# Required
DEEPSEEK_API_KEY=your-key              # Get from platform.deepseek.com
PENTAGI_POSTGRES_PASSWORD=changeme     # Database password
COOKIE_SIGNING_SALT=changeme           # Session security

# Optional
DUCKDUCKGO_ENABLED=true                # Search engine (default: true)
MAX_GENERAL_AGENT_TOOL_CALLS=100       # Agent limits
AGENT_PLANNING_STEP_ENABLED=false      # Planning mode
```

### Code Generation

When modifying `backend/pkg/graph/schema.graphqls`, re-run gqlgen to regenerate resolvers. When modifying REST handler annotations, re-run swag to update Swagger docs. When modifying `frontend/src/graphql/*.graphql` queries, re-run `npm run graphql:generate`.

## Adding Features Back

If you need features that were removed (multiple LLM providers, OAuth, observability, etc.), you'll need to:

1. Restore the relevant config fields in `pkg/config/config.go`
2. Restore provider implementations in `pkg/providers/providers.go`
3. Update `.env.example` with the additional variables
4. Potentially restore docker-compose override files for additional services

For reference, check the original PentAGI repository for the full implementation.

## Development Tips

- Use `docker compose logs -f pentagi` to see backend logs
- Frontend dev server proxies API requests to backend
- Database migrations run automatically on startup
- Default admin user must be created on first launch via UI
