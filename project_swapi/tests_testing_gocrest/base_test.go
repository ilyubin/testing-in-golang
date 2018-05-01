package tests_testing_gocrest

import (
	"testing-in-golang/project_swapi/framework/swapi"
	"github.com/spf13/viper"
	log "github.com/sirupsen/logrus"
	"os"
)

func init() {
	setupCfg()
	setupLog()
	setupApp()
}

func setupCfg() {
	viper.SetConfigName("config")
	viper.AddConfigPath("../..")

	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
}

func setupLog() {
	log.SetFormatter(&log.JSONFormatter{})
	log.SetOutput(os.Stdout)
	log.SetLevel(log.DebugLevel)
}

func setupApp() {
	swapi.Init()
}
