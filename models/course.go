package models

import "github.com/jinzhu/gorm"

type (
	CourseModel struct {
		gorm.Model
		MainTitle           string `json:"main_title"`
		SubTitle            string `json:"sub_title"`
		CourseCode          string `json:"course_code"`
		CourseUnits         uint   `json:"course_units"`
		CourseDescription   string `json:"course_description"`
		CourseType          string `json:"course_type"`
		CourseIsVisible     bool   `json:"course_is_visible"`
		CourseAcceptComment bool   `json:"course_accept_comment"`
	}
)
