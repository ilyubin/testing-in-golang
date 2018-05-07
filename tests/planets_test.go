package tests

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"testing-in-golang/framework/swapi"
)

func Test_GetPlanets_200(t *testing.T) {
	t.Parallel()
	planets := swapi.GetPlanets()
	assert.Equal(t, planets.Count, 61)
}
