package swapi_tests

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"testing-in-golang/project_httpbin/framework/httpbin"
	"testing-in-golang/project_httpbin/framework/httpbin/model"
	"github.com/google/uuid"
	"fmt"
)

var _ = Describe("POST /post", func() {

	It("200 and return params", func() {
		req := model.PostRequest{
			IntField: 123,
			StringField: fmt.Sprintf("str+%s", uuid.New()),
		}
		resp := httpbin.PostPost(req)
		Ω(resp.Json.IntField).Should(Equal(req.IntField))
		Ω(resp.Json.StringField).Should(Equal(req.StringField))
	})

})
