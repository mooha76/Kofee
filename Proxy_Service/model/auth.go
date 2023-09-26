package model

type UserSignupRequest struct {
	FirstName  string `json:"first_name" binding:"required"`
	MiddleName string `json:"middle_name" binding:"required"`
	LastName   string `json:"last_name" binding:"required"`
	Age        uint64 `json:"age" binding:"required"`
	Gender     string `json:"gender" binding:"required,min=4,max=20"`
	Email      string `json:"email" binding:"required,email"`
	Phone      string `json:"phone" binding:"required,min=8,max=20"`
	Account    string `json:"account" binding:"required,min=8,max=20"`
	Password   string `json:"password" binding:"required,min=6,max=30"`
}

type UserLoginRequest struct {
	Email    string `json:"email" binding:"omitempty,email"`
	Phone    string `json:"phone" binding:"omitempty,min=10,max=10"`
	Password string `json:"password" binding:"required"`
}
