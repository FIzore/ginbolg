package model

import (
	"errors"
	"ginblog/utils/errmsg"
	"gorm.io/gorm"
	"log"
)

type Article struct {
	Category Category `gorm:"foreignKey:Cid;"`
	gorm.Model
	Title   string `gorm:"type:varchar(100);not null" json:"title"`
	Desc    string `gorm:"type:varchar(200)" json:"desc"` //简述
	Content string `gorm:"type:longtext" json:"content"`  //todo longtext?
	Img     string `gorm:"type:varchar(100)" json:"img"`
	Cid     int    `gorm:"type:int;not nul" json:"cid"` //分类
}

// CreateArticle 新增文章
func CreateArticle(data *Article) int {
	err := db.Create(&data).Error
	if err != nil {
		return errmsg.ERROR //500
	}
	return errmsg.SUCCESS //200

}

// GetCateArt 查询分类下所有文章
func GetCateArt(id int, pageSize int, pageNum int) ([]Article, int) {
	var cateArtList []Article
	err = db.Preload("Category").Limit(pageSize).Offset((pageNum-1)*pageSize).Where("cid = ?", id).Find(&cateArtList).Error
	if err != nil {
		return nil, errmsg.ERROR_ART_NOT_EXIST //todo 可以加一个分类是否存在的前提检测
	}
	return cateArtList, errmsg.SUCCESS
}

// GetArtInfo 查询单个文章
func GetArtInfo(id int) (Article, int) { //使用文章标题查询
	var art Article
	err = db.Preload("Category").Where("id = ?", id).First(&art).Error
	if err != nil {
		return art, errmsg.ERROR_ART_NOT_EXIST
	}
	return art, errmsg.SUCCESS
}

// GetArt 查询文章列表
func GetArt(pageSize int, pageNum int) ([]Article, int) {
	var articlelist []Article
	err := db.Preload("Category").Limit(pageSize).Offset((pageNum - 1) * pageSize).Find(&articlelist).Error //分页查询
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		log.Printf("Error fetching category: %v", err)
		return []Article{}, errmsg.ERROR // 返回空的切片而不是 nil
	}
	return articlelist, errmsg.SUCCESS
}

// EditArt 编辑文章
func EditArt(id int, data *Article) int {
	var art Article
	var maps = make(map[string]interface{})
	maps["title"] = data.Title
	maps["desc"] = data.Desc
	maps["content"] = data.Content
	maps["img"] = data.Img
	maps["cid"] = data.Cid
	err = db.Model(&art).Where("id = ?", id).Updates(maps).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}

// DeleteArt 删除文章
func DeleteArt(id int) int {
	var art Article
	err = db.Where("id = ?", id).Delete(&art).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}
