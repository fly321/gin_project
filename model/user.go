package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name string
	Age  int64
}

func (User) TableName() string {
	return "users"
}
