package http_client

import (
	"encoding/json"
	"net/http"
	"strings"

	"ginkgo_pracrtice/framework/extensions"

	"github.com/ahmetb/go-linq"
	"github.com/ddliu/go-httpclient"
	"github.com/google/uuid"
	. "github.com/onsi/gomega"
	log "github.com/sirupsen/logrus"
)

const BaseUrl = "https://swapi.co/api"

func GetOk(path string, model interface{}) {
	response := get(path)
	Ω(response.StatusCode).Should(Equal(http.StatusOK))
	decode(response, &model)
}

func GetErr(path string, statusCode int) {
	response := get(path)
	Ω(response.StatusCode).Should(Equal(statusCode))
}

func get(path string) *http.Response {
	url := BaseUrl + path

	logRequest := log.WithFields(log.Fields{"request_id": uuid.New()})

	logRequest.WithFields(log.Fields{
		"method":  "GET",
		"url":     url,
		"handler": linq.From(strings.Split(extensions.CallerName2(), "/")).Last(),
	}).Info("Begin request")

	httpclient.Defaults(httpclient.Map{
		httpclient.OPT_USERAGENT: "my awesome http client",
		"Accept-Language":        "en-us",
	})

	response, err1 := httpclient.Get(url)
	//response2 := &httpclient.Response{}
	//response2 := response.Response
	//
	//body, err2 := response.ToString()

	logRequest.WithFields(log.Fields{
		"error_http_client": err1,
		//"error_to_string": err2,
		"status_code": response.StatusCode,
		//"response_body": body,
	}).Info("End request")

	return response.Response
}

func decode(response *http.Response, model interface{}) {
	json.NewDecoder(response.Body).Decode(model)
}
