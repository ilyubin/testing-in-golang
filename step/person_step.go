package step

import (
	"fmt"
	"testing-in-golang/framework/swapi"
	"testing-in-golang/framework/swapi/model"
)

var Person model.PersonResponse

func SendGetPersonRequest(personId int) error {
	Person = swapi.GetPerson(personId)
	return nil
}

func ShouldReceivedPerson(name string) error {
	if Person.Name != name {
		return fmt.Errorf("expected Person Name %q, but was %q", Person.Name, name)
	}
	return nil
}

func ItHairIs(hair string) error {
	if Person.HairColor != hair {
		return fmt.Errorf("expected Person HairColor %q, but was %q", Person.HairColor, hair)
	}
	return nil
}

func ItEyeIs(eye string) error {
	if Person.EyeColor != eye {
		return fmt.Errorf("expected Person EyeColor %q, but was %q", Person.EyeColor, eye)
	}
	return nil
}
