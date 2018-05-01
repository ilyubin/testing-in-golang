package swapi

import (
	"fmt"
	"testing-in-golang/core/http_client"
	"testing-in-golang/project_swapi/framework/swapi/model"
	"github.com/spf13/viper"
	"net/http"
)

var api http_client.HttpClient

func Init() {
	api = http_client.Init(viper.GetString("swapi.base_url"))
}

func GetUrls() (out model.RootResponse) {
	api.Get("/", http.StatusOK, &out)
	return
}

func GerPeople() (out model.PeopleResponse) {
	api.Get("/people", http.StatusOK, &out)
	return
}

func GetPerson(num int) (out model.PersonResponse) {
	api.Get(fmt.Sprintf("/people/%d", num), http.StatusOK, &out)
	return
}

func GetPersonErr(num int, statusCode int) (out model.ErrorResponse) {
	api.Get(fmt.Sprintf("/people/%d", num), statusCode, &out)
	return
}

func GetPlanets() (out model.PlanetsResponse) {
	api.Get("/planets", http.StatusOK, &out)
	return
}

func GetPlanet(num int) (out model.PlanetResponse) {
	api.Get(fmt.Sprintf("/planets/%d", num), http.StatusOK, &out)
	return
}

func GetPlanetErr(num int, statusCode int) (out model.ErrorResponse) {
	api.Get(fmt.Sprintf("/planets/%d", num), statusCode, &out)
	return
}