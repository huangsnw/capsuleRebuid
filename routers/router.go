package routers

import (
	"capsuleRebuild/controllers"

	"github.com/gin-gonic/gin"
)

func AssembleRouters(engine *gin.Engine) *gin.Engine {
	// 路由组
	statistic := engine.Group("/statistic")
	// 1. 详细数据
	statistic.GET("/detailDate", controllers.DetailBloodPressure)

	return engine
}
