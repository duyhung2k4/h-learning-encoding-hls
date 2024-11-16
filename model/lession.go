package model

import "gorm.io/gorm"

type Lession struct {
	gorm.Model
	Name        string `json:"name"`
	Description string `json:"description"`

	ChapterId uint `json:"chapterId"`
	CourseId  uint `json:"courseId"`

	Chapter      *Chapter      `json:"chapter" gorm:"foreignKey:ChapterId; constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Course       *Course       `json:"course" gorm:"foreignKey:CourseId; constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	VideoLession *VideoLession `json:"cideoLession" gorm:"foreignKey:LessionId"`
}
