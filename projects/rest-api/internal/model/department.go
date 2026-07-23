package model


type Department struct{
	Base
	Name string `json:"name"`
	DepartmentCode string `json:"department_code"`
}