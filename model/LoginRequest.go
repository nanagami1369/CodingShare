package model

type LoginRequest struct {
	Id    string `json:"id"`
	RowPassword string `json:"password"`
}
