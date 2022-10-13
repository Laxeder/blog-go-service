package configs

var cfg *config

type config struct {
	API APIConfig
}

type APIConfig struct {
	DEV_PORT string
	PORT     string
	ENV      string
}

func Load() {
	cfg = new(config)

	cfg.API = APIConfig{
		ENV:      "development",
		DEV_PORT: ":3000",
		PORT:     ":9000",
	}
}

func GetAPI() APIConfig {
	return cfg.API
}

func GetAPIPort() string {
	if cfg.API.ENV == "production" {
		return cfg.API.PORT
	} else {
		return cfg.API.DEV_PORT
	}
}
