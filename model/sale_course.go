package model

import (
	"time"

	"gorm.io/gorm"
)

type SaleCourse struct {
	gorm.Model
	Start    time.Time `json:"start"`
	Finish   time.Time `json:"finish"`
	NewValue float64   `json:"newValue"`

	CourseId        uint             `json:"courseId"`
	Course          *Course          `json:"course" gorm:"foreignKey:CourseId; constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	CourseRegisters []CourseRegister `json:"courseRegisters" gorm:"foreignKey:SaleCourseId;"`
}
