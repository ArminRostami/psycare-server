package postgres

import (
	"fmt"
	app "psycare/internal/domain"

	"github.com/jmoiron/sqlx"
)

// UserRepo implements the app.UserRepo interface using postgresql
type UserRepo struct {
	db *sqlx.DB
}

func (ur *UserRepo) connect(connStr string) error {
	db, err := sqlx.Connect(pgDriver, connStr)
	if err != nil {
		return fmt.Errorf("db connection error: %w", err)
	}
	ur.db = db
	return nil
}

func (ur *UserRepo) getUserWithName(username string) (*app.User, error) {
	u := &app.User{}
	err := ur.db.Get(u, "SELECT * FROM users WHERE (user_name=$1)", username)
	if err != nil {
		return nil, fmt.Errorf("no such user: %w", err)
	}
	return u, nil
}

func (ur *UserRepo) addUser(u app.User) error {
	tx := ur.db.MustBegin()
	tx.NamedExec("INSERT INTO users (user_name,email,password) VALUES (:user_name,:email,:password)", u)
	err := tx.Commit()
	if err != nil {
		return fmt.Errorf("inserting new user failed: %w", err)
	}
	return nil
}
