package tests

import (
	"github.com/corbym/gocrest/is"
	"github.com/corbym/gocrest/then"
	"testing"
	"testing-in-golang/framework/swapi"
)

func Test_GetFilms_200(t *testing.T) {
	t.Parallel()
	films := swapi.GetFilms()
	then.AssertThat(t, films.Count, is.EqualTo(7))
}
