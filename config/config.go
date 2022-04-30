package config

type Config struct {
	Host     string `env:"HOST, default=127.0.0.1"`
	Port     int    `env:"PORT, default=8080"`
	Database struct {
		Host     string `env:"DB_HOST, default=localhost"`
		Port     string `env:"DB_PORT, default=5432"`
		Name     string `env:"DB_NAME, required"`
		Username string `env:"DB_USERNAME, required"`
		Password string `env:"DB_PASSWORD, required"`
		SSL      string `env:"DB_SSL, default=disable"`
	}
}
