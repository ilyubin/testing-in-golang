package tests

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"ginkgo_pracrtice/framework/swapi"
)

var _ = Describe("GET /people", func() {
	It("200 and results should not be zero", func() {
		people := swapi.GerPeople()
		Ω(len(people.Results)).ShouldNot(BeZero())
	})
})
