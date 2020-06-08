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

	return &Config{
		Host:     "10.50.1.182",
		Username: "arturo",
		Password: "Hydr0gen7892",
	}
}
