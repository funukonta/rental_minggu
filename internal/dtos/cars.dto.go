package dtos

type CreateCarReq struct {
	Model      string  `json:"model"`
	PlatNumber string  `json:"plat_number"`
	DailyRate  float64 `json:"daily_rate"`
	Status     bool    `json:"status"`
	Color      string  `json:"color"`
	Make       string  `json:"make"`
	Year       int     `json:"year"`
}

func (c *CreateCarReq) Normalize() {
	if !c.Status {
		c.Status = true
	}
}
