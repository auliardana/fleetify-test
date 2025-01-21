package handler

import (
	"net/http"

	"github.com/auliardana/fleetify-test/internal/model"
	"github.com/auliardana/fleetify-test/internal/service"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/sirupsen/logrus"
)

type AttendanceHandler interface {
	HandleClockIn(c *gin.Context)
	HandleClockOut(c *gin.Context)
}

type attendanceHandler struct {
	Service  service.AttendanceService
	Log      *logrus.Logger
	Validate *validator.Validate
}

func NewAttendanceHandler(service service.AttendanceService, logger *logrus.Logger, validate *validator.Validate) AttendanceHandler {
	return &attendanceHandler{
		Service:  service,
		Log:      logger,
		Validate: validate,
	}
}

func (h *attendanceHandler) HandleClockIn(c *gin.Context) {

	attendanceRequest := new(model.ClockInAttendanceRequest)
	err := c.ShouldBindJSON(&attendanceRequest)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid format", "details": err.Error()})
		return
	}

	if err = h.Validate.Struct(attendanceRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input", "details": err.Error()})
		return
	}

	// idParam := c.Param("id")

	err = h.Service.ClockIn(c, attendanceRequest)
	if err != nil {
		h.Log.Warn("Failed to clock in: ", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "internal server error"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "clock in success"})

}

func (h *attendanceHandler) HandleClockOut(c *gin.Context) {
	idParam := c.Param("id")

	err := h.Service.ClockOut(c, idParam)
	if err != nil {
		h.Log.Warn("Failed to clock out: ", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "internal server error"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "clock out success"})
}
