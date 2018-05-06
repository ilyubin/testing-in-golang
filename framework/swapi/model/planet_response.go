package model

import "time"

type PlanetResponse struct {
	Name           string        `json:"name"`
	RotationPeriod string        `json:"rotation_period"`
	OrbitalPeriod  string        `json:"orbital_period"`
	Diameter       string        `json:"diameter"`
	Climate        string        `json:"climate"`
	Gravity        string        `json:"gravity"`
	Terrain        string        `json:"terrain"`
	SurfaceWater   string        `json:"surface_water"`
	Population     string        `json:"population"`
	Residents      []interface{} `json:"residents"`
	Films          []string      `json:"films"`
	Created        time.Time     `json:"created"`
	Edited         time.Time     `json:"edited"`
	URL            string        `json:"url"`
}
