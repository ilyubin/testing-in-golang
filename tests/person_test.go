package tests

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"ginkgo_pracrtice/framework/swapi"
	"net/http"
)

var _ = Describe("GET /people/1", func() {

	It("200 if person 1 is Luke Skywalker", func() {
		person := swapi.GetPerson(1)
		Ω(person.Name).Should(Equal("Luke Skywalker"))
	})


	It("404 if nonexistent personId", func() {
		swapi.GetPersonErr(0, http.StatusNotFound)
	})
})
