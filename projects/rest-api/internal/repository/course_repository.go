package repository

import (
	"context"

	"github.com/ShyamSundhar1411/rest-api/internal/model"
	"gorm.io/gorm"
)

type CourseRepository interface {
	GetByID(ctx context.Context, id uint)(*model.Course, error)
	GetAll(ctx context.Context)(*[]model.Course, error)
	Create(ctx context.Context, courseData *model.Course)(*model.Course, error)
	Update(ctx context.Context, id uint, courseData *model.Course)(*model.Course, error)
	Delete(ctx context.Context, id uint) error
	Search(ctx context.Context, search string)(*[]model.Course, error)
}
type courseRepository struct{
	db *gorm.DB
}

func NewCourseRepository(db *gorm.DB) CourseRepository {
	return &courseRepository{db: db}
}

func (courseRepository *courseRepository) GetByID(ctx context.Context,id uint)(*model.Course, error){
	var course model.Course
	err := courseRepository.db.WithContext(ctx).First(&course,id).Error
	if err!=nil{
		return nil,err
	}
	return &course,nil
}

func (courseRepository *courseRepository) GetAll(ctx context.Context)(*[]model.Course,error){
	var courses[] model.Course
	err := courseRepository.db.WithContext(ctx).Find(&courses).Error
	if err != nil {
		return nil,err
	}
	return &courses,nil
}

func(courseRepository *courseRepository) Create(ctx context.Context,courseData *model.Course)(*model.Course,error){
	err := courseRepository.db.WithContext(ctx).Create(&courseData).Error
	if err != nil{
		return nil,err
	}
	return courseData,nil

}
func(courseRepository *courseRepository) Update(ctx context.Context,id uint, courseData *model.Course)(*model.Course,error){
	err := courseRepository.db.Model(&model.Course{}).WithContext(ctx).Where("id = ?",id).Select("*").Updates(courseData).Error
	if err != nil{
		return nil,err
	}
	return courseData,nil
}
func(courseRepository *courseRepository) Delete(ctx context.Context, id uint) error{
	err := courseRepository.db.WithContext(ctx).Delete(&model.Course{}, id).Error
	if err != nil{
		return err
	}
	return nil
}
func(courseRepository *courseRepository) Search(ctx context.Context, search string)(*[]model.Course, error){
	var courses []model.Course
	searchPattern := "%" + search + "%"
	err := courseRepository.db.WithContext(ctx).Where("course_name ILIKE ? OR course_code ILIKE ?",searchPattern,searchPattern).Find(&courses).Error
	if err != nil{
		return nil,err
	}
	return &courses,nil
}
