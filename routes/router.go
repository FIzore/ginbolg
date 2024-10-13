package routes

import (
	"ginblog/api/v1"
	"ginblog/middleware"
	"ginblog/utils"
	"github.com/gin-gonic/gin"
)

func InitRouter() {
	gin.SetMode(utils.AppMode)
	r := gin.Default()

	auth := r.Group("/api/v1")
	auth.Use(middleware.JwtToken())
	{
		//user模块的路由接口
		auth.PUT("user/:id", v1.EditUser)      //编辑用户
		auth.DELETE("user/:id", v1.DeleteUser) //删除用户
		//文章模块的路由接口
		auth.POST("article/add", v1.AddArticle)  //新增文章
		auth.PUT("article/:id", v1.EditArt)      //编辑文章
		auth.DELETE("article/:id", v1.DeleteArt) //删除文章
		//分类模块的路由接口
		auth.POST("category/add", v1.AddCategory)  //新增分类
		auth.PUT("category/:id", v1.EditCate)      //编辑分类
		auth.DELETE("category/:id", v1.DeleteCate) //删除分类
		//上传文件
		auth.POST("upload", v1.UpLoad)
	}
	router := r.Group("api/v1")
	{
		router.POST("user/add", v1.AddUser)                   //新增用户
		router.GET("users", v1.GetUsers)                      //查询用户列表
		router.GET("category", v1.GetCate)                    //查询分类列表
		router.GET("article", v1.GetArt)                      //查询文章列表
		router.GET("article/CategoryList/:id", v1.GetCateArt) //查询分类下的所有文章
		router.GET("article/info/:id", v1.GetArtInfo)         //查询单个文章
		router.POST("login", v1.Login)                        //登录
	}
	_ = r.Run(utils.HttpPort)
}
