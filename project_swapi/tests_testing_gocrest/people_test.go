package tests_testing_gocrest

import (
	"testing-in-golang/project_swapi/framework/swapi"
	"testing"
	"github.com/corbym/gocrest/then"
	"github.com/corbym/gocrest/is"
)

func TestGetPeople(t *testing.T) {
	people := swapi.GerPeople()
	then.AssertThat(t, len(people.Results), is.GreaterThan(0))
}
