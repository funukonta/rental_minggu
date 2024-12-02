package dtos

type CreateBookingReq struct {
	CarsID        string `json:"cars_id" gorm:"foreignKey:cars_id"`
	UserID        string `json:"user_id" gorm:"foreignKey:user_id"`
	StartDate     string `json:"start_date"`
	EndDate       string `json:"end_date"`
	PaymentMethod string `json:"payment_method"`
}

type UpdateBookingReq struct {
	CarsID        string `json:"cars_id"`
	UserID        string `json:"user_id"`
	StartDate     string `json:"start_date"`
	EndDate       string `json:"end_date"`
	Status        string `json:"status"`
	PaymentMethod string `json:"payment_method"`
}
