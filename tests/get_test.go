package swapi_tests

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"testing-in-golang/framework/httpbin"
)

var _ = Describe("GET /get", func() {
	It("200 and return params", func() {
		resp := httpbin.GetGet()
		Ω(resp.Headers.ContentType).Should(Equal("application/json"))
	})
})
