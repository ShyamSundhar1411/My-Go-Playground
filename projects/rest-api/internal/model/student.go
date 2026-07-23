package model

type Student struct{
	BaseUser
	EmergencyContact string 
	GradeLevel int `json:"grade_level"`
	DepartmentID uint `json:"department_id" gorm:"not null"`
	Department Department `json:"department" gorm:"foreignKey: DepartmentID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

