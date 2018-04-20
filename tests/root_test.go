package tests

import (
	"testing-in-golang/framework/swapi"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
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
