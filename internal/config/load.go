package config

import (
	"os"
	"path/filepath"

	"github.com/spf13/viper"
)

type Config struct {
	Server   ServerConfig
	Database DatabaseConfig
	JWT      JWTConfig
}

func Load() (*Config, error) {
	v := viper.New()
	v.SetConfigType("yaml")

	configDir := "./config"

	files, err := os.ReadDir(configDir)
	if err != nil {
		return nil, err
	}

	first := true
	for _, file := range files {
		if filepath.Ext(file.Name()) == ".yaml" {
			name := file.Name()
			v.SetConfigFile(filepath.Join(configDir, name))

			if first {
				if err := v.ReadInConfig(); err != nil {
					return nil, err
				}
				first = false
			} else {
				if err := v.MergeInConfig(); err != nil {
					return nil, err
				}
			}
		}
	}

	var config Config
	if err := v.Unmarshal(&config); err != nil {
		return nil, err
	}

	return &config, nil
}
