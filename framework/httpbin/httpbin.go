package httpbin

import (
	"github.com/ilyubin/go-packages/http_client"
	"github.com/spf13/viper"
	"net/http"
	"testing-in-golang/framework/httpbin/model"
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
