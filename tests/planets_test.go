package tests

import (
	"github.com/corbym/gocrest/is"
	"github.com/corbym/gocrest/then"
	"testing"
	"testing-in-golang/framework/swapi"
)

func Test_GetPlanets_200(t *testing.T) {
	t.Parallel()
	planets := swapi.GetPlanets()
	then.AssertThat(t, planets.Count, is.EqualTo(61))
}
