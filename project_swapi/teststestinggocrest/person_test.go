package teststestinggocrest

import (
	"github.com/corbym/gocrest/is"
	"github.com/corbym/gocrest/then"
	"net/http"
	"testing"
	"testing-in-golang/project_swapi/framework/swapi"
)

func Test_GetPerson_200(t *testing.T) {
	t.Parallel()
	person := swapi.GetPerson(1)
	then.AssertThat(t, person.Name, is.EqualTo("Luke Skywalker"))
	then.AssertThat(t, person.HairColor, is.EqualTo("blond"))
	then.AssertThat(t, person.EyeColor, is.EqualTo("blue"))
}

func Test_GetPerson_404_if_nonexistent_personId(t *testing.T) {
	t.Parallel()
	err := swapi.GetPersonErr(0, http.StatusNotFound)
	then.AssertThat(t, err.Detail, is.EqualTo("Not found"))
}
