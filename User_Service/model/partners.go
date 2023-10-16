package model

import "time"

type PARTNERS struct {
	PARTNERID               uint64    `json:"PARTNERID" db:"PARTNERID"`
	FIRSTNAME               string    `json:"FIRSTNAME,omitempty" db:"FIRSTNAME" binding:"required,min=2,max=50"`
	MIDDLENAME              string    `json:"MIDDLENAME" db:"MIDDLENAME" binding:"required,min=2,max=50"`
	LASTNAME                string    `json:"LASTNAME" db:"LASTNAME" binding:"required,min=1,max=50"`
	IDENTITYTYPEID          uint64    `json:"IDENTITYTYPEID" db:"IDENTITYTYPEID"`
	MOBILE                  string    `json:"MOBILE" db:"MOBILE" binding:"required,min=6,max=10"`
	ACCOUNTNO               string    `json:"ACCOUNTNO" db:"ACCOUNTNO" binding:"required,min=6,max=10"`
	COMMISSIONPLANID        uint64    `json:"COMMISSIONPLANID" db:"COMMISSIONPLANID"`
	WITHDRAWALREQUESTTYPEID uint64    `json:"WITHDRAWALREQUESTTYPEID" db:"WITHDRAWALREQUESTTYPEID"`
	AGE                     uint64    `json:"AGE" db:"AGE" binding:"required,numeric"`
	GENDER                  string    `json:"GENDER" db:"GENDER" binding:"required,min=4,max=10"`
	EMAIL                   string    `json:"EMAIL" db:"EMAIL" binding:"required,email"`
	USERNAME                string    `json:"USERNAME" db:"USERNAME" binding:"required,email"`
	PASSWORD                string    `json:"PASSWORD" db:"PASSWORD" binding:"required,min=6,max=12"`
	ISACTIVE                bool      `json:"ISACTIVE" db:"ISACTIVE"`
	ISMOBILEACTIVE          bool      `json:"ISMOBILEACTIVE" db:"ISMOBILEACTIVE"`
	ISBLOCKED               bool      `json:"ISBLOCKED" db:"ISBLOCKED"`
	ISBLACKLIST             bool      `json:"ISBLACKLIST" db:"ISBLACKLIST"`
	CreatedAt               time.Time `json:"created_at" db:"created_at"`
	UpdatedAt               time.Time `json:"updated_at" db:"updated_at"`
}

type PARTNERSSINGUP_REQUEST struct {
	FIRSTNAME               string `json:"FIRSTNAME,omitempty" db:"FIRSTNAME" binding:"required,min=2,max=50"`
	MIDDLENAME              string `json:"MIDDLENAME" db:"MIDDLENAME" binding:"required,min=2,max=50"`
	LASTNAME                string `json:"LASTNAME" db:"LASTNAME" binding:"required,min=1,max=50"`
	MOBILE                  string `json:"MOBILE" db:"MOBILE" binding:"required,min=6,max=10"`
	ACCOUNTNO               string `json:"ACCOUNTNO" db:"ACCOUNTNO" binding:"required,min=6,max=10"`
	WITHDRAWALREQUESTTYPEID uint64 `json:"WITHDRAWALREQUESTTYPEID" db:"WITHDRAWALREQUESTTYPEID"`
	AGE                     uint64 `json:"AGE" db:"AGE" binding:"required,numeric"`
	GENDER                  string `json:"GENDER" db:"GENDER" binding:"required,min=4,max=10"`
	EMAIL                   string `json:"EMAIL" db:"EMAIL" binding:"required,email"`
	USERNAME                string `json:"USERNAME" db:"USERNAME" binding:"required,email"`
	PASSWORD                string `json:"PASSWORD" db:"PASSWORD" binding:"required,min=6,max=12"`
}
