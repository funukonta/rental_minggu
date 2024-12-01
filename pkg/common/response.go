package common

type RespCreate struct {
	Success     bool   `json:"success"`
	Message     string `json:"message"`
	ReferenceID string `json:"reference_id"`
	Data        any    `json:"data"`
}
