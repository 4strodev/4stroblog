package post

import (
	emoji "github.com/4strodev/go-markdown-emoji"
	"github.com/gomarkdown/markdown/html"
	"github.com/gomarkdown/markdown/parser"
	"github.com/microcosm-cc/bluemonday"
)

var renderPool *RenderPool
var sanitizePolicy = bluemonday.UGCPolicy()

func init() {
	renderPool = new(RenderPool)
	renderPool.New = func() any {
		htmlFlags := html.CommonFlags | html.HrefTargetBlank
		opts := html.RendererOptions{Flags: htmlFlags, RenderNodeHook: emoji.Renderer}
		return html.NewRenderer(opts)
	}
}

// RenderMarkdown renders and sanitizes markdown input converting it into
// save html
func RenderMarkdown(md []byte) []byte {
	extensions := parser.CommonExtensions | parser.AutoHeadingIDs | parser.NoEmptyLineBeforeBlock
	p := parser.NewWithExtensions(extensions)
	p.Opts = parser.Options{
		ParserHook: emoji.Parser,
	}
	doc := p.Parse(md)

	// Generate html from parsed markdown
	rawHtml := renderPool.Render(doc)

	// Sanitize generated html
	html := sanitizePolicy.SanitizeBytes(rawHtml)

	return html
}
