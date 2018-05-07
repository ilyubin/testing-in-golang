package tests

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"testing-in-golang/framework/swapi"
)

func Test_GetStartships_200(t *testing.T) {
	t.Parallel()
	ships := swapi.GetStarships()
	assert.Equal(t, ships.Count, 37)
}
