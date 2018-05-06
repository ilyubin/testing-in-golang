package model

import "time"

type StarshipResponse struct {
	Name                 string        `json:"name"`
	Model                string        `json:"model"`
	Manufacturer         string        `json:"manufacturer"`
	CostInCredits        string        `json:"cost_in_credits"`
	Length               string        `json:"length"`
	MaxAtmospheringSpeed string        `json:"max_atmosphering_speed"`
	Crew                 string        `json:"crew"`
	Passengers           string        `json:"passengers"`
	CargoCapacity        string        `json:"cargo_capacity"`
	Consumables          string        `json:"consumables"`
	HyperdriveRating     string        `json:"hyperdrive_rating"`
	MGLT                 string        `json:"MGLT"`
	StarshipClass        string        `json:"starship_class"`
	Pilots               []interface{} `json:"pilots"`
	Films                []string      `json:"films"`
	Created              time.Time     `json:"created"`
	Edited               time.Time     `json:"edited"`
	URL                  string        `json:"url"`
}
