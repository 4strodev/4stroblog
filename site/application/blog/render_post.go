package blog

import (
	"fmt"
	"os"
	"path/filepath"
)

var baseFolder = "./uploads/posts"

func RenderPost(postId string) ([]byte, error) {
	fileName := fmt.Sprintf("%s.md", postId)
	filePath := filepath.Join(baseFolder, fileName)

	content, err := os.ReadFile(filePath)
	if err != nil {
		return nil, err
	}

	html := RenderMarkdown(content)

	return html, nil
}
