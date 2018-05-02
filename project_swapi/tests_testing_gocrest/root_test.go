package tests_testing_gocrest

import (
	"github.com/corbym/gocrest/is"
	"github.com/corbym/gocrest/then"
	"testing"
	"testing-in-golang/project_swapi/framework/swapi"
)

func Test_GetRoot_200(t *testing.T) {
	t.Parallel()
	body := swapi.GetUrls()
	then.AssertThat(t, body.People, is.ValueContaining("http"))
}
