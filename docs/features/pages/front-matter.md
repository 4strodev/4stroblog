# Page front matter

A page template can declare settings for the router in a front matter block at the top of the file.
Today it sets the **layout**. It is the place to add future per-page settings.

## Embed a front matter

Put a TOML block between two `+++` lines, inside a Go template comment, as the **first thing** in the page:

```html
{{/*
+++
layout = "layouts/full-width"
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

All fields with their defaults:

```toml
# Layout template used to render the page (views/layouts/)
layout = "layouts/main"
```

Available layouts (`views/layouts/`):

- `layouts/main`: content limited to a reading width.
- `layouts/full-width`: content uses the whole screen. For editing pages.

## Rules

- The block must be at the top of the file. A block anywhere else is ignored.
- A page without a block uses the defaults.
- Invalid TOML stops the app at startup with an error naming the file.
- Front matter is read once at startup. Restart the app after changing it (`task dev` does it on file changes).

## Use it in templates

Pages rendered by the page controller receive the values as `.Meta`:

```html
<p>Layout: {{ .Meta.Layout }}</p>
```

Pages rendered by other controllers (e.g. the blog controller) do not get `.Meta`.

## Add a new field

1. Add the field to `PageMeta` in `packages/site/server/site/page/page_meta.go`.
2. Read it from the parsed values in `parseFrontMatter`.
3. Use it in the controller (`page_controller.go`) or in templates via `.Meta`.
4. Add a case to `page_meta_test.go`.
