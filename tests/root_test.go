package tests

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"testing-in-golang/framework/swapi"
)

var _ = Describe("Handler GET /", func() {

	It("should return 200 and urls", func() {
		body := swapi.GetUrls()
		Ω(body.People).ShouldNot(BeEmpty())
		Ω(body.Films).ShouldNot(BeEmpty())
		Ω(body.Planets).ShouldNot(BeEmpty())
		Ω(body.Species).ShouldNot(BeEmpty())
		Ω(body.StarShips).ShouldNot(BeEmpty())
	})

})
