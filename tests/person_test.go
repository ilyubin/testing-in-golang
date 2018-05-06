package tests

import (
	. "github.com/smartystreets/goconvey/convey"
	"net/http"
	"testing"
	"testing-in-golang/framework/swapi"
)

func Test_GetPerson(t *testing.T) {

	Convey("Handler GET /people/1", t, func() {

		Convey("When send request with personId = 1", func() {
			person := swapi.GetPerson(1)

			Convey("Then Luke Skywalker should received", func() {
				So(person.Name, ShouldEqual, "Luke Skywalker")
				So(person.HairColor, ShouldEqual, "blond")
				So(person.EyeColor, ShouldEqual, "blue")
			})
		})

		Convey("When send request with personId = 3", func() {
			person := swapi.GetPerson(3)

			Convey("Then R2-D2 should received", func() {
				So(person.Name, ShouldEqual, "R2-D2")
				So(person.EyeColor, ShouldEqual, "black")
			})
		})

		Convey("When send request with personId = 0", func() {
			err := swapi.GetPersonErr(0, http.StatusNotFound)

			Convey("Then error Not Found should received", func() {
				So(err.Detail, ShouldEqual, "Not found")
			})
		})
	})
}
