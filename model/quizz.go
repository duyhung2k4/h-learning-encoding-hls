package model

import (
	"github.com/lib/pq"
	"gorm.io/gorm"
)

type Quizz struct {
	gorm.Model
	Ask        string         `json:"ask"`
	ResultType string         `json:"resultType"`
	Result     pq.StringArray `json:"result" gorm:"type:text[]"`
	Pption     pq.StringArray `json:"option" gorm:"type:text[]"`
	Time       int            `json:"time"`

	EntityType string `json:"entityType"`
	EntityId   uint   `json:"entityId"`
}
