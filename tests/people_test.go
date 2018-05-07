package tests

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"testing-in-golang/framework/swapi"
)

func Test_GetPeople_200(t *testing.T) {
	t.Parallel()
	people := swapi.GetPeople()
	assert.Equal(t, people.Count, 87)
}

func Test_GetPeople_200_twice(t *testing.T) {
	t.Parallel()
	people1 := swapi.GetPeople()
	people2 := swapi.GetPeople()
	assert.Equal(t, people2, people1)
}
