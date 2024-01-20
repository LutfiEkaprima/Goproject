package models

import (
	"database/sql"
	"github.com/LutfiEkaprima/Goproject/config"
	"github.com/LutfiEkaprima/Goproject/entities"
)
type UserModel struct {
	db *sql.DB
}

func NewUserModel() *UserModel {
	conn, err := config.DBConn()

	if err != nil {
		panic(err)
	}

	return &UserModel{
		db: conn,
	}
}

func (u UserModel) Where(user *entities.User, fieldName, fieldValue string) error {

	row, err := u.db.Query("SELECT id, nama_lengkap, email, username, password FROM users WHERE "+fieldName+" = ? limit 1", fieldValue)

	if err != nil {
		return err
	}

	defer row.Close()

	for row.Next() {
		row.Scan(&user.Id, &user.NamaLengkap, &user.Email, &user.Username, &user.Password)
	}

	return nil

}