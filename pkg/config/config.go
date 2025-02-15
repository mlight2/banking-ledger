package config

import (
	"log"

	"github.com/spf13/viper"
)

// Config structure
type Config struct {
	ServerPort   string `mapstructure:"SERVER_PORT"`
	PostgresURL  string `mapstructure:"POSTGRES_URL"`
	MongoURI     string `mapstructure:"MONGO_URI"`
	KafkaBrokers string `mapstructure:"KAFKA_BROKERS"`
}

// LoadConfig loads configuration from file
func LoadConfig() (*Config, error) {
	viper.SetConfigFile("./config.yml")
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		log.Fatal("Error reading config file", err)
		return nil, err
	}

	var config Config
	if err := viper.Unmarshal(&config); err != nil {
		log.Fatal("Error unmarshalling config", err)
		return nil, err
	}

	return &config, nil
}
