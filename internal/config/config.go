package config

import (
	"github.com/Nikik0/dataCollectorBot/internal/logger"
	"github.com/pkg/errors"
	"gopkg.in/yaml.v3"
	"os"
)

const configFile = "config.yaml"

type Config struct {
	Token              string `yaml:"token"`              // Токен бота в телеграме.
	ConnectionStringDB string `yaml:"ConnectionStringDB"` // Строка подключения в базе данных.
}

func New() (*Config, error) {
	s := &Config{}

	rawYAML, err := os.ReadFile(configFile)
	if err != nil {
		logger.Error("Ошибка reading config file", "err", err)
		return nil, errors.Wrap(err, "reading config file")
	}

	err = yaml.Unmarshal(rawYAML, &s)
	if err != nil {
		logger.Error("Ошибка parsing yaml", "err", err)
		return nil, errors.Wrap(err, "parsing yaml")
	}

	return s, nil
}

func getToken(c *Config) string {
	return c.Token
}
