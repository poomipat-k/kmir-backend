package user

import (
	"database/sql"
	"log/slog"
	"strings"
)

type store struct {
	db *sql.DB
}

func NewStore(db *sql.DB) *store {
	return &store{
		db: db,
	}
}

func (s *store) GetUserByUsername(username string) (User, error) {
	var user User
	row := s.db.QueryRow(getUserByUsernameSQL, strings.ToLower(username))
	err := row.Scan(&user.Id, &user.Username, &user.Password, &user.DisplayName, &user.UserRole)
	switch err {
	case sql.ErrNoRows:
		slog.Error("GetUserByUsername() no row were returned!")
		return User{}, err
	case nil:
		return user, nil
	default:
		slog.Error(err.Error())
		return User{}, err
	}
}
