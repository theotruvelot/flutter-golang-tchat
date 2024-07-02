package user

import (
	"context"
	"database/sql"
)

type DBTX interface {
	ExecContext(ctx context.Context, query string, args ...interface{}) (sql.Result, error)
	PrepareContext(context.Context, string) (*sql.Stmt, error)
	QueryContext(context.Context, string, ...interface{}) (*sql.Rows, error)
	QueryRowContext(context.Context, string, ...interface{}) *sql.Row
}

type repository struct {
	db DBTX
}

func NewRepository(db DBTX) Repository {
	return &repository{db: db}
}

func (r *repository) CreateUser(ctx context.Context, user *User) (*User, error) {
	var lastInsertId int
	rawQuery := `INSERT INTO "user"(username, password, email) VALUES ($1, $2, $3) RETURNING id`
	err := r.db.QueryRowContext(ctx, rawQuery, user.Username, user.Password, user.Email).Scan(&lastInsertId)
	if err != nil {
		return &User{}, err
	}

	user.ID = int64(lastInsertId)
	return user, nil
}

func (r *repository) CheckIfUserExists(ctx context.Context, email string) (bool, error) {
	var exists bool
	rawQuery := `SELECT EXISTS(SELECT 1 FROM "user" WHERE email = $1)`
	err := r.db.QueryRowContext(ctx, rawQuery, email).Scan(&exists)
	if err != nil {
		return false, err
	}

	return exists, nil
}

func (r *repository) GetUserByEmail(ctx context.Context, email string) (*User, error) {
	u := User{}
	rawQuery := `SELECT id, email, username, password FROM "user" WHERE email = $1`
	err := r.db.QueryRowContext(ctx, rawQuery, email).Scan(&u.ID, &u.Email, &u.Username, &u.Password)
	if err != nil {
		return &User{}, nil
	}

	return &u, nil
}
