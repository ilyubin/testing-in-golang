package httpbin

import (
	"testing-in-golang/core/http_client"
	"testing-in-golang/project_httpbin/framework/httpbin/model"
	"github.com/spf13/viper"
	"net/http"
)

var api http_client.HttpClient

func Init() {
	api = http_client.Init(viper.GetString("httpbin.base_url"))
}

func GetGet() (out model.GetResponse) {
	api.Get("/get", http.StatusOK, &out)
	return
}

func PostPost(in model.PostRequest) (out model.PostResponse) {
	api.Post("/post", in, http.StatusOK, &out)
	return
}
