package model

import "gorm.io/gorm"

type Budget struct {
	gorm.Model
	Value float64 `json:"value"`

	ProfileId uint `json:"profileId"`

	Profile *Profile `json:"profile" gorm:"foreignKey:ProfileId; constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}
