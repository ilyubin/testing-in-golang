package swapi_tests

import (
	"fmt"
	"github.com/google/uuid"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"testing-in-golang/framework/httpbin"
	"testing-in-golang/framework/httpbin/model"
)

var _ = Describe("POST /post", func() {

	It("200 and return params", func() {
		req := model.PostRequest{
			IntField:    123,
			StringField: fmt.Sprintf("str+%s", uuid.New()),
		}
		resp := httpbin.PostPost(req)
		Ω(resp.Json.IntField).Should(Equal(req.IntField))
		Ω(resp.Json.StringField).Should(Equal(req.StringField))
	})

})
