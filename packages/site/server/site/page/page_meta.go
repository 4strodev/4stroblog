package page

import (
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"strings"

	"github.com/knadh/koanf/parsers/toml"
)

const (
	frontMatterDelimiter = "+++"
	commentStart         = "{{/*"
	commentEnd           = "*/}}"
)

// PageMeta holds the settings a page declares in its front matter.
type PageMeta struct {
	Layout string
	Title  string
}

// LoadPagesMeta reads the front matter of every page under viewsDir/pages.
// Keys are template names without extension, e.g. "pages/admin/post/new".
func LoadPagesMeta(viewsDir string) (map[string]PageMeta, error) {
	metas := map[string]PageMeta{}
	pagesDir := filepath.Join(viewsDir, "pages")

	err := filepath.WalkDir(pagesDir, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if d.IsDir() || filepath.Ext(path) != ".html" {
			return nil
		}

		content, err := os.ReadFile(path)
		if err != nil {
			return err
		}

		meta, err := parseFrontMatter(content)
		if err != nil {
			return fmt.Errorf("invalid front matter in %s: %w", path, err)
		}

		rel, err := filepath.Rel(viewsDir, path)
		if err != nil {
			return err
		}
		metas[strings.TrimSuffix(rel, ".html")] = meta
		return nil
	})

	return metas, err
}

// parseFrontMatter reads the TOML block between "+++" lines inside a
// leading template comment. Pages without it get an empty PageMeta.
func parseFrontMatter(content []byte) (PageMeta, error) {
	meta := PageMeta{}
	text := strings.TrimSpace(string(content))
	if !strings.HasPrefix(text, commentStart) {
		return meta, nil
	}

	end := strings.Index(text, commentEnd)
	if end == -1 {
		return meta, nil
	}
	comment := text[len(commentStart):end]

	parts := strings.SplitN(comment, frontMatterDelimiter, 3)
	if len(parts) < 3 {
		return meta, nil
	}

	values, err := toml.Parser().Unmarshal([]byte(parts[1]))
	if err != nil {
		return meta, err
	}

	meta.Layout, _ = values["layout"].(string)
	meta.Title, _ = values["title"].(string)
	return meta, nil
}
