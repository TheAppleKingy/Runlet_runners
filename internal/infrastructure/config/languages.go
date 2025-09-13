package config

import (
	"log/slog"
	"os"

	"gopkg.in/yaml.v3"
)

type LanguageConfig struct {
	CompileCommandArgs []string `yaml:"compile"`
	RunCommandArgs     []string `yaml:"run"`
}

type languagesConfig struct {
	path      string                    `yaml:"-"`
	Languages map[string]LanguageConfig `yaml:"languages"`
}

func (cfg *languagesConfig) parse() {
	if cfg.path == "" {
		slog.Error("Unable to get runners config path from environment")
		os.Exit(1)
	}

	confData, err := os.ReadFile(cfg.path)
	if err != nil {
		slog.Error("Unable to read runners config file", "error", err)
		os.Exit(1)
	}

	if err := yaml.Unmarshal(confData, cfg); err != nil {
		slog.Error("Unable to parse runners config file", "error", err)
		os.Exit(1)
	}
}

var LaunguagesConfig = languagesConfig{
	path: os.Getenv("LANGS_CONF_PATH"),
}
