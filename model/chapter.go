package model

import "gorm.io/gorm"

type Chapter struct {
	gorm.Model
	Name        string `json:"name"`
	Description string `json:"description"`

	CourseId uint `json:"courseId"`

	Course   *Course   `json:"course" gorm:"foreignKey:CourseId; constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Lessions []Lession `json:"lessions" gorm:"foreignKey:ChapterId"`
}
