package model

type PlanetsResponse struct {
	Count    int              `json:"count"`
	Next     string           `json:"next"`
	Previous interface{}      `json:"previous"`
	Results  []PlanetResponse `json:"results"`
}
