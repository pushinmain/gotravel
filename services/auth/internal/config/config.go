package config

import (
	"flag"
	"os"
	"time"

	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	Env      string     `yml:"env" env-default:"local"`
	TokenTTL string     `yml:"token_ttl" env-required:"1h"`
	Grpc     GrpcConfig `yml:"grpc"`
}

type GrpcConfig struct {
	Port    string        `yml:"port"`
	Timeout time.Duration `yml:"timeout"`
}

func MustLoad() *Config {
	path := fetchConfigPath()

	if path == "" {
		panic("config is empty")
	}

	if _, err := os.Stat(path); os.IsNotExist(err) {
		panic("config file does not exist: ")
	}

	var cfg Config

	if err := cleanenv.ReadConfig(path, &cfg); err != nil {
		panic("failed to read config")
	}

	return &cfg
}

// Fetches the config path in the local development environment.
// Example: go run ./path/to/entrypoint --config=./path/to/config/config.yaml
// For remote development, it is better to set the env variables explicitly.
func fetchConfigPath() string {
	var res string

	flag.StringVar(&res, "config", "", "path to config file")
	flag.Parse()

	if res == "" {
		res = os.Getenv("CONFIG_PATH")
	}
	return res
}
