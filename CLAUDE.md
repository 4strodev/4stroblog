# 4stroblog — CLAUDE.md

## What this project is

A personal blog application. Server-rendered Go web app using HTMX for partial page updates, with a JSON REST API layer. Blog posts are Markdown files on disk. Authentication uses cookie-based sessions for the public site and JWT for the admin area.

## Stack at a glance

- **Go 1.24** — Gofiber v3 (HTTP), GORM + SQLite (storage), Koanf (config)
- **MinIO** — S3-compatible object storage for file uploads
- **HTMX** — frontend interactivity (no JS build step)
- **Tailwind CSS** — compiled from `packages/theme/` into `packages/site/assets/index.css`
- **Docker Compose** — local dev services (MinIO)

## How the code is organized

```
packages/
  site/           # Main Go application
    cmd/          # Entry point
    server/       # HTTP layer — modules, controllers, routes
    features/     # Domain logic (DDD: domain / application / infrastructure)
    shared/       # Cross-cutting: DB, S3, config, logger, i18n, domain errors
    views/        # Go html/template files
    config/       # config.toml, i18n TOML files
  theme/          # Tailwind CSS source → compiled to site/assets/index.css
  go-markdown-emoji/  # Local Go module (markdown extension)
infrastructure/
  docker-compose.yaml  # MinIO for local dev
```

## Architecture conventions

- **Module/DI pattern** (NestJS-inspired via `wiring_graphs`): each `core.Module` declares singletons, transients, exports, and controllers.
- **DDD layers inside `features/`**: `domain/` → pure entities and interfaces; `application/` → use-case services and DTOs; `infrastructure/` → GORM implementations and mappers.
- **Repository pattern**: domain interfaces in `domain/`, GORM implementations in `infrastructure/`. Keep ORM annotations out of domain objects.
- **Typed domain errors**: use `domainerrors.Errorf` / `domainerrors.Is` — never match errors by string.

## Coding style

- **Prefer clear over clever.** Code should be readable by someone with zero context.
- Flat, explicit code beats indirection. Name things for what they do.
- No speculative abstractions. Solve the problem in front of you.

## Agents to consult

- **Architecture questions** → `@architecture-oracle` agent. Use it before making structural changes or adding new features.
- **After changing behavior** → `@docs-sync-writer` agent. Use it to keep documentation in sync with what the code actually does.
