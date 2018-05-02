package tests_testing_gocrest

import (
	"github.com/corbym/gocrest/is"
	"github.com/corbym/gocrest/then"
	"net/http"
	"testing"
	"testing-in-golang/project_swapi/framework/swapi"
)

func Test_GetPlanet_200(t *testing.T) {
	t.Parallel()
	planet := swapi.GetPlanet(1)
	then.AssertThat(t, planet.Name, is.EqualTo("Tatooine"))
}

func Test_GetPlanet_404_if_nonexistent_planetId(t *testing.T) {
	t.Parallel()
	err := swapi.GetPersonErr(0, http.StatusNotFound)
	then.AssertThat(t, err.Detail, is.EqualTo("Not found"))
}

func Test_GetPlanet_twice(t *testing.T) {
	t.Parallel()
	planet1 := swapi.GetPlanet(2)
	planet2 := swapi.GetPlanet(3)
	then.AssertThat(t, planet1, is.Not(is.EqualTo(planet2)))
}
