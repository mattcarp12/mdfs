package config

import "github.com/caarlos0/env/v6"

type Config struct {
	RedisAddr       string `env:"REDIS_ADDR"`
	ChunkmasterAddr string `env:"CHUNKMASTER_ADDR"`
	ChunkDir        string `env:"CHUNK_DIR"`
}

func NewConfig() Config {
	var config Config
	if err := env.Parse(&config); err != nil {
		panic(err)
	}

	return config
}
