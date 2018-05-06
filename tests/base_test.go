package tests

import (
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"os"
	"testing-in-golang/framework/swapi"
)

func init() {
	setupCfg()
	setupLog()
	setupApp()
}

func setupCfg() {
	viper.SetConfigName("config")
	viper.AddConfigPath("..")
	viper.AddConfigPath(".")

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
