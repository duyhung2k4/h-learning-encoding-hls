package model

import (
	"gorm.io/gorm"
)

type Profile struct {
	gorm.Model
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Phone     string `json:"phone"`

	Email    string `json:"email" gorm:"unique"`
	Username string `json:"username"`
	Password string `json:"password"`
	Active   bool   `json:"active"`

	RoleId         uint  `json:"roleId"`
	OrganizationId *uint `json:"organizationId"`

	Role            *Role            `json:"role" gorm:"foreignKey:RoleId; constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Organization    *Organization    `json:"organization" gorm:"foreignKey:OrganizationId; constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	CourseRegisters []CourseRegister `json:"courseRegisters" gorm:"foreignKey:ProfileId;"`
}
