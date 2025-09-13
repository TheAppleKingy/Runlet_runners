package config

type config interface {
	parse()
}

func LoadConfigs() {
	configs := []config{
		&LaunguagesConfig,
	}
	for _, cfg := range configs {
		cfg.parse()
	}
}
