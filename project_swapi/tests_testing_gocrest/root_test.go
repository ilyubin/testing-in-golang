package tests_testing_gocrest

import (
	"testing-in-golang/project_swapi/framework/swapi"
	"testing"
	"github.com/corbym/gocrest/then"
	"github.com/corbym/gocrest/is"
)

func Test_GetRoot_200(t *testing.T) {
	body := swapi.GetUrls()
	then.AssertThat(t, body.People, is.ValueContaining("http"))
}