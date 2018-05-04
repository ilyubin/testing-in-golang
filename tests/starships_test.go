package tests

import (
	"github.com/corbym/gocrest/is"
	"github.com/corbym/gocrest/then"
	"testing"
	"testing-in-golang/framework/swapi"
)

func Test_GetStartships_200(t *testing.T) {
	t.Parallel()
	ships := swapi.GetStarships()
	then.AssertThat(t, ships.Count, is.EqualTo(37))
}
