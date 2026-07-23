package model

type Course struct{
	Base
	CourseName string `json:"course_name" gorm:"type:varchar(255) not null"`
	CourseCode string `json:"course_code" gorm:"type:varchar(255) not null;unique"`
}