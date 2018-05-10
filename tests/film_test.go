package tests

import (
	"github.com/DATA-DOG/godog"
	"testing-in-golang/framework/step"
)

func ResetFilm(s *godog.Suite) {
	s.Step(`^reset film$`, step.ResetFilm)
}

func GetFilm_Ok(s *godog.Suite) {
	s.Step(`^send get film request with (\d+)$`, step.SendGetFilmRequest)
	s.Step(`^should receive film ([^"]*)$`, step.ShouldReceivedFilm)
}
