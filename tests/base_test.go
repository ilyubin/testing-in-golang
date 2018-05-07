package tests

import (
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	. "gopkg.in/check.v1"
	"os"
	"testing"
	"testing-in-golang/framework/swapi"
)

func init() {
	setupCfg()
	setupLog()
	setupApp()
}

func Test(t *testing.T) { TestingT(t) }

type MySuite struct{}

var _ = Suite(&MySuite{})

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
