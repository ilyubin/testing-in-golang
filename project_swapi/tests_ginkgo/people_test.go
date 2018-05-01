package swapi_tests

import (
	"testing-in-golang/project_swapi/framework/swapi"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("GET /people", func() {
	It("200 and results should not be zero", func() {
		people := swapi.GerPeople()
		Ω(len(people.Results)).ShouldNot(BeZero())
	})
})
