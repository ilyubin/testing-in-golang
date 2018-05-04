package tests

import (
	"github.com/corbym/gocrest/is"
	"github.com/corbym/gocrest/then"
	"net/http"
	"testing"
	"testing-in-golang/framework/swapi"
)

func Test_GetFilm_200_film_1(t *testing.T) {
	t.Parallel()
	film := swapi.GetFilm(1)
	then.AssertThat(t, film.Title, is.EqualTo("A New Hope"))
}

func Test_GetFilm_200_film_2(t *testing.T) {
	t.Parallel()
	film := swapi.GetFilm(2)
	then.AssertThat(t, film.Title, is.EqualTo("The Empire Strikes Back"))
}

func Test_GetFilm_404_if_nonexistent_filmId(t *testing.T) {
	t.Parallel()
	err := swapi.GetFilmErr(0, http.StatusNotFound)
	then.AssertThat(t, err.Detail, is.EqualTo("Not found"))
}
