package models

type Configuration struct {
	Version int    `json:"_version"`
	Data    string `json:"data"`
}
