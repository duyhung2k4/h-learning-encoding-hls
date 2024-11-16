package model

import "gorm.io/gorm"

type TextNote struct {
	gorm.Model
	Time    string `json:"time"`
	Context string `json:"context"`

	VideoLessionId uint          `json:"videoLessionId"`
	VideoLession   *VideoLession `json:"videoLession" gorm:"foreignKey:VideoLessionId; constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}
