package repository

import "gorm.io/gorm"
type Repositories struct{
	Course CourseRepository
}

func InitializeRepositories(db *gorm.DB) *Repositories{
	return &Repositories{
		Course: NewCourseRepository(db),
	}
}