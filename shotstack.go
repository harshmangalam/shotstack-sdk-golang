package shotstacksdkgolang

type Config struct {
	ApiKey string
	Env    string
}

func NewConfig() *Config {
	return &Config{}
}

func (c *Config) SetApiKey(apiKey string) *Config {
	c.ApiKey = apiKey
	return c
}

func (c *Config) SetEnv(env string) *Config {
	c.Env = env
	return c
}
