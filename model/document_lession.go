package model

import "gorm.io/gorm"

type DocumentLession struct {
	gorm.Model
	Code    string `json:"code"`
	Content string `json:"content"`

	LessionId uint `json:"lessionId"`

	Lession *Lession `json:"lession" gorm:"foreignKey:LessionId; constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}
