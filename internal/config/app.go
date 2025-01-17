package config

import (
	"github.com/auliardana/fleetify-test/internal/delivery/http/route"
	"github.com/gin-gonic/gin"
)

func NewApp(route *route.ConfigRoute) *gin.Engine {
	var app = gin.New()

	route.Setup(app)

	return app
}
