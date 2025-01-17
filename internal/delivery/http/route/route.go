package route

import (
	"github.com/auliardana/fleetify-test/internal/delivery/http/handler"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

type ConfigRoute struct {
	EmployeeHandler          handler.EmployeeHandler
	DepartementHandler       handler.DepartementHandler
	AttendanceHandler        handler.AttendanceHandler
	AttendanceHistoryHandler handler.AttendanceHistoryHandler
}

// , attendanceHandler handler.AttendanceHandler, attendanceHistoryHandler handler.AttendanceHistoryHandler

func NewRoute(employeeHandler handler.EmployeeHandler, departementHandler handler.DepartementHandler, attendanceHandler handler.AttendanceHandler, attendanceHistoryHandler handler.AttendanceHistoryHandler) *ConfigRoute {
	return &ConfigRoute{
		EmployeeHandler:          employeeHandler,
		DepartementHandler:       departementHandler,
		AttendanceHandler:        attendanceHandler,
		AttendanceHistoryHandler: attendanceHistoryHandler,
	}
}

func (c *ConfigRoute) Setup(app *gin.Engine) {
	app.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"}, // Bisa diganti sesuai dengan origin yang diizinkan
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))

	c.guestApiRoute(app)
	// c.protectedApiRoute(app)
	// c.adminApiRoute(app)
}

func (c *ConfigRoute) guestApiRoute(app *gin.Engine) {
	api := apiGroup(app)
	{
		// employee
		api.POST("/employee", c.EmployeeHandler.CreateEmployee)
		api.GET("/employee", c.EmployeeHandler.ListEmployee)
		// api.GET("/employee/:id", c.EmployeeHandler.GetEmployeeByID)
		api.PATCH("/employee/:id", c.EmployeeHandler.UpdateEmployee)
		api.DELETE("/employee/:id", c.EmployeeHandler.DeleteEmployee)

		// // departement
		api.POST("/departement", c.DepartementHandler.CreateDepartement)
		// api.GET("/departement/:id", c.DepartementHandler.GetDepartementByID)
		api.GET("/departement", c.DepartementHandler.ListDepartement)
		api.PATCH("/departement/:id", c.DepartementHandler.UpdateDepartement)
		api.DELETE("/departement/:id", c.DepartementHandler.DeleteDepartement)

		// attendance
		api.POST("/attendance", c.AttendanceHandler.HandleClockIn)
		api.PUT("/attendance/:id", c.AttendanceHandler.HandleClockOut)

		// // attendance history
		api.GET("/attendance", c.AttendanceHistoryHandler.ListAttendanceHistory)

	}
}

// func (c *ConfigRoute) protectedApiRoute(app *gin.Engine) {
// 	api := apiGroup(app).Use(c.AuthMiddleware.TokenAuth)
// 	{
// 		// employee
// 		api.GET("/employee", c.EmployeeHandler.GetProducts)

// 	}
// }

// func (c *ConfigRoute) adminApiRoute(app *gin.Engine) {
// 	api := apiGroup(app).Use(c.AuthMiddleware.AdminAuth)
// 	{
// 		api.PATCH("/admin/orders-confirm/:id", c.OrderHandler.ConfirmOrder)
// 		api.GET("/admin/orders", c.OrderHandler.GetOrders)
// 	}
// }

func apiGroup(app *gin.Engine) *gin.RouterGroup {
	return app.Group("/api/v1")
}
