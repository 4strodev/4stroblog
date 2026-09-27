package config

import (
	"fmt"

	"github.com/knadh/koanf/parsers/toml"
	"github.com/knadh/koanf/providers/file"
	"github.com/knadh/koanf/v2"
)

type SiteConfig struct {
	AdminEmail string `koanf:"admin_email" validate:"required"`
}

type EnvironmentConfig struct {
	// GodMode allows you to access to all restricted
	// sections of the site. Use it to test integrations and
	// features hidden to public users
	GodMode bool `koanf:"god_mode"`
}

type AuthConfig struct {
	JWK struct {
		Secret string `koanf:"secret"`
	} `koanf:"jwk"`
}

type DbConfig struct {
	Sqlite struct {
		Path string `koanf:"path"`
	} `koanf:"sqlite"`
}

type StorageConfig struct {
	S3 struct {
		Url    string `koanf:"url"`
		Bucket string `koanf:"bucket"`
	} `koanf:"s3"`
}

type I18nConfig struct {
	Folder  string `koanf:"folder"`
	Default string `koanf:"default"`
}

type ViewsConfig struct {
	// Folder with the html templates (pages, layouts, components)
	Folder string `koanf:"folder"`
}

type Config struct {
	Site        SiteConfig        `koanf:"site"`
	Environment EnvironmentConfig `koanf:"environment"`
	Auth        AuthConfig        `koanf:"auth"`
	Db          DbConfig          `koanf:"db"`
	Storage     StorageConfig     `koanf:"storage"`
	I18n        I18nConfig        `koanf:"i18n"`
	Views       ViewsConfig       `koanf:"views"`
}

var config Config
var loaded bool

func GetConfig() (Config, error) {
	var err error
	if !loaded {
		config, err = loadConfig()
		if err != nil {
			return config, err
		}
		loaded = true
	}
	return config, nil
}

func loadConfig() (Config, error) {
	k := koanf.New(".")
	parser := toml.Parser()
	config := Config{}

	if err := k.Load(file.Provider("config/config.toml"), parser); err != nil {
		return config, fmt.Errorf("cannot load config: %w", err)
	}

	if err := k.Unmarshal("", &config); err != nil {
		return config, fmt.Errorf("cannot unmarshall config: %w", err)
	}

	return config, nil
}
