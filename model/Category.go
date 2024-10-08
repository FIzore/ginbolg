package model

import (
	"errors"
	"ginblog/utils/errmsg"
	"gorm.io/gorm"
	"log"
)

// Category 分类
type Category struct {
	ID   uint   `gorm:"primary_key;auto_increment" json:"id"`
	Name string `gorm:"type:varchar(20);not null" json:"name"`
}

// CreateCategory 新增分类
func CreateCategory(data *Category) int {
	err := db.Create(&data).Error
	if err != nil {
		return errmsg.ERROR //500
	}
	return errmsg.SUCCESS //200

}

// CheckCate 查询分类是否存在
func CheckCate(name string) int {
	var cate Category
	db.Select("id").Where("name = ?", name).First(&cate) //todo 优化RowsAffected
	if cate.ID > 0 {
		return errmsg.ERROR_CATEGORY_USED //1001
	}
	return errmsg.SUCCESS //200
}

// GetCate 查询分类列表
func GetCate(pageSize int, pageNum int) []Category {
	var cate []Category
	err := db.Limit(pageSize).Offset((pageNum - 1) * pageSize).Find(&cate).Error //分页查询
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		log.Printf("Error fetching category: %v", err)
		return []Category{} // 返回空的切片而不是 nil
	}
	return cate
}

// EditCate 编辑分类
func EditCate(id int, data *Category) int {
	var cate Category
	var maps = make(map[string]interface{})
	maps["name"] = data.Name
	err = db.Model(&cate).Where("id = ?", id).Updates(maps).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}

// DeleteCate 删除分类
func DeleteCate(id int) int {
	var cate Category
	err = db.Where("id = ?", id).Delete(&cate).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}

//todo 查询分类下的文章
