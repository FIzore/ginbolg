package v1

import (
	"ginblog/middleware"
	"ginblog/model"
	"ginblog/utils/errmsg"
	"github.com/gin-gonic/gin"
	"net/http"
)

// Login 查询用户是否存在
func Login(c *gin.Context) {
	var data model.User
	var token string
	var code int
	err := c.ShouldBindJSON(&data)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  errmsg.ERROR,
			"message": errmsg.GetErrMsg(errmsg.ERROR),
		})
		return
	}
	code = model.CheckLogin(data.Username, data.Password)
	if code == errmsg.SUCCESS {
		token, code = middleware.SetToken(data.Username)
	}
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"message": errmsg.GetErrMsg(code),
		"token":   token,
	})
}

//添加用户

//查询单个用户

//查询用户列表

//编辑用户

//删除用户
