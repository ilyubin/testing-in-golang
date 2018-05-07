package tests

import (
	"github.com/onsi/ginkgo"
	"github.com/onsi/gomega"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"testing"

	"fmt"
	"github.com/onsi/ginkgo/config"
	"github.com/onsi/ginkgo/reporters"
	"os"
	"testing-in-golang/framework/swapi"
)

func TestApi(t *testing.T) {
	setupCfg()
	setupLog()
	setupApp()
	gomega.RegisterFailHandler(ginkgo.Fail)
	node := config.GinkgoConfig.ParallelNode
	junitReporter := reporters.NewJUnitReporter(fmt.Sprintf("junit_%d.xml", node))
	ginkgo.RunSpecsWithDefaultAndCustomReporters(t, "Api", []ginkgo.Reporter{junitReporter})
}

func setupCfg() {
	viper.SetConfigName("config")
	viper.AddConfigPath(".")
	viper.AddConfigPath("..")

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
