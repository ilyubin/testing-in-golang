package swapi

import (
	"fmt"
	"testing-in-golang/framework/http_client"
	"testing-in-golang/framework/swapi/model"
	"github.com/spf13/viper"
)

var api http_client.HttpClient

func Init() {
	api = http_client.Init(viper.GetString("swapi.base_url"))
}

func GetUrls() (out model.RootResponse) {
	api.GetOk("/", &out)
	return
}

func GerPeople() (out model.PeopleResponse) {
	api.GetOk("/people", &out)
	return
}

func GetPerson(num int) (out model.PersonResponse) {
	api.GetOk(fmt.Sprintf("/people/%d", num), &out)
	return
}

func GetPersonErr(num int, statusCode int) {
	api.GetErr(fmt.Sprintf("/people/%d", num), statusCode)
}
