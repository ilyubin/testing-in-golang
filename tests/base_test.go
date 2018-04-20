package tests

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	log "github.com/sirupsen/logrus"
	"testing"

	"os"
)

func TestApi(t *testing.T) {
	setupLog()
	RegisterFailHandler(Fail)
	RunSpecs(t, "Api")
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
