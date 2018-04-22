package http_client

import (
	"encoding/json"
	"net/http"
	"strings"

	"testing-in-golang/framework/extensions"

	"github.com/google/uuid"
	. "github.com/onsi/gomega"
	log "github.com/sirupsen/logrus"
	"net/http/httputil"
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
		"handler": strings.Split(extensions.CallerName2(), "/")[2],
	}).Info("Begin request")

	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Set("X-Custom-Header", "myvalue")
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	response, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()

	dump, _ := httputil.DumpResponse(response, true)

	logRequest.WithFields(log.Fields{
		"status_code":   response.StatusCode,
		"response_body": strings.Split(string(dump), "\r\n\r\n")[1],
	}).Info("End request")

	return response
}

func decode(response *http.Response, model interface{}) {
	json.NewDecoder(response.Body).Decode(model)
}
