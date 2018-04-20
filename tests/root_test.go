package tests

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"ginkgo_pracrtice/framework/swapi"
)

var _ = Describe("GET /", func() {

	It("200 and should return urls", func() {
		body := swapi.GetUrls()
		Ω(body.People).ShouldNot(BeEmpty())
		Ω(body.Films).ShouldNot(BeEmpty())
		Ω(body.Planets).ShouldNot(BeEmpty())
		Ω(body.Species).ShouldNot(BeEmpty())
		Ω(body.StarShips).ShouldNot(BeEmpty())
	})

})
