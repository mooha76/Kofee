package repository

import (
	"context"
	"database/sql"

	"github.com/jmoiron/sqlx"
	"github.com/mooha76/Kofee/User_Service/model"
	"github.com/mooha76/Kofee/User_Service/repository/interfaces"
)

type partnerDatabase struct {
	db *sqlx.DB
}

func NewPartnerRepository(db *sqlx.DB) interfaces.PartnerRepository {
	return &partnerDatabase{
		db: db,
	}

}

func (c *partnerDatabase) FindPartnerByEmail(ctx context.Context, email string) (Partner model.PARTNERS, err error) {
	query := `SELECT * FROM PARTNERS WHERE email = $1`

	err = c.db.GetContext(ctx, &Partner, query, email)

	if err != nil {
		if err == sql.ErrNoRows {
			return Partner, nil // User not found, return an empty user and no error
		}
		return Partner, err // Return any other error
	}

	return Partner, nil
}

func (c *partnerDatabase) FindPartnerByPhone(ctx context.Context, phone string) (Partner model.PARTNERS, err error) {

	quary := `SELECT * FROM PARTNERS WHERE phone = $1`
	err = c.db.GetContext(ctx, &Partner, quary, phone)
	if err != nil {
		if err == sql.ErrNoRows {
			return Partner, nil // User not found, return an empty user and no error
		}
		return Partner, err // Return any other error
	}

	return Partner, nil
}

// Save a new user
func (c *partnerDatabase) SaveParner(ctx context.Context, partner model.PARTNERSSINGUP_REQUEST) (uint64, error) {
	query := `
		INSERT INTO public."PARTNERS"(
			"PARTNERID", "FIRSTNAME", "MIDDLENAME", "LASTNAME", "USERNAME", "PASSWORD", 
			"WITHDRAWALREQUESTTYPEID", "IDENTITYTYPEID", "MOBILE", "ACCOUNTNO", 
			"COMMISSIONPLANID", "AGE", "GENDER", "EMAIL")
		VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)
		RETURNING "PARTNERID"
	`

	var partnerID uint64

	// Execute the SQL query
	row := c.db.QueryRowContext(ctx, query,
		partner.FIRSTNAME, partner.MIDDLENAME, partner.LASTNAME,
		partner.USERNAME, partner.PASSWORD, partner.WITHDRAWALREQUESTTYPEID,
		//partner.IDENTITYTYPEID, partner.MOBILE, partner.ACCOUNTNO,
		//partner.COMMISSIONPLANID, partner.AGE, partner.GENDER,
		partner.EMAIL,
	)

	// Scan the result into partnerID
	if err := row.Scan(&partnerID); err != nil {
		// Handle the error
		return 0, err
	}

	return partnerID, nil
}
