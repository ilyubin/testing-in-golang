package teststestinggocrest

import (
	"github.com/corbym/gocrest/has"
	"github.com/corbym/gocrest/then"
	"testing"
	"testing-in-golang/project_swapi/framework/swapi"
)

func Test_GetPlanets_200(t *testing.T) {
	t.Parallel()
	planets := swapi.GetPlanets()
	then.AssertThat(t, planets.Results, has.Length(10))
}
