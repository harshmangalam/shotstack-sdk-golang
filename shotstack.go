package shotstacksdkgolang

type Config struct {
	ApiKey      string // provide api_key from related environment
	Environment string // Environment can be "Sandbox" | "Production"
}

func New(apiKey, env string) *Config {

	return &Config{
		ApiKey:      apiKey,
		Environment: env,
	}
}
