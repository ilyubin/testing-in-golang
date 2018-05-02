package swapi_tests

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"testing-in-golang/project_swapi/framework/swapi"
)

var _ = Describe("GET /people", func() {
	It("200 and results should not be zero", func() {
		people := swapi.GetPeople()
		Ω(len(people.Results)).ShouldNot(BeZero())
	})
})
