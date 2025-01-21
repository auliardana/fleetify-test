package config

import (
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func InitGin(config *viper.Viper) *gin.Engine {
	gin.SetMode(config.GetString("gin.mode"))

	return gin.Default()
}