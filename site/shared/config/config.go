package config

import (
	"fmt"

	"github.com/knadh/koanf/parsers/toml"
	"github.com/knadh/koanf/providers/file"
	"github.com/knadh/koanf/v2"
)

type Config struct {
	JWK struct {
		Secret string `koanf:"secret"`
	} `koanf:"jwk"`
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
