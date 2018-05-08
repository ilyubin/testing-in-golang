package spec

import (
	"github.com/DATA-DOG/godog"
	"testing-in-golang/framework/swapi/model"
	"testing-in-golang/step"
)

func GeTPerson_Ok(s *godog.Suite) {
	s.Step(`^send get person request with personId (\d+)$`, step.SendGetPersonRequest)
	s.Step(`^should received person "([^"]*)"$`, step.ShouldReceivedPerson)
	s.Step(`^it hair is "([^"]*)"$`, step.ItHairIs)
	s.Step(`^it eye is "([^"]*)"$`, step.ItEyeIs)

	s.BeforeScenario(func(interface{}) {
		step.Person = model.PersonResponse{}
	})
}
