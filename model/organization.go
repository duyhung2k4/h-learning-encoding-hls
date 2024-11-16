package model

import "gorm.io/gorm"

type Organization struct {
	gorm.Model
	Name string `json:"name"`

	Profiles []Profile `json:"profiles" gorm:"foreignKey:OrganizationId;"`
}
