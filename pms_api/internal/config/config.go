package config

type Config struct {
	Port string
}

func NewConfig() *Config {
	return &Config{}
}

func (c *Config) GetAddress() string {
	return c.Port
}
