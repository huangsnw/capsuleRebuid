package main

import (
	"capsuleRebuild/config"
	"capsuleRebuild/routers"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func init() {
	config.LogConfig()
	config.ViperConfig()
	config.InitMongoDB()
	log.Println("config init over")
}

func main() {
	engine := gin.Default()
	engine = routers.AssembleRouters(engine)
	engine.Run(":" + viper.GetString("httpport"))
}
