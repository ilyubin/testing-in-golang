package tests

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"testing-in-golang/framework/swapi"
)

func Test_GetFilms_200(t *testing.T) {
	t.Parallel()
	films := swapi.GetFilms()
	assert.Equal(t, films.Count, 7)
}
