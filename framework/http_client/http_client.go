package http_client

import (
	"encoding/json"
	"net/http"
	"strings"

	"testing-in-golang/framework/extensions"

	"github.com/ahmetb/go-linq"
	"github.com/ddliu/go-httpclient"
	"github.com/google/uuid"
	. "github.com/onsi/gomega"
	log "github.com/sirupsen/logrus"
)

type HttpClient struct {
	BaseUrl string
}

func Init(baseUrl string) HttpClient  {
	c := new(HttpClient)
	c.BaseUrl = baseUrl
	return *c
}

func (c *HttpClient) GetOk(path string, model interface{}) {
	response := c.get(path)
	Ω(response.StatusCode).Should(Equal(http.StatusOK))
	decode(response, &model)
}

func (c *HttpClient) GetErr(path string, statusCode int) {
	response := c.get(path)
	Ω(response.StatusCode).Should(Equal(statusCode))
}

func (c *HttpClient) get(path string) *http.Response {
	url := c.BaseUrl + path

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
		"status_code":   response.StatusCode,
		"response_body": response.Response.Body,
	}).Info("End request")

	return response.Response
}

func decode(response *http.Response, model interface{}) {
	json.NewDecoder(response.Body).Decode(model)
}
