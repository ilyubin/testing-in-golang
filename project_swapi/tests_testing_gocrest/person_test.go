package tests_testing_gocrest

import (
	"testing-in-golang/project_swapi/framework/swapi"
	"net/http"
	"testing"
	"github.com/corbym/gocrest/then"
	"github.com/corbym/gocrest/is"
)

func Test_GetPerson_200(t *testing.T) {
	person := swapi.GetPerson(1)
	then.AssertThat(t, person.Name, is.EqualTo("Luke Skywalker"))
}

func Test_GetPerson_404_if_nonexistent_personId(t *testing.T) {
	err := swapi.GetPersonErr(0, http.StatusNotFound)
	then.AssertThat(t, err.Detail, is.EqualTo("Not found"))
}