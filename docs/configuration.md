# Configuration

The app reads `config/config.toml`, relative to the directory it runs from (`packages/site` in dev).
The HTTP port comes from the `PORT` environment variable, not from this file.

```toml
[site]
title = "4strodev"
admin_email = "admin@example.com"

[views]
folder = "./views"

[i18n]
default = "en"
folder = "./config/i18n"

[db.sqlite]
path = "./test.db"

[storage.s3]
url = "minio://user:password@localhost:9000"
bucket = "uploads"

[auth.jwk]
secret = "<random hex string>"

[environment]
god_mode = false
```

| Key | Purpose |
|-----|---------|
| `site.title` | Site name. Default page title and `.Site.Title` in templates. |
| `site.admin_email` | Email of the admin user created on first start. |
| `views.folder` | Folder with the UI templates. See [Views](features/views/README.md). |
| `i18n.default` | Language used when the requested one has no translation. |
| `i18n.folder` | Folder with the translation TOML files. |
| `db.sqlite.path` | SQLite database file. Created if missing. |
| `storage.s3.url` | S3 endpoint with credentials, e.g. `minio://user:password@host:port`. |
| `storage.s3.bucket` | Bucket for uploads. Created on startup if missing. |
| `auth.jwk.secret` | Secret used to sign and verify JWTs. Keep it private. |
| `environment.god_mode` | Access all restricted sections, for testing. Not used yet. |

## First start

On first start the app creates an admin user with `site.admin_email` and the password `1234`.
Change that password after the first login.

`[server] port` in the default file is not read. Set `PORT` instead (`.dev.env` in dev).
