package model

type PeopleResponse struct {
	Count    int                  `json:"count"`
	Next     string           `json:"next"`
	Previous interface{}      `json:"previous"`
	Results  []PersonResponse `json:"results"`
}
