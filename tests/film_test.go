package tests

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/ginkgo/extensions/table"
	. "github.com/onsi/gomega"
	"net/http"
	"testing-in-golang/framework/swapi"
)

var _ = Describe("Handler GET /film/1", func() {

	It("should return 200 if filmId = 1", func() {
		film := swapi.GetFilm(1)
		Ω(film.Title).Should(Equal("A New Hope"))
	})

	It("should return 200 if filmId = 2", func() {
		film := swapi.GetFilm(2)
		Ω(film.Title).Should(Equal("The Empire Strikes Back"))
	})

	It("should return 404 if nonexistent filmId", func() {
		err := swapi.GetFilmErr(0, http.StatusNotFound)
		Ω(err.Detail).Should(Equal("Not found"))
	})

})

var _ = Describe("Handler GET /film/1", func() {
	DescribeTable("should return 200",
		func(filmId int, title string) {
			film := swapi.GetFilm(filmId)
			Ω(film.Title).Should(Equal(title))
		},
		Entry("film_1", 1, "A New Hope"),
		Entry("film_2", 2, "The Empire Strikes Back"),
		Entry("film_3", 3, "Return of the Jedi"),
	)
})
