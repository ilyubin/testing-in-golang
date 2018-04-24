package swapi_tests

import (
	"github.com/onsi/ginkgo"
	"github.com/onsi/gomega"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"testing"

	"os"
	"github.com/onsi/ginkgo/reporters"
	"fmt"
	"github.com/onsi/ginkgo/config"
	"testing-in-golang/project_httpbin/framework/httpbin"
)

func TestApi(t *testing.T) {
	setupCfg()
	setupLog()
	setupApp()
	gomega.RegisterFailHandler(ginkgo.Fail)
	junitReporter := reporters.NewJUnitReporter(fmt.Sprintf("junit_%d.xml", config.GinkgoConfig.ParallelNode))
	ginkgo.RunSpecsWithDefaultAndCustomReporters(t, "Api", []ginkgo.Reporter{junitReporter})
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
	httpbin.Init()
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
