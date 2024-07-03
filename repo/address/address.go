package address

import (
	"app/models"
	"context"
	"database/sql"
	"errors"
	"strings"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type AddressRepo struct {
	// you cold inject interface of methods to have more options when comes to db. but postgres now
	DB *sql.DB
}

// Creates a address for a user
func (repo *AddressRepo) Create(ctx context.Context, in models.Address) (err error) {
	r, err := repo.DB.ExecContext(ctx, `
	INSERT INTO 'address' (ID, USER_ID, ZIP_CODE, DETAILS, STATE, COUNTRY, CITY) VALUES($1,$2,$3,$4,$5,$6,$7)`,
		in.ID, in.UserID, in.ZipCode, in.Details, in.State, in.Country, in.City)
	if err != nil {
		return
	}
	qtd, err := r.RowsAffected()
	if err != nil {
		return
	}
	if qtd == 0 {
		return errors.New("no rows affected")
	}
	return
}

// Gets address from user, or get 1 address, in that case the slice will only have 1 obj
func (repo *AddressRepo) Get(ctx context.Context, userID, addressID string) (out []models.Address, err error) {
	query := `SELECT 
	ID,  ZIP_CODE, DETAILS, STATE, COUNTRY, CITY FROM 'address' WHERE USER_ID = ?`

	args := []interface{}{userID}

	if addressID != "" {
		query = strings.Join([]string{query, "AND ID = ?"}, "")
		args = append(args, addressID)
	}

	rows, err := repo.DB.QueryContext(ctx, sqlx.Rebind(sqlx.DOLLAR, query), args...)
	if err != nil {
		return
	}

	for rows.Next() {
		obj := models.Address{}
		err = rows.Scan(&obj.ID, &obj.ZipCode, &obj.Details, &obj.State, &obj.Country, &obj.City)
		if err != nil {
			return nil, err
		}
		out = append(out, obj)
	}

	return
}

// Deletes a user address based on userID and address ID
func (repo *AddressRepo) Delete(ctx context.Context, userID, addressID string) (err error) {
	r, err := repo.DB.ExecContext(ctx, `DELETE FROM 'address' WHERE USER_ID = $1 AND ID = $2`, userID, addressID)
	if err != nil {
		return
	}
	qtd, err := r.RowsAffected()
	if err != nil {
		return
	}
	if qtd == 0 {
		return errors.New("no rows affected")
	}
	return
}

// Patch patchs a address info
func (repo *AddressRepo) Patch(ctx context.Context, addressID string, in models.Address) (err error) {
	query := "UPDATE 'users' SET "
	args := []interface{}{}
	if in.Details != "" {
		query = strings.Join([]string{query, "address = ?"}, "")
		args = append(args, in.Details)
	}

	query = strings.Join([]string{query, "WHERE ID = ? AND USER_ID = ?"}, "")
	args = append(args, addressID, in.UserID)

	r, err := repo.DB.ExecContext(ctx, sqlx.Rebind(sqlx.DOLLAR, query), args...)
	if err != nil {
		return
	}
	qtd, err := r.RowsAffected()
	if err != nil {
		return
	}
	if qtd == 0 {
		return errors.New("no rows affected")
	}
	return
}
