package statesstruct

type StateResponse struct {
	Id     uint   `json:"id"`
	Name   string `json:"name"`
	Active bool   `json:"active"`
}
