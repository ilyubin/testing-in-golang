package model

type FilmsResponse struct {
	Count    int            `json:"count"`
	Next     string         `json:"next"`
	Previous interface{}    `json:"previous"`
	Results  []FilmResponse `json:"results"`
}
