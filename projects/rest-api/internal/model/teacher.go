package model


type Teacher struct{
	BaseUser
	DepartmentID uint `json:"department_id" gorm:"not null"`
	Department Department `json:"department" gorm:"foreignKey: DepartmentID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Courses []Course `json:"courses,omitempty"`
}