package models

type Config struct {
	HTTP HTTPConfiguration
}

type HTTPConfiguration struct {
	Host string `env:"HTTP_HOST,default=localhost"`
	Port string `env:"HTTP_PORT,default=8080"`
}
