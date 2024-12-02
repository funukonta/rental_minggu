package common

type RespCreate struct {
	Success     bool   `json:"success"`
	Message     string `json:"message"`
	ReferenceID string `json:"reference_id"`
	Data        any    `json:"data"`
}

type RespList struct {
	TotalItem int         `json:"total_item"`
	Success   bool        `json:"success"`
	Message   string      `json:"message"`
	Data      interface{} `json:"data"`
}

type RespDetail struct {
	Success bool        `json:"success"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}
