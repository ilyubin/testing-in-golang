package tests

import (
	"github.com/DATA-DOG/godog"
	"testing-in-golang/framework/step"
	"testing-in-golang/framework/swapi/model"
)

func GetPerson_Ok(s *godog.Suite) {
	s.Step(`^send get person request with personId (\d+)$`, step.SendGetPersonRequest)
	s.Step(`^should received person "([^"]*)"$`, step.ShouldReceivedPerson)
	s.Step(`^it hair is "([^"]*)"$`, step.ItHairIs)
	s.Step(`^it eye is "([^"]*)"$`, step.ItEyeIs)

	s.BeforeScenario(func(interface{}) {
		step.Person = model.PersonResponse{}
	})
}

func GetPerson_Error(s *godog.Suite) {
	s.Step(`^send get person request with personId (\d+) expected error (\d+)$`, step.SendGetPersonErrorRequest)
	s.Step(`^should be error "([^"]*)"$`, step.ShouldBeError)

	s.BeforeScenario(func(interface{}) {
		step.Error = model.ErrorResponse{}
	})
}
