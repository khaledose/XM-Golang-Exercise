package config

import (
	"fmt"

	"github.com/tkanos/gonfig"
	"go.uber.org/zap"
)

type Configuration struct {
	Server   ServerConfigurations   `json:"server"`
	Database DatabaseConfigurations `json:"database"`
	Kafka    KafkaConfigurations    `json:"kafka"`
}

type ServerConfigurations struct {
	Port   int    `json:"port"`
	Secret string `json:"secret"`
}

type DatabaseConfigurations struct {
	Dsn         string `json:"dsn"`
	Pool        int    `json:"pool"`
	VerboseMode bool   `json:"verbose-mode"`
}

type KafkaConfigurations struct {
	SecurityProtocol string `json:"security-protocol"`
	SaslMechanism    string `json:"sasl-mechanisms"`
	Servers          string `json:"servers"`
	User             string `json:"user"`
	Pass             string `json:"pass"`
	ClientName       string `json:"client-name"`
}

func GetConfig(env string, logger *zap.SugaredLogger) Configuration {
	logger.Infof("Running on environment: %s", env)
	configuration := Configuration{}

	fileName := fmt.Sprintf("./resources/%s_config.json", env)
	err := gonfig.GetConf(fileName, &configuration)
	if err != nil {
		logger.Fatalf("Error while trying to read config with error: %s", err)
		panic(err)
	}

	return configuration
}
