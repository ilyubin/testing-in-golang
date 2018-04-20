package tests

import (
	"ginkgo_pracrtice/framework/swapi"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("GET /people", func() {
	It("200 and results should not be zero", func() {
		people := swapi.GerPeople()
		Ω(len(people.Results)).ShouldNot(BeZero())
	})
})
