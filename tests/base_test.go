package tests

import (
	"github.com/onsi/ginkgo"
	"github.com/onsi/gomega"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"testing"

	"os"
	"fmt"
	"testing-in-golang/framework/swapi"
	"github.com/onsi/ginkgo/reporters"
)

func TestApi(t *testing.T) {
	setupCfg()
	setupLog()
	setupApp()
	gomega.RegisterFailHandler(ginkgo.Fail)
	junitReporter := reporters.NewJUnitReporter("../junit.xml")
	ginkgo.RunSpecsWithDefaultAndCustomReporters(t, "Api", []ginkgo.Reporter{junitReporter})
}

func setupCfg() {
	viper.SetConfigName("config")
	viper.AddConfigPath("..")

	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
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

var _ = ginkgo.BeforeSuite(func() {
	log.WithFields(log.Fields{
		"suite": "before",
	}).Info()
})

var _ = ginkgo.AfterSuite(func() {
	log.WithFields(log.Fields{
		"suite": "after",
	}).Info()
})
