package entity

type Response struct {
	Message string
}

type ReponseWithData[T any] struct {
	Data    T `json:"data"`
	Message string
}
