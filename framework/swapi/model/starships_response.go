package model

//StarshipsResponse
type StarshipsResponse struct {
	Count    int                `json:"count"`
	Next     string             `json:"next"`
	Previous interface{}        `json:"previous"`
	Results  []StarshipResponse `json:"results"`
}
