# 4stroblog docs

4stroblog is a personal blog. It is a server-rendered Go web app: pages are built on the server
with Go templates, and HTMX swaps page fragments without a JavaScript build step. Posts are
written in Markdown. Uploaded files live in S3-compatible object storage.

## Built with

- **Go + Fiber**: HTTP server, routing and template rendering.
- **SQLite + GORM**: application data (users, sessions, posts metadata, uploads).
- **S3 storage (RustFS)**: uploaded files.
- **HTMX**: partial page updates from HTML attributes.
- **Tailwind CSS**: styling, compiled to a single CSS file.
- **Task + Docker Compose**: dev commands and local services.

See [architecture/stack.md](architecture/stack.md) for the full list and why each piece is there.

## Contents

- [Architecture](architecture/README.md)
  - [Stack](architecture/stack.md)
- Features
  - Pages
    - [Front matter](features/pages/front-matter.md)
