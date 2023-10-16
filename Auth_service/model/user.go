package model

type UserSignupRequest struct {
	FirstName  string `json:"first_name,omitempty" db:"first_name" binding:"required,min=2,max=50"`
	MiddleName string `json:"middle_name" db:"middle_name" binding:"required,min=2,max=50"`
	LastName   string `json:"last_name" db:"last_name" binding:"required,min=1,max=50"`
	Age        uint64 `json:"age" db:"age" binding:"required,numeric"`
	Gender     string `json:"gender" db:"gender" binding:"required,min=4,max=20"`
	Email      string `json:"email" db:"email" binding:"required,email"`
	Phone      string `json:"phone" db:"phone" binding:"required,min=8,max=10"`
	Account    string `json:"account" db:"account" binding:"required,min=8,max=10"`
	Password   string `json:"password" db:"password" binding:"required"`
}

type UserLoginRequest struct {
	Email    string
	Phone    string
	Password string
}

type User struct {
	ID          uint64
	FirstName   string
	LastName    string
	Age         uint64
	Email       string
	Phone       string
	Password    string
	Verified    bool
	BlockStatus bool
}
