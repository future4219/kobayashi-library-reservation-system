package main

import (
	"api/domain"
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"gorm.io/driver/mysql"
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

func main() {
	// MySQL接続のためのDSN（Data Source Name）設定
	dsn := "user:password@tcp(127.0.0.1:3306)/testdb?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println("Error connecting to database:", err)
		return
	}

	userdatabase := NewUserDatabase(db)

	// マイグレーション（テーブル作成）
	db.AutoMigrate(&domain.User{})

	// 新しいユーザーの追加
	db.Create(&domain.User{Name: "Alice", Email: "alice@example.com"})

	// ユーザーの取得
	var user domain.User
	db.First(&user, 1) // IDが1のユーザーを取得
	fmt.Println("User:", user.Name, user.Email)

	e := echo.New()

	// Route to get all users
	e.POST("/users", userdatabase.Create)

	// Route to add a new user
	// e.POST("/users", createUser)
}

func (ud UserDatabase) Create(c echo.Context) error {
	// ユーザーを作成する処理
	var user domain.User
	if err := c.Bind(&user); err != nil {
		return err
	}

	err := ud.db.Create(&domain.User{Name: user.Name, Email: user.Email}).Error
	if err != nil {
		return fmt.Errorf("dfdfda")
	}

	return c.JSON(http.StatusCreated, user)
}
