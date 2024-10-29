package db

import (
	"database/sql"
	"golauth/model"
)

type DB interface {
	// Auth
	MatchEmailAndPassword(email string, password string) (bool, *model.User)
	AssignProfile(user_uuid string, profile_uuid string) (bool, error)
	GetUserByID(id int64) (*model.Auth, error)
	Register(user *model.Auth) (bool, error, int64)
}

type MySQLDB struct {
	db *sql.DB
}

func NewDB(db *sql.DB) DB {
	return MySQLDB{db: db}
}
