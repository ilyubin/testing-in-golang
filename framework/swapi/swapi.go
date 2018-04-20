package swapi

import (
	"fmt"
	"testing-in-golang/framework/http_client"
	. "testing-in-golang/framework/model"
)

func GetUrls() (out RootResponse) {
	http_client.GetOk("/", &out)
	return
}

func GerPeople() (out PeopleResponse) {
	http_client.GetOk("/people", &out)
	return
}

func GetPerson(num int) (out PersonResponse) {
	http_client.GetOk(fmt.Sprintf("/people/%d", num), &out)
	return
}

func GetPersonErr(num int, statusCode int) {
	http_client.GetErr(fmt.Sprintf("/people/%d", num), statusCode)
}
