package swapi

import (
	"fmt"
	"github.com/ilyubin/go-packages/http_client"
	"github.com/spf13/viper"
	"net/http"
	"testing-in-golang/framework/swapi/model"
)

var api http_client.HttpClient

func Init() {
	api = http_client.Init(viper.GetString("swapi.base_url"))
}

func GetUrls() (out model.RootResponse) {
	api.Get("/", http.StatusOK, &out)
	return
}

func GetPeople() (out model.PeopleResponse) {
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

func GetFilms() (out model.FilmsResponse) {
	api.Get("/films", http.StatusOK, &out)
	return
}

func GetFilm(num int) (out model.FilmResponse) {
	api.Get(fmt.Sprintf("/films/%d", num), http.StatusOK, &out)
	return
}

func GetFilmErr(num int, statusCode int) (out model.ErrorResponse) {
	api.Get(fmt.Sprintf("/films/%d", num), statusCode, &out)
	return
}

func GetStarships() (out model.StarshipsResponse) {
	api.Get("/starships", http.StatusOK, &out)
	return
}

func GetStarship(num int) (out model.StarshipResponse) {
	api.Get(fmt.Sprintf("/starships/%d", num), http.StatusOK, &out)
	return
}

func GetStarshipErr(num int, statusCode int) (out model.ErrorResponse) {
	api.Get(fmt.Sprintf("/starships/%d", num), statusCode, &out)
	return
}
