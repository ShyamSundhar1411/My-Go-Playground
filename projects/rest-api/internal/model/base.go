package model

import (
	"time"

)

type Base struct{
	ID uint `gorm:"primary_key" json:"id"`
	CreatedAt time.Time 
	UpdatedAt time.Time 
}

type Role struct{
	Base
	Name string	`json:"name" gorm:"type:varchar(100);not null;unique"`
	Code string `json:"code" gorm:"type:varchar(100);not null;unique"`
}
type BaseUser struct{
	Base
	UserName string `json:"username" gorm:"type:varchar(100);not null;unique"`
	Password  string `json:"password" gorm:"not null"`
	FirstName string `json:"firstname" gorm:"type:varchar(100);not null"`
	LastName  string `json:"lastname" gorm:"type:varchar(100);not null"`
	Roles []Role `json:"roles" gorm:"many2many:user_roles;constraint:OnDelete:CASCADE;"`
}
