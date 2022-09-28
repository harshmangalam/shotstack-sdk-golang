package shotstacksdkgolang

type EnvType string

const (
	Prod  EnvType = "prod"
	Stage EnvType = "stage"
)

type Config struct {
	ApiKey string
	Env    EnvType
}

func NewConfig() *Config {
	return &Config{}
}

func (c *Config) SetApiKey(apiKey string) *Config {
	c.ApiKey = apiKey
	return c
}

func (c *Config) SetEnv(env EnvType) *Config {
	c.Env = env
	return c
}
