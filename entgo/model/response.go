package model

type Response struct {
	Code   int    `json:"code""'`
	Status bool   `json:"status"`
	Error  string `json:"error"`
	Data   any    `json:"data"`
}
