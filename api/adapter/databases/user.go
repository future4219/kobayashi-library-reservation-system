// データベースをgormでいじるための関数を定義する

package databases

import (
	"api/api/domain"

	"gorm.io/gorm"
)

type UserDatabase struct {
	db *gorm.DB
}

func NewUserDatabase(
	db *gorm.DB,
) UserDatabase {
	return UserDatabase{db: db}
}

func (ud UserDatabase) Create(name string, mail string) error {
	res := ud.db.Create(&domain.User{Name: name, Email: mail}).Error
	return res
}
