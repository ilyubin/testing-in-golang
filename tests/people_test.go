package tests

import (
	"github.com/corbym/gocrest/is"
	"github.com/corbym/gocrest/then"
	"testing"
	"testing-in-golang/framework/swapi"
)

func Test_GetPeople_200(t *testing.T) {
	t.Parallel()
	people := swapi.GetPeople()
	then.AssertThat(t, people.Count, is.EqualTo(87))
}

func Test_GetPeople_200_twice(t *testing.T) {
	t.Parallel()
	people1 := swapi.GetPeople()
	people2 := swapi.GetPeople()
	then.AssertThat(t, people2, is.EqualTo(people1))
}
