# Views

Views are the Go `html/template` files that build the site UI: layouts, pages and components.
They are loaded from a folder set in the config, so you can replace the default UI with your own.

## Custom views

Point `views.folder` in `packages/site/config/config.toml` to your views folder:

```toml
[views]
folder = "./views"
```

The path is relative to the directory the app runs from (`packages/site` in dev).
The default UI lives in `packages/site/views/`. Copy it as a starting point for your own.

## Required views

The app checks the views folder on startup. It refuses to start if any of these is missing:

| File | Why |
|------|-----|
| `layouts/main.html` | Default layout for every page. |
| `pages/index.html` | Home page (`/site`). |

The error lists every missing file:

```
views folder "./my-theme" is missing required templates: layouts/main.html, pages/index.html
```

The check covers only these two files. If your layout includes other templates
(e.g. `{{ template "components/head" . }}`), they must exist too, or the page fails when rendered.
The default views also use `scaffolds/post.html` (blog posts) and `pages/not-found.html`.

## Folder structure

```
views/
  layouts/     # page shells, e.g. main.html, full-width.html
  pages/       # one file per route: pages/about.html -> /site/about
  components/  # shared pieces included from layouts and pages
  scaffolds/   # templates rendered directly by controllers
```

- A request to `/site/<path>` renders `pages/<path>.html`, or `pages/<path>/index.html`.
- Admin routes (`/site/admin/<path>`) render from `pages/admin/`.
- A page picks its layout with a [front matter](pages/front-matter.md) block. Without it, `layouts/main` is used.

## Template functions

Functions available in every template:

| Function | Purpose |
|----------|---------|
| `translate lang key [fallback]` | Translated text for `key` in `lang`. |
| `renderPost postId` | Renders the Markdown post `postId` to HTML. |
| `unescape text` | Outputs `text` as raw HTML. |

## Pages

- [Front matter](pages/front-matter.md): set a page's layout from the page itself.
