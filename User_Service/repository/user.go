package repository

import (
	"context"
	"time"

	"github.com/mooha76/Kofee/User_Service/model"
	"github.com/mooha76/Kofee/User_Service/repository/interfaces"
	"gorm.io/gorm"
)

type userDatabase struct {
	db *gorm.DB
}

// New UserRepository using gorm postgres databse
func NewUserRepository(db *gorm.DB) interfaces.UserRepository {
	return &userDatabase{
		db: db,
	}
}

func (c *userDatabase) FindUserByEmail(ctx context.Context, email string) (user model.User, err error) {

	query := `SELECT * FROM users WHERE email = $1`
	err = c.db.Raw(query, email).Scan(&user).Error

	return user, err
}

func (c *userDatabase) FindUserByPhone(ctx context.Context, phone string) (user model.User, err error) {

	query := `SELECT * FROM users WHERE phone = $1`
	err = c.db.Raw(query, phone).Scan(&user).Error

	return user, err
}

// Save a new user
func (c *userDatabase) SaveUser(ctx context.Context, user model.User) (uint64, error) {
	query := `INSERT INTO users ( first_name, middle_name , last_name, age, gender, email, phone, account , password,created_at) 
	VALUES ($1, $2, $3, $4, $5, $6, $7 ,$8 ,$9 ,$10 ) RETURNING id`

	createdAt := time.Now()
	err := c.db.Raw(query, user.FirstName, user.MiddleName, user.LastName,
		user.Age, user.Gender, user.Email, user.Phone, user.Account, user.Password, createdAt).Scan(&user).Error

	return user.ID, err
}
