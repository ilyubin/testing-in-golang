package http_client

import (
	"encoding/json"
	"net/http"
	"strings"

	"testing-in-golang/core/extensions"

	"github.com/google/uuid"
	log "github.com/sirupsen/logrus"
	"net/http/httputil"
	"bytes"
	"github.com/ahmetb/go-linq"
)

type HttpClient struct {
	BaseUrl string
}

func Init(baseUrl string) HttpClient {
	c := new(HttpClient)
	c.BaseUrl = baseUrl
	return *c
}

func (c *HttpClient) Get(path string, statusCode int, out interface{}) {
	response := c.get(path)
	checkStatusCode(response.StatusCode, statusCode)
	decode(response, &out)
}

func (c *HttpClient) Post(path string, in interface{}, statusCode int, out interface{}) {
	response := c.post(path, in)
	checkStatusCode(response.StatusCode, statusCode)
	decode(response, &out)
}

func checkStatusCode(actualStatusCode int, expectedStatusCode int) {
	if actualStatusCode != expectedStatusCode {
		log.Fatalf("Status code is %d, expected %d", actualStatusCode, expectedStatusCode)
	}
}

func (c *HttpClient) get(path string) *http.Response {
	url := c.BaseUrl + path

	logRequest := log.WithFields(log.Fields{"request_id": uuid.New()})

	logRequest.WithFields(log.Fields{
		"method":  "GET",
		"url":     url,
		"handler": getHandlerName(extensions.GrandParentalFunctionName()),
	}).Info("Begin request")

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		panic(err)
	}

	//req.Header.Set("X-Custom-Header", "myvalue")
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	response, err := client.Do(req)
	if err != nil {
		panic(err)
	}

	defer response.Body.Close()

	dump, err := httputil.DumpResponse(response, true)
	if err != nil {
		panic(err)
	}

	logRequest.WithFields(log.Fields{
		"status_code":   response.StatusCode,
		"response_body": strings.Split(string(dump), "\r\n\r\n")[1],
	}).Info("End request")

	return response
}

func (c *HttpClient) post(path string, data interface{}) *http.Response {
	url := c.BaseUrl + path
	b := new(bytes.Buffer)
	json.NewEncoder(b).Encode(data)

	logRequest := log.WithFields(log.Fields{"request_id": uuid.New()})

	logRequest.WithFields(log.Fields{
		"method":  "POST",
		"url":     url,
		"data":    data,
		"handler": getHandlerName(extensions.GrandParentalFunctionName()),
	}).Info("Begin request")

	req, err := http.NewRequest("POST", url, b)
	if err != nil {
		panic(err)
	}

	//req.Header.Set("X-Custom-Header", "myvalue")
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	response, err := client.Do(req)
	if err != nil {
		panic(err)
	}

	defer response.Body.Close()

	dump, err := httputil.DumpResponse(response, true)
	if err != nil {
		panic(err)
	}

	logRequest.WithFields(log.Fields{
		"status_code":   response.StatusCode,
		"response_body": strings.Split(string(dump), "\r\n\r\n")[1],
	}).Info("End request")

	return response
}

func decode(response *http.Response, model interface{}) {
	json.NewDecoder(response.Body).Decode(model)
}

func getHandlerName(name string) interface{} {
	return linq.From(strings.Split(name, "/")).Last()
}
