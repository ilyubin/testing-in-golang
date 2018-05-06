package tests

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"testing-in-golang/framework/swapi"
)

var _ = Describe("Handler GET /people", func() {

	It("should return 200 and count people", func() {
		people := swapi.GetPeople()
		Ω(people.Count).Should(Equal(87))
	})

})
