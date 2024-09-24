package main

import (
	configuration "github.com/Nikik0/dataCollectorBot/internal/config"
	"github.com/Nikik0/dataCollectorBot/internal/logger"
)

var (
	dbConnectionString = ""
	tgToken            = ""
)

func main() {

	logger.Info("Starting application")
	config, err := configuration.New()
	if err != nil {
		logger.Fatal("Failed to get config")
	}
	setConfigSettings(config)
}

func setConfigSettings(conif *configuration.Config) {
	dbConnectionString = conif.ConnectionStringDB
	tgToken = conif.Token
}
