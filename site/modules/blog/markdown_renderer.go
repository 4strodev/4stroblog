package blog

import (
	"github.com/gomarkdown/markdown"
	"github.com/gomarkdown/markdown/html"
	"github.com/gomarkdown/markdown/parser"
	emoji "github.com/4strodev/go-markdown-emoji"
	"github.com/microcosm-cc/bluemonday"
)

// RenderMarkdown renders and sanitizes markdown input converting it into
// save html
func RenderMarkdown(md []byte) []byte {
	extensions := parser.CommonExtensions | parser.AutoHeadingIDs | parser.NoEmptyLineBeforeBlock
	p := parser.NewWithExtensions(extensions)
	p.Opts = parser.Options{
		ParserHook: emoji.Parser,
	}
	doc := p.Parse(md)
	htmlFlags := html.CommonFlags | html.HrefTargetBlank
	opts := html.RendererOptions{Flags: htmlFlags, RenderNodeHook: emoji.Renderer}
	renderer := html.NewRenderer(opts)

	rawHtml := markdown.Render(doc, renderer)
	html := bluemonday.UGCPolicy().SanitizeBytes(rawHtml)
	return html
}
