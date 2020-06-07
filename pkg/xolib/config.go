package xolib

// Config are the configuration env parameters
type Config struct {
	Host     string `env:"host"`
	Username string `env:"username"`
	Password string `env:"password"`
	Token    string `env:"token"`
}

// LoadConfig loads the Config from env
func LoadConfig() *Config {
	// TODO: Placeholder
	return nil
}
