package model

import "gorm.io/gorm"

type Role struct {
	gorm.Model
	Name string `json:"name"`
	Code ROLE   `json:"code" gorm:"unique"`

	Profiles []Profile `json:"profiles" gorm:"foreignKey:RoleId;"`
}

type ROLE string

const (
	ADMIN ROLE = "admin"
	USER  ROLE = "user"
)
