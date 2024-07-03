package user

import (
	"app/models"
	"context"
	"database/sql"
	"errors"
	"strings"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type UserRepo struct {
	// you cold inject interface of methods to have more options when comes to db. but postgres now
	DB *sql.DB
}

func (repo *UserRepo) Create(ctx context.Context, in models.User) (err error) {
	r, err := repo.DB.ExecContext(ctx, `INSERT INTO 'users' (ID,NAME,LAST_NAME,AGE) VALUES($1,$2,$3,$4)`, in.ID, in.Name, in.LastName, in.Age)
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
func (repo *UserRepo) ExistID(ctx context.Context, id string) (ok bool, err error) {
	err = repo.DB.QueryRowContext(ctx, `SELECT TRUE FROM 'users' WHERE ID= $1`, id).Scan(&ok)
	return
}
func (repo *UserRepo) Get(ctx context.Context, id string) (out models.User, err error) {
	err = repo.DB.QueryRowContext(ctx, `SELECT NAME, LAST_NAME, AGE FROM 'users' WHERE ID= $1`, id).Scan(
		&out.Name, &out.LastName, &out.Age,
	)
	return
}
func (repo *UserRepo) Delete(ctx context.Context, id string) (err error) {
	r, err := repo.DB.ExecContext(ctx, `DELETE FROM 'users' WHERE ID = $1`, id)
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

func (repo *UserRepo) Patch(ctx context.Context, in models.User) (err error) {
	query := "UPDATE 'users' SET "
	args := []interface{}{}
	if in.Age != 0 {
		query = strings.Join([]string{query, "AGE = ?"}, "")
		args = append(args, in.Age)
	}

	if in.LastName != nil {
		query = strings.Join([]string{query, "LAST_NAME = ?"}, "")
		args = append(args, in.LastName)
	}

	if in.Name != "" {
		query = strings.Join([]string{query, "NAME = ?"}, "")
		args = append(args, in.LastName)
	}

	query = strings.Join([]string{query, "WHERE ID = ?"}, "")
	args = append(args, in.ID)

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
