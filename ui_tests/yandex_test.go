package ui_tests

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/sclevine/agouti"
)

var _ = Describe("Browser", func() {

	It("ya.ru", func() {
		driver := agouti.Selenium()
		if err := driver.Start(); err != nil {
			panic(err)
		}
		page, err := driver.NewPage(agouti.Browser("chrome"))
		if err != nil {
			panic(err)
		}

		page.Navigate("https://ya.ru/")

		yaUrl, _ := page.URL()

		Ω(yaUrl).Should(Equal("https://ya.ru/"))

		page.Find(".input__input").SendKeys("Agouti")
		page.Find(".search2__button button").Click()

	})

})
