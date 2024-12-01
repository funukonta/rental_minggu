package dtos

type RegisterReq struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
	Role     string `json:"role"`
}

type LoginReq struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
