package i18n

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/knadh/koanf/parsers/toml"
	"github.com/knadh/koanf/providers/file"
	"github.com/knadh/koanf/v2"
	"github.com/tidwall/gjson"
	"golang.org/x/text/language"
)

type TranslationService struct {
	translationFiles map[string][]byte
}

func NewTranslationsService(folder string) (*TranslationService, error) {
	log.Println("creating translation service")
	dirEntries, err := os.ReadDir(folder)
	if err != nil {
		return nil, err
	}

	service := &TranslationService{
		translationFiles: make(map[string][]byte),
	}

	for _, file := range dirEntries {
		if file.IsDir() {
			continue
		}

		content, err := loadFile(filepath.Join(folder, file.Name()))
		if err != nil {
			return nil, err
		}

		lang := strings.Split(file.Name(), ".")[0]
		service.translationFiles[lang] = content
	}

	return service, nil
}

func loadFile(path string) ([]byte, error) {
	k := koanf.New(".")
	parser := toml.Parser()

	if err := k.Load(file.Provider(path), parser); err != nil {
		return nil, fmt.Errorf("cannot load translation: %w", err)
	}

	content := map[string]any{}
	if err := k.Unmarshal("", &content); err != nil {
		return nil, fmt.Errorf("cannot unmarshall translation: %w", err)
	}

	return json.Marshal(content)
}

func (t *TranslationService) Translate(lang, key string) string {
	if _, ok := t.translationFiles[lang]; !ok {
		return ""
	}
	return gjson.Get(string(t.translationFiles[lang]), key).String()
}

func (t *TranslationService) TranslateOr(lang, key string, fallback string) string {
	if _, ok := t.translationFiles[lang]; !ok {
		return ""
	}
	result := gjson.Get(string(t.translationFiles[lang]), key)
	if !result.Exists() {
		return fallback
	}
	return result.String()
}
