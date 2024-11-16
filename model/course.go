package model

import "gorm.io/gorm"

type Course struct {
	gorm.Model
	Name        string  `json:"name"`
	Code        string  `json:"code" gorm:"unique"`
	Description string  `json:"description"`
	MultiLogin  bool    `json:"multiLogin"`
	Value       float64 `json:"value"`

	CategoryId uint `json:"categoryId"`
	CreateId   uint `json:"createId"`

	Category    *Category    `json:"category" gorm:"foreignKey:CategoryId; constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Create      *Profile     `json:"create" gorm:"foreignKey:CreateId; constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Chapters    []Chapter    `json:"chapters" gorm:"foreignKey:CourseId;"`
	SaleCourses []SaleCourse `json:"saleCourses" gorm:"foreignKey:CourseId;"`
}
