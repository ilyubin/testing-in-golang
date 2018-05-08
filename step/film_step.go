package step

import (
	"fmt"
	"testing-in-golang/framework/swapi"
	"testing-in-golang/framework/swapi/model"
)

var Film model.FilmResponse

func SendGetFilmRequest(filmId int) error {
	Film = swapi.GetFilm(filmId)
	return nil
}

func ShouldReceivedFilm(title string) error {
	if Film.Title != title {
		return fmt.Errorf("expected Film Title %q, but was %q", Film.Title, title)
	}
	return nil
}

func ResetFilm() error {
	Film = model.FilmResponse{}
	return nil
}
