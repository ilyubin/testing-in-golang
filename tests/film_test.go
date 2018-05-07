package tests

import (
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
	"testing-in-golang/framework/swapi"
)

func Test_GetFilm_200(t *testing.T) {
	t.Parallel()
	cases := []struct {
		FilmId int
		Title  string
	}{
		{1, "A New Hope"},
		{2, "The Empire Strikes Back"},
		{3, "Return of the Jedi"},
		{4, "The Phantom Menace"},
		{5, "Attack of the Clones"},
		{6, "Revenge of the Sith"},
		{7, "The Force Awakens"},
	}
	for _, c := range cases {
		film := swapi.GetFilm(c.FilmId)
		assert.Equal(t, film.Title, c.Title)
	}
}

func Test_GetFilm_404_if_nonexistent_filmId(t *testing.T) {
	t.Parallel()
	err := swapi.GetFilmErr(0, http.StatusNotFound)
	assert.Equal(t, err.Detail, "Not found")
}
