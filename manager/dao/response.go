package dao

//common
type (
	Response struct {
		Code    int    `json:"code"`
		Message string `json:"message"`
		Success bool   `json:"success"`
	}
)

type (
	GetHelloRes struct{
		Response
		Data string `json:"data"`
	}
)
