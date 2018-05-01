package tests_testing_gocrest

import (
	"testing-in-golang/project_swapi/framework/swapi"
	"net/http"
	"testing"
	"github.com/corbym/gocrest/then"
	"github.com/corbym/gocrest/is"
)

func Test_GetPlanet_200(t *testing.T) {
	planet := swapi.GetPlanet(1)
	then.AssertThat(t, planet.Name, is.EqualTo("Tatooine"))
}

func Test_GetPlanet_404_if_nonexistent_planetId(t *testing.T) {
	err := swapi.GetPersonErr(0, http.StatusNotFound)
	then.AssertThat(t, err.Detail, is.EqualTo("Not found"))
}