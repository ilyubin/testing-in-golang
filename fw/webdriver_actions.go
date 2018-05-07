package fw

import (
	"github.com/sclevine/agouti"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

var wd *agouti.Page

func Init() {
	driver := agouti.Selenium()
	if err := driver.Start(); err != nil {
		panic(err)
	}
	instance, err := driver.NewPage(agouti.Browser(viper.GetString("driver.browser")))
	if err != nil {
		panic(err)
	}
	instance.Navigate(viper.GetString("yandex.base_url"))
	wd = instance
}

func find(css string) *agouti.Selection {
	element := wd.Find(css)
	return element
}

func Click(css string) {
	element := find(css)
	log.WithFields(log.Fields{
		"method": "Click",
		"css":    css,
	}).Info()
	element.Click()
}

func Type(css string, text string) {
	element := find(css)
	log.WithFields(log.Fields{
		"method": "Type",
		"css":    css,
		"text":   text,
	}).Info()
	element.Clear()
	element.SendKeys(text)
}

func Url() string {
	url, _ := wd.URL()
	log.WithFields(log.Fields{
		"method": "Url",
		"url":    url,
	}).Info()
	return url
}

func GetText(css string) string {
	text, _ := find(css).Text()
	log.WithFields(log.Fields{
		"method": "GetText",
		"css":    css,
		"text":   text,
	}).Info()
	return text
}
