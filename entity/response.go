package entity

type Response struct {
	Message string `json:"message"`
}

type ReponseWithData[T any] struct {
	Message string `json:"message"`
	Data    T      `json:"data"`
}
