package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	UserID   int64  `db:"user_id"`
	Username string `db:"username"`
	Password string `db:"password"`
}
