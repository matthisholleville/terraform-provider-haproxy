package models

type Transaction struct {
	Version int    `json:"_version"`
	Id      string `json:"id"`
	Status  string `json:"status"`
}
