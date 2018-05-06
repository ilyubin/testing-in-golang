package tests

import (
	"github.com/corbym/gocrest/is"
	"github.com/corbym/gocrest/then"
	"net/http"
	"testing"
	"testing-in-golang/framework/swapi"
)

func Test_GetPerson_200_person_1(t *testing.T) {
	t.Parallel()
	person := swapi.GetPerson(1)
	then.AssertThat(t, person.Name, is.EqualTo("Luke Skywalker"))
	then.AssertThat(t, person.HairColor, is.EqualTo("blond"))
	then.AssertThat(t, person.EyeColor, is.EqualTo("blue"))
}

func Test_GetPerson_200_person_3(t *testing.T) {
	t.Parallel()
	person := swapi.GetPerson(3)
	then.AssertThat(t, person.Name, is.EqualTo("R2-D2"))
	then.AssertThat(t, person.EyeColor, is.EqualTo("red"))
}

func Test_GetPerson_404_if_nonexistent_personId(t *testing.T) {
	t.Parallel()
	err := swapi.GetPersonErr(0, http.StatusNotFound)
	then.AssertThat(t, err.Detail, is.EqualTo("Not found"))
}

func Test_GetPerson_200_person_1_te(t *testing.T) {
	t.Parallel()
	person := swapi.GetPerson(5)
	if person.Name != "Luke Skywalker" {
		t.Errorf("Expected person.Name \"Luke Skywalker\", but was %q.", person.Name)
	}
}
