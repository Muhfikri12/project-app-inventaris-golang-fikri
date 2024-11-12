package apirepository

import (
	"database/sql"
	"errors"
	"time"

	"github.com/Muhfikri12/project-app-inventaris-golang-fikri/model"
	"github.com/google/uuid"
)

type AuthRepo struct {
	db *sql.DB
}

func NewAuthRepo(db *sql.DB) *AuthRepo {
	return &AuthRepo{db}
}

func (u *AuthRepo) LoginUser(user *model.User) error {

	query := `SELECT id token FROM users WHERE email=$1 AND password=$2`

	err := u.db.QueryRow(query, user.Email, user.Password).Scan(&user.ID)
	if err != nil {
		if err == sql.ErrNoRows {
			return errors.New("invalid email or password")
		}
	}

	token := uuid.New().String()

	expiration := time.Now().Add(1 * time.Hour) 

	updateQuery := `UPDATE users SET token=$1, expired=$2 WHERE id=$3`

	_, err = u.db.Exec(updateQuery, token, expiration, user.ID)
	if err != nil {
		return err
	}

	user.Token = token

	return nil
}

func (u *AuthRepo) CheckToken(token string) (bool, error) {
	var expired time.Time

	// Query untuk mengecek token dan expired
	query := `SELECT expired FROM users WHERE token=$1`
	err := u.db.QueryRow(query, token).Scan(&expired)
	if err != nil {
		if err == sql.ErrNoRows {
			return false, errors.New("token tidak ditemukan")
		}
		return false, err
	}

	// Cek apakah token sudah kadaluarsa
	if time.Now().After(expired) {
		return false, errors.New("token sudah kadaluarsa")
	}

	return true, nil
}
