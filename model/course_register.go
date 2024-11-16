package model

import (
	"github.com/lib/pq"
	"gorm.io/gorm"
)

type CourseRegister struct {
	gorm.Model
	IpLogin pq.StringArray `json:"ipLogin" gorm:"type:text[]"`

	ProfileId    uint `json:"profileId"`
	CourseId     uint `json:"courseId"`
	SaleCourseId uint `json:"saleCourseId"`

	Profile    *Profile    `json:"profile" gorm:"foreignKey:ProfileId; constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Course     *Course     `json:"course" gorm:"foreignKey:CourseId; constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	SaleCourse *SaleCourse `json:"saleCourse" gorm:"foreignKey:SaleCourseId; constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}
