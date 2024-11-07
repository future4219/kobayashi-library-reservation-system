// ドメイン知識を定義する

package domain

// User 定義
type User struct {
	ID    uint `gorm:"primaryKey"`
	Name  string
	Email string
}
