//go:build wireinject
// +build wireinject

package server

import (
	"github.com/auliardana/fleetify-test/internal/config"
	"github.com/auliardana/fleetify-test/internal/delivery/http/handler"
	"github.com/auliardana/fleetify-test/internal/delivery/http/route"
	"github.com/auliardana/fleetify-test/internal/repository"
	"github.com/auliardana/fleetify-test/internal/service"
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
)

var configSet = wire.NewSet(config.LoadConfig, config.NewLogger, config.NewDatabase, config.NewValidator)
var repositorySet = wire.NewSet(repository.NewEmployeeRepository, repository.NewDepartementRepository, repository.NewAttendanceRepository, repository.NewAttendanceHistoryRepository)
var serviceSet = wire.NewSet(service.NewEmployeeService, service.NewDepartementService, service.NewAttendanceService, service.NewAttendanceHistoryService)
var handlerSet = wire.NewSet(handler.NewEmployeeHandler, handler.NewDepartementHandler, handler.NewAttendanceHandler, handler.NewAttendanceHistoryHandler)

func InitializeServer() *gin.Engine {
	wire.Build(configSet, repositorySet, serviceSet, handlerSet, route.NewRoute, config.NewApp)
	return nil
}
