package repository

import (
	"context"
	"database/sql"

	"github.com/jmoiron/sqlx"
	"github.com/mooha76/Kofee/User_Service/model"
	"github.com/mooha76/Kofee/User_Service/repository/interfaces"
)

type userDatabase struct {
	db *sqlx.DB
}

// New UserRepository using gorm postgres databse
func NewUserRepository(db *sqlx.DB) interfaces.UserRepository {
	return &userDatabase{
		db: db,
	}
}

func (c *userDatabase) FindUserByEmail(ctx context.Context, email string) (user model.User, err error) {
	query := `SELECT * FROM users WHERE email = $1`

	err = c.db.GetContext(ctx, &user, query, email)

	if err != nil {
		if err == sql.ErrNoRows {
			return user, nil // User not found, return an empty user and no error
		}
		return user, err // Return any other error
	}

	return user, nil
}

func (c *userDatabase) FindUserByPhone(ctx context.Context, phone string) (user model.User, err error) {

	quary := `SELECT * FROM users WHERE phone = $1`
	err = c.db.GetContext(ctx, &user, quary, phone)
	if err != nil {
		if err == sql.ErrNoRows {
			return user, nil // User not found, return an empty user and no error
		}
		return user, err // Return any other error
	}

	return user, nil
}

// Save a new user
func (c *userDatabase) SaveUser(ctx context.Context, user model.User) (uint64, error) {
	query := `
        INSERT INTO users (first_name, middle_name, last_name, age, gender, email, phone, account, password, created_at) 
        VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10) 
        RETURNING id
    `

	var userID uint64

	err := c.db.QueryRowContext(ctx, query,
		&user,
	).Scan(&userID)

	if err != nil {
		return 0, err
	}

	return userID, nil
}
