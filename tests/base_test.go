package tests

import (
	"github.com/sirupsen/logrus"
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
	logrus.SetFormatter(&logrus.JSONFormatter{})
	logrus.SetOutput(os.Stdout)
	logrus.SetLevel(logrus.DebugLevel)
}

func setupApp() {
	swapi.Init()
}
