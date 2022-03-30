package models

import "github.com/jinzhu/gorm"

type Group struct {
	gorm.Model
	GroupName  string `json:"group_name"`
	SubName    string `json:"sub_name"`
	AbbrName   string `json:"abbr_name"`
	SubGroupTo *Group `json:"sub_group_to"`
}
