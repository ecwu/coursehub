package models

import "github.com/jinzhu/gorm"

type GroupModel struct {
	gorm.Model
	GroupName    string      `json:"group_name"`
	SubName      string      `json:"sub_name"`
	AbbrName     string      `json:"abbr_name"`
	SubGroupTo   *GroupModel `json:"sub_group_to" gorm:"foreignKey:GroupModelID"`
	GroupModelID *uint
}
