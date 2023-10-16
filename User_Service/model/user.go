package model

import "time"

type Users struct {
	ID          uint64    `json:"id" db:"id"`
	FirstName   string    `json:"first_name,omitempty" db:"first_name" binding:"required,min=2,max=50"`
	MiddleName  string    `json:"middle_name" db:"middle_name" binding:"required,min=2,max=50"`
	LastName    string    `json:"last_name" db:"last_name" binding:"required,min=1,max=50"`
	Age         uint64    `json:"age" db:"age" binding:"required,numeric"`
	Gender      string    `json:"gender" db:"gender" binding:"required,min=4,max=20"`
	Email       string    `json:"email" db:"email" binding:"required,email"`
	Phone       string    `json:"phone" db:"phone" binding:"required,min=8,max=10"`
	Account     string    `json:"account" db:"account" binding:"required,min=8,max=10"`
	Password    string    `json:"password" db:"password" binding:"required"`
	Verified    bool      `json:"verified" db:"verified"`
	BlockStatus bool      `json:"block_status" db:"block_status"`
	CreatedAt   time.Time `json:"created_at" db:"created_at"`
	UpdatedAt   time.Time `json:"updated_at" db:"updated_at"`
}
