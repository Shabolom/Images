package models

type Error struct {
	Code  int    `json:"code"`
	Error string `json:"error"`
}
