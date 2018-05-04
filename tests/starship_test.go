package tests

import (
	"github.com/corbym/gocrest/is"
	"github.com/corbym/gocrest/then"
	"net/http"
	"testing"
	"testing-in-golang/framework/swapi"
)

func Test_GetStarship_200(t *testing.T) {
	t.Parallel()
	ship := swapi.GetStarship(9)
	then.AssertThat(t, ship.Name, is.EqualTo("Death Star"))
}

func Test_GetStarship_404_if_nonexistent_starshipId(t *testing.T) {
	t.Parallel()
	err := swapi.GetStarshipErr(0, http.StatusNotFound)
	then.AssertThat(t, err.Detail, is.EqualTo("Not found"))
}
