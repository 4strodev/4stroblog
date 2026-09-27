package post

import (
	"sync"

	"github.com/gomarkdown/markdown"
	"github.com/gomarkdown/markdown/ast"
	"github.com/gomarkdown/markdown/html"
)

type RenderPool struct {
	sync.Pool
}

func (p *RenderPool) get() *html.Renderer {
	return p.Pool.Get().(*html.Renderer)
}

func (p *RenderPool) put(r *html.Renderer) {
	p.Pool.Put(r)
}

func (p *RenderPool) Render(doc ast.Node) []byte {
	renderer := p.get()
	defer p.put(renderer)
	return markdown.Render(doc, renderer)
}
