package conf

import "github.com/caarlos0/env/v6"

// AppConfig presents app conf
type AppConfig struct {
	Port     string `env:"PORT" envDefault:"8003"`
	EnableDB string `env:"ENABLE_DB" envDefault:"false"`
	// ENV
	EnvName string `env:"ENV_NAME" envDefault:"dev"`
	//KAFKA
	KafkaHost string `env:"KAFKA_HOST" envDefault:"kafka"`
}

var config AppConfig

func LoadConfig() {
	_ = env.Parse(&config)
}

func GetConfig() AppConfig {
	return config
}
