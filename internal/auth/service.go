package auth

import (
	"database/sql"
	"errors"
)

type AuthService struct {
	db *sql.DB
}

func NewAuthService(db *sql.DB) *AuthService {
    return &AuthService{db: db}
}

func (s *AuthService) RegisterUser(username, password string) (int, error) {
	var userID int
	err := s.db.QueryRow("INSERT INTO users (username, password) VALUES ($1, $2) RETURNING id", username, password).Scan(&userID)
	if err != nil {
		return 0, err
	}
	return userID, nil
}

func (s *AuthService) LoginUser(username, password string) (int, bool, error) {
	var storedPassword string
	var userID int

	err := s.db.QueryRow("SELECT id, password FROM users WHERE username = $1", username).Scan(&userID, &storedPassword)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return 0, false, nil
		}
		return 0, false, err
	}

	if storedPassword == password {
		return userID, true, nil
	}

	return 0, false, nil
}