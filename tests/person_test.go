package tests

import (
	. "gopkg.in/check.v1"
	"net/http"
	"testing-in-golang/framework/swapi"
)

func (s *MySuite) Test_GetPerson_200_person_1(c *C) {
	person := swapi.GetPerson(1)
	c.Assert(person.Name, Equals, "Luke Skywalker")
	c.Assert(person.HairColor, Equals, "blond")
	c.Assert(person.EyeColor, Equals, "blue")
}

func (s *MySuite) Test_GetPerson_200_person_3(c *C) {
	person := swapi.GetPerson(3)
	c.Assert(person.Name, Equals, "R2-D2")
	c.Assert(person.EyeColor, Equals, "red")
}

func (s *MySuite) Test_GetPerson_404_if_nonexistent_personId(c *C) {
	err := swapi.GetPersonErr(0, http.StatusNotFound)
	c.Assert(err.Detail, Equals, "Not found")
}
