package model

import "gorm.io/gorm"

// Category 分类
type Category struct {
	gorm.Model
	Name string `gorm:"type:varchar(20);not null" json:"name"`
}
