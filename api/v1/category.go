package v1

import (
	"ginblog/model"
	"ginblog/utils/errmsg"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// AddCategory 添加 分类
func AddCategory(c *gin.Context) {
	var data model.Category
	_ = c.ShouldBindJSON(&data)
	code = model.CheckCate(data.Name)
	if code == errmsg.SUCCESS {
		model.CreateCategory(&data)
	}
	if code == errmsg.ERROR_CATEGORY_USED {
		code = errmsg.ERROR_CATEGORY_USED
	}
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"message": errmsg.GetErrMsg(code),
	})
}

//todo 查询分类下的所有文章

// GetCate 查询 分类列表
func GetCate(c *gin.Context) {
	pageSize, _ := strconv.Atoi(c.Query("pageSize"))
	pageNum, _ := strconv.Atoi(c.Query("pageNum"))
	if pageSize == 0 {
		pageSize = -1 //查询所有数据,取消分页限制
	}
	if pageNum == 0 {
		pageNum = -1 //取消分页限制
	}
	data := model.GetCate(pageSize, pageNum)
	code = errmsg.SUCCESS
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"data":    data,
		"message": errmsg.GetErrMsg(code),
	})
}

// EditCate 编辑 分类
func EditCate(c *gin.Context) { //todo username不能重复或许使用unique约束
	var data model.Category
	id, _ := strconv.Atoi(c.Param("id"))
	c.ShouldBindJSON(&data)           //判断一下如果 分类名存在且id不是当前 分类id就说明真实存在，不能修改
	code = model.CheckCate(data.Name) //目前这个逻辑如果我不修改 分类名，就无法修改邮箱和头像等其他信息
	if code == errmsg.SUCCESS {       //todo 逻辑问题优化
		model.EditCate(id, &data)
	} else {
		c.Abort()
	}
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"message": errmsg.GetErrMsg(code),
	})
}

// DeleteCate 删除 分类
func DeleteCate(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	code = model.DeleteCate(id)
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"message": errmsg.GetErrMsg(code),
	})
}

//查询单个分类下的文章

//查询分类列表
