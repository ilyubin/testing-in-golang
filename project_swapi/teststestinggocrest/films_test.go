package teststestinggocrest

import (
	"github.com/corbym/gocrest/has"
	"github.com/corbym/gocrest/then"
	"testing"
	"testing-in-golang/project_swapi/framework/swapi"
)

func Test_GetFilms_200(t *testing.T) {
	t.Parallel()
	films := swapi.GetFilms()
	then.AssertThat(t, films.Results, has.Length(7))
}
