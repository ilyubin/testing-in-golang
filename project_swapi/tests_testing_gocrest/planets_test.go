package tests_testing_gocrest

import (
	"testing"
	"testing-in-golang/project_swapi/framework/swapi"
	"github.com/corbym/gocrest/then"
	"github.com/corbym/gocrest/has"
)

func Test_GetPlanets_200(t *testing.T) {
	planets := swapi.GetPlanets()
	then.AssertThat(t, planets.Results, has.Length(10))
}
