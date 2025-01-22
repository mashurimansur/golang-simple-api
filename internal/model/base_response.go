package model

type BaseResponse[T any] struct {
	Code int `json:"code"`
	Data T   `json:"data"`
}
