package main

import (
    "gorm.io/driver/mysql"
    "gorm.io/gorm"
    "fmt"
)

// User 定義
type User struct {
    ID    uint   `gorm:"primaryKey"`
    Name  string
    Email string
}

func main() {
    // MySQL接続のためのDSN（Data Source Name）設定
    dsn := "user:password@tcp(127.0.0.1:3306)/testdb?charset=utf8mb4&parseTime=True&loc=Local"
    db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
    if err != nil {
        fmt.Println("Error connecting to database:", err)
        return
    }

    // マイグレーション（テーブル作成）
    db.AutoMigrate(&User{})

    // 新しいユーザーの追加
    db.Create(&User{Name: "Alice", Email: "alice@example.com"})

    // ユーザーの取得
    var user User
    db.First(&user, 1) // IDが1のユーザーを取得
    fmt.Println("User:", user.Name, user.Email)
}

func createUser(db *gorm.DB, name, email string) {
    db.Create(&User{Name: name, Email: email})
}



