package tests

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"testing"

	"os"
)

func TestApi(t *testing.T) {
	setupCfg()
	setupLog()
	RegisterFailHandler(Fail)
	RunSpecs(t, "Api")
}

func setupCfg() {
	viper.SetDefault("ContentDir", "content")
}

func setupLog() {
	log.SetFormatter(&log.JSONFormatter{})
	log.SetOutput(os.Stdout)
	log.SetLevel(log.DebugLevel)
}

var _ = BeforeSuite(func() {
	log.WithFields(log.Fields{
		"suite": "before",
	}).Info()
})

var _ = AfterSuite(func() {
	log.WithFields(log.Fields{
		"suite": "after",
	}).Info()
})
