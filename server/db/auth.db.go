package db

import (
	"database/sql"
	"fmt"
	"golauth/model"
)

func (d MySQLDB) Register(user *model.Auth) (bool, error, int64) {
	var result sql.Result
	db, err := d.db.Prepare("INSERT INTO auth(email, password) VALUES (?, ?)")
	if err != nil {
		return false, err, 0
	}

	result, err = db.Exec(user.Email, user.Password)
	id, err := result.LastInsertId()
	if err != nil {
		return false, err, 0
	}
	defer db.Close()

	return true, err, id
}

func (d MySQLDB) AssignProfile(user_uuid string, profile_uuid string) (bool, error) {
	db, err := d.db.Prepare("UPDATE auth SET profile_uuid =  ? where uuid = ?")
	if err != nil {
		return false, err
	}

	_, err = db.Exec(profile_uuid, user_uuid)
	if err != nil {
		return false, err
	}
	defer db.Close()

	return true, err
}

func (d MySQLDB) MatchEmailAndPassword(email string, password string) (bool, *model.User) {
	query := fmt.Sprintf("SELECT email FROM auth WHERE email = '%s' AND password = SHA1('%s');", email, password)
	fmt.Println(query)
	rows, err := d.db.Query(query)
	if err != nil {
		return false, &model.User{}
	}
	defer rows.Close()

	t := new(model.User)
	for rows.Next() {
		switch err := rows.Scan(&t.Email); err {
		case sql.ErrNoRows:
			return false, t
		case nil:
			return true, t
		default:
			return false, t
		}
	}

	return false, &model.User{}
}

func (d MySQLDB) GetUserByID(id int64) (*model.Auth, error) {
	query := fmt.Sprintf(`SELECT uuid FROM auth WHERE id =  %d`, id)
	rows, err := d.db.Query(query)
	if err != nil {
		return &model.Auth{}, err
	}
	defer rows.Close()

	for rows.Next() {
		user := new(model.Auth)
		if err := rows.Scan(&user.Uuid); err == nil {
			return user, err
		}
	}
	return &model.Auth{}, err
}
