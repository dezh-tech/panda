package config

import (
	"os"

	"github.com/dezh-tech/geb/delivery/grpc"
	"github.com/dezh-tech/geb/delivery/http"
	"github.com/dezh-tech/geb/infrastructure/database"
	grpcclient "github.com/dezh-tech/geb/infrastructure/grpc_client"
	"github.com/dezh-tech/geb/infrastructure/redis"
	"github.com/dezh-tech/geb/pkg/logger"
	"github.com/joho/godotenv"
	"gopkg.in/yaml.v3"
)

// Config represents the configs used by relay and other concepts on system.
type Config struct {
	Environment string            `yaml:"environment"`
	GRPCClient  grpcclient.Config `yaml:"grpc_client"`
	Database    database.Config   `yaml:"database"`
	RedisConf   redis.Config      `yaml:"redis"`
	GRPCServer  grpc.Config       `yaml:"grpc_server"`
	HTTPServer  http.Config       `yaml:"http_server"`
	Logger      logger.Config     `yaml:"logger"`
}

// Load loads config from file and env.
func Load(path string) (*Config, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	config := &Config{}

	decoder := yaml.NewDecoder(file)

	if err := decoder.Decode(config); err != nil {
		return nil, err
	}

	if config.Environment != "prod" {
		if err := godotenv.Load(); err != nil {
			return nil, err
		}
	}

	config.Database.URI = os.Getenv("MONGO_URI")
	config.RedisConf.URI = os.Getenv("REDIS_URI")

	if err := config.basicCheck(); err != nil {
		return nil, err
	}

	return config, nil
}

// basicCheck validates the basic stuff in config.
func (c *Config) basicCheck() error {
	return nil
}
