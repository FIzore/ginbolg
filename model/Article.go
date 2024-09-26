package model

import "gorm.io/gorm"

type Article struct {
	Category Category `gorm:"foreignKey:Cid;"`
	gorm.Model
	Title   string `gorm:"type:varchar(100);not null" json:"title"`
	Desc    string `gorm:"type:varchar(200)" json:"desc"` //简述
	Content string `gorm:"type:longtext" json:"content"`  //todo longtext?
	Img     string `gorm:"type:varchar(100)" json:"img"`
	Cid     int    `gorm:"type:int;not nul" json:"cid"` //分类
}
