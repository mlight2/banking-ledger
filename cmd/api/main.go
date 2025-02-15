package config

import (
	"log"

	"github.com/spf13/viper"
)

type Config struct {
	ServerPort  string `mapstructure:"SERVER_PORT"`
	PostgresURL string `mapstructure:"POSTGRES_URL"`
	MongoURI    string `mapstructure:"MONGO_URI"`
}

func LoadConfig() (*Config, error) {
	viper.SetConfigFile("./config.yml")
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("[ERROR] Failed to read config file: %v", err)
		return nil, err
	}

	var cfg Config
	if err := viper.Unmarshal(&cfg); err != nil {
		log.Fatalf("[ERROR] Failed to unmarshal config: %v", err)
		return nil, err
	}

	log.Println("[INFO] Configuration loaded successfully")
	return &cfg, nil
}
