package config

import (
	"github.com/jinzhu/configor"
)

// Config config info
type Config struct {
	MySQL struct {
		Name     string
		User     string
		Password string
		Host     string
	}
	Twitter struct {
		Token            string
		Secret           string
		RequestURI       string
		AuthorizationURI string
		TokenRequestURI  string
		CallbackURI      string
	}
}

// New load config.yaml
func New(file ...string) *Config {
	config := new(Config)

	if len(file) < 1 {
		file = append(file, "./config/config.yaml")
	}

	if err := configor.New(&configor.Config{Debug: false}).Load(config, file[0]); err != nil {
		panic(err)
	}

	return config
}
