package routes

import (
	"bi-activity/middleware"
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	router := gin.Default()
	router.Use(middleware.CORSMiddleware())
	// TODO: 注册路由

	// 登录注册相关路由
	loginRegisterRouter(router)
	// 学院相关的路由
	College(router)

	return router
}
