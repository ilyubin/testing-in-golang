package swapi_tests

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"testing-in-golang/project_httpbin/framework/httpbin"
	"testing-in-golang/project_httpbin/framework/httpbin/model"
)

var _ = Describe("POST /post", func() {
	It("200 and return params", func() {
		req := model.PostRequest{IntField: 123, StringField: "xyz"}
		resp := httpbin.PostPost(req)
		Ω(resp.Headers.ContentType).Should(Equal("application/json"))
	})
})
