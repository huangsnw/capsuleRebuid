package routers

import (
	"capsuleRebuild/controllers"

	"github.com/gin-gonic/gin"
)

func AssembleRouters(engine *gin.Engine) *gin.Engine {
	statistic := engine.Group("/statistic")
	statistic.GET("/detailDate", controllers.DetailBloodPressure)
	statistic.GET("/oneBP", controllers.LatestBloodPressure)

	return engine
}
