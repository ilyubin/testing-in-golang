package tests

import (
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
	"testing-in-golang/framework/swapi"
)

func Test_GetPlanet_200(t *testing.T) {
	t.Parallel()
	planet := swapi.GetPlanet(1)
	assert.Equal(t, planet.Name, "Tatooine")
}

func Test_GetPlanet_404_if_nonexistent_planetId(t *testing.T) {
	t.Parallel()
	err := swapi.GetPlanetErr(0, http.StatusNotFound)
	assert.Equal(t, err.Detail, "Not found")
}

func Test_GetPlanet_twice(t *testing.T) {
	t.Parallel()
	planet1 := swapi.GetPlanet(2)
	planet2 := swapi.GetPlanet(3)
	assert.NotEqual(t, planet1, planet2)
}
