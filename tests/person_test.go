package tests

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"net/http"
	"testing-in-golang/framework/swapi"
)

var _ = Describe("Handler GET /people/1", func() {
	It("should return 200 if personId = 1", func() {
		person := swapi.GetPerson(1)
		Ω(person.Name).Should(Equal("Luke Skywalker"))
		Ω(person.HairColor).Should(Equal("blond"))
		Ω(person.EyeColor).Should(Equal("blue"))
	})

	It("should return 200 if personId = 3", func() {
		person := swapi.GetPerson(3)
		Ω(person.Name).Should(Equal("R2-D2"))
		Ω(person.EyeColor).Should(Equal("red"))
	})

	It("should return 404 if nonexistent personId", func() {
		err := swapi.GetPersonErr(0, http.StatusNotFound)
		Ω(err.Detail).Should(Equal("Not found"))
	})
})
