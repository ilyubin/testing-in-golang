package tests

import (
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
	"testing-in-golang/framework/swapi"
)

func Test_GetPerson_200_person_1(t *testing.T) {
	t.Parallel()
	person := swapi.GetPerson(1)
	assert.Equal(t, person.Name, "Luke Skywalker")
	assert.Equal(t, person.HairColor, "blond")
	assert.Equal(t, person.EyeColor, "blue")
}

func Test_GetPerson_200_person_3(t *testing.T) {
	t.Parallel()
	person := swapi.GetPerson(3)
	assert.Equal(t, person.Name, "R2-D2")
	assert.Equal(t, person.EyeColor, "black")
}

func Test_GetPerson_404_if_nonexistent_personId(t *testing.T) {
	t.Parallel()
	err := swapi.GetPersonErr(0, http.StatusNotFound)
	assert.Equal(t, err.Detail, "Not found")
}
