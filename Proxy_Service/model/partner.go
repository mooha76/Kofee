package model

import "time"

type Partner struct {
	FIRSTNAME               string    `json:"FIRSTNAME,omitempty" db:"FIRSTNAME" binding:"required,min=2,max=50"`
	MIDDLENAME              string    `json:"MIDDLENAME" db:"MIDDLENAME" binding:"required,min=2,max=50"`
	LASTNAME                string    `json:"LASTNAME" db:"LASTNAME" binding:"required,min=1,max=50"`
	MOBILE                  string    `json:"MOBILE" db:"MOBILE" binding:"required,min=6,max=10"`
	ACCOUNTNO               string    `json:"ACCOUNTNO" db:"ACCOUNTNO" binding:"required,min=6,max=10"`
	WITHDRAWALREQUESTTYPEID uint64    `json:"WITHDRAWALREQUESTTYPEID" db:"WITHDRAWALREQUESTTYPEID"`
	AGE                     uint64    `json:"AGE" db:"AGE" binding:"required,numeric"`
	GENDER                  string    `json:"GENDER" db:"GENDER" binding:"required,min=4,max=10"`
	EMAIL                   string    `json:"EMAIL" db:"EMAIL" binding:"required,email"`
	USERNAME                string    `json:"USERNAME" db:"USERNAME" binding:"required,min=4,max=10"`
	PASSWORD                string    `json:"PASSWORD" db:"PASSWORD" binding:"required,min=6,max=12"`
	CreatedAt               time.Time `json:"created_at" db:"created_at"`
	UpdatedAt               time.Time `json:"updated_at" db:"updated_at"`
}

type PartnerSingUpRequest struct {
	FIRSTNAME               string `json:"FIRSTNAME,omitempty" db:"FIRSTNAME" binding:"required,min=2,max=50"`
	MIDDLENAME              string `json:"MIDDLENAME" db:"MIDDLENAME" binding:"required,min=2,max=50"`
	LASTNAME                string `json:"LASTNAME" db:"LASTNAME" binding:"required,min=1,max=50"`
	MOBILE                  string `json:"MOBILE" db:"MOBILE" binding:"required,min=6,max=10"`
	ACCOUNTNO               string `json:"ACCOUNTNO" db:"ACCOUNTNO" binding:"required,min=6,max=10"`
	WITHDRAWALREQUESTTYPEID uint64 `json:"WITHDRAWALREQUESTTYPEID" db:"WITHDRAWALREQUESTTYPEID"`
	AGE                     uint64 `json:"AGE" db:"AGE" binding:"required,numeric"`
	GENDER                  string `json:"GENDER" db:"GENDER" binding:"required,min=4,max=10"`
	EMAIL                   string `json:"EMAIL" db:"EMAIL" binding:"required,email"`
	USERNAME                string `json:"USERNAME" db:"USERNAME" binding:"required,min=4,max=10"`
	PASSWORD                string `json:"PASSWORD" db:"PASSWORD" binding:"required,min=6,max=12"`
}
