package tests

import (
	. "github.com/smartystreets/goconvey/convey"
	"net/http"
	"testing"
	"testing-in-golang/framework/swapi"
)

func Test_GetFilm(t *testing.T) {

	Convey("Handler GET /films/1", t, func() {

		Convey("When send request with filmId = 1", func() {
			film := swapi.GetFilm(1)

			Convey("Then Luke Skywalker should received", func() {
				So(film.Title, ShouldEqual, "A New Hope")
			})
		})

		Convey("When send request with filmId = 3", func() {
			film := swapi.GetFilm(2)

			Convey("Then R2-D2 should received", func() {
				So(film.Title, ShouldEqual, "The Empire Strikes Back")
			})
		})

		Convey("When send request with filmId = 0", func() {
			err := swapi.GetFilmErr(0, http.StatusNotFound)

			Convey("Then error Not Found should received", func() {
				So(err.Detail, ShouldEqual, "Not found")
			})
		})
	})
}
