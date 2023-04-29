package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	ID       int32
	Name     string
	RoleId   int64
	Username string
	Role     *Role `json:Role`
}

type Role struct {
	Id          int32
	Name        string
	Description string
}
