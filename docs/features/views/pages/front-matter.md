# Page front matter

A page template can declare settings for the router in a front matter block at the top of the file.
Today it sets the **layout** and the **title**. It is the place to add future per-page settings.

## Embed a front matter

Put a TOML block between two `+++` lines, inside a Go template comment, as the **first thing** in the page:

```html
{{/*
+++
layout = "layouts/full-width"
title = "New post"
+++
*/}}
<form>
  ...
</form>
```

The template engine ignores the comment, so the block never reaches the HTML.

## Fields

| Field | Type | Default | Purpose |
|-------|------|---------|---------|
| `layout` | string | `layouts/main` | Layout template used to render the page. |
| `title` | string | `site.title` from config | Browser tab title. |

All fields with their defaults:

```toml
# Layout template used to render the page (views/layouts/)
layout = "layouts/main"

# Browser tab title, available to templates as .Meta.Title
# Defaults to site.title in config/config.toml
title = "4strodev"
```

Available layouts (`views/layouts/`):

- `layouts/main`: content limited to a reading width.
- `layouts/full-width`: content uses the whole screen. For editing pages.

## Rules

- The block must be at the top of the file. A block anywhere else is ignored.
- A page without a block, or a field missing or empty, uses the defaults. Templates never need a fallback.
- Invalid TOML stops the app at startup with an error naming the file.
- Front matter is read once at startup. Restart the app after changing it (`task dev` does it on file changes).

## Use it in templates

Pages rendered by the page controller receive the values as `.Meta`. Layouts and components included with `.` get it too:

```html
<title>{{ .Meta.Title }}</title>
```

Controllers that render templates on their own must pass `Meta` too (the blog controller passes the post title).

## Add a new field

1. Add the field to `PageMeta` in `packages/site/server/site/page/page_meta.go`.
   Set its default in `defaultMeta` (`page_controller.go`), from config if it should be configurable.
2. Read it from the parsed values in `parseFrontMatter`.
3. Use it in the controller (`page_controller.go`) or in templates via `.Meta`.
4. Add a case to `page_meta_test.go`.
