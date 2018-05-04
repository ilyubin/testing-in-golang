package model

type RootResponse struct {
	People    string `json:"people"`
	Planets   string `json:"planets"`
	Films     string `json:"films"`
	Species   string `json:"species"`
	Vehicles  string `json:"vehicles"`
	StarShips string `json:"starships"`
}
