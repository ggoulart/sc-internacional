package mongodb

import (
	"github.com/caarlos0/env/v11"
)

type Config struct {
	MongoURI string `env:"MONGO_URI,required"`
	DBName   string `env:"DB_NAME,required"`
}

func NewConfig() (*Config, error) {
	cfg := Config{}
	if err := env.Parse(&cfg); err != nil {
		return nil, err
	}
	return &cfg, nil
}
