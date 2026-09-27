# Stack

## Backend (`packages/site`)

| Technology | Purpose |
|------------|---------|
| [Go](https://go.dev) 1.25 | Language of the application. |
| [Fiber v3](https://gofiber.io) | HTTP server and router. Serves pages, the JSON API and static assets. |
| [Fiber html template](https://github.com/gofiber/template) | Renders the Go `html/template` files in `views/`. |
| [wiring_graphs](https://github.com/4strodev/wiring_graphs) | Dependency injection. Wires modules, services and controllers together. |
| [GORM](https://gorm.io) + SQLite | ORM and embedded database for app data. SQLite needs no separate server (requires `CGO_ENABLED=1`). |
| [Koanf](https://github.com/knadh/koanf) | Loads configuration from `config/config.toml`. |
| [minio-go](https://github.com/minio/minio-go) | S3 client used to store and delete uploaded files. |
| [gomarkdown](https://github.com/gomarkdown/markdown) | Converts post Markdown to HTML. |
| [go-markdown-emoji](../../packages/go-markdown-emoji) | Local module. Adds `:emoji:` shortcodes to the Markdown renderer. |
| [bluemonday](https://github.com/microcosm-cc/bluemonday) | Sanitizes the rendered HTML to block XSS (Cross-Site Scripting). |
| [cristalhq/jwt](https://github.com/cristalhq/jwt) | Builds and verifies JWTs (JSON Web Tokens) for the admin area. |
| [x/crypto (bcrypt)](https://pkg.go.dev/golang.org/x/crypto/bcrypt) | Hashes and checks user passwords. |
| [x/text](https://pkg.go.dev/golang.org/x/text) | Language matching for i18n (internationalization). |
| [gjson](https://github.com/tidwall/gjson) | Reads nested translation keys in the i18n service. |
| [uuid](https://github.com/google/uuid) | Generates entity ids. |

## Frontend

| Technology | Purpose |
|------------|---------|
| [HTMX](https://htmx.org) | Partial page updates driven by HTML attributes. No JS bundler. Loaded from CDN. |
| [Tailwind CSS](https://tailwindcss.com) | Utility CSS. Source in `packages/theme/`, compiled to `packages/site/assets/index.css`. |
| PostCSS + autoprefixer | Tailwind build pipeline. Adds vendor prefixes. |
| [highlight.js](https://highlightjs.org) | Syntax highlighting of code blocks in posts. Loaded from CDN. |
| [Feather Icons](https://feathericons.com) | SVG icons. Loaded from CDN. |

## Storage and infrastructure (`infrastructure/`)

| Technology | Purpose |
|------------|---------|
| [RustFS](https://rustfs.com) | S3-compatible object storage for uploads. Runs in Docker Compose (API `:9000`, console `:9001`). |
| [Docker](https://www.docker.com) | Multi-stage build (`packages/Dockerfile`): builds the theme, compiles Go, ships a distroless image. |
| Docker Compose | Starts object storage and the app locally. |

## Tooling

| Technology | Purpose |
|------------|---------|
| [Task](https://taskfile.dev) | Command runner (`Taskfile.yaml`): `install`, `dev`, `build`, `docker:init`. |
| [pnpm](https://pnpm.io) | Installs the theme dependencies. |
| [arelo](https://github.com/makiuchi-d/arelo) | Live reload in `task dev`. Restarts the server on Go, view or config changes. |
| [Bruno](https://www.usebruno.com) | API request collection in `bruno/`. |
