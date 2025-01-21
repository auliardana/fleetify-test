package handler

import (
	"net/http"

	"github.com/auliardana/fleetify-test/internal/model"
	"github.com/auliardana/fleetify-test/internal/service"
	"github.com/auliardana/fleetify-test/internal/utils"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
)

type AttendanceHistoryHandler interface {
	ListAttendanceHistory(c *gin.Context)
}

type attendanceHistoryHandler struct {
	Service  service.AttendanceHistoryService
	Log      *logrus.Logger
	Validate *validator.Validate
}

func NewAttendanceHistoryHandler(service service.AttendanceHistoryService, logger *logrus.Logger, validate *validator.Validate) AttendanceHistoryHandler {
	return &attendanceHistoryHandler{
		Service:  service,
		Log:      logger,
		Validate: validate,
	}
}

func (h *attendanceHistoryHandler) ListAttendanceHistory(c *gin.Context) {
	queries := new(model.AttendanceHistoryFilter)

	// Query parameters
	startDateStr := c.Query("start_date")
	endDateStr := c.Query("end_date")
	departementIDStr := c.Query("departement_id")

	// Parse and validate dates using utils
	startDate, endDate, err := utils.ParseDateRange(startDateStr, endDateStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	queries.StartDate = startDate
	queries.EndDate = endDate
	departementID, err := uuid.Parse(departementIDStr)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Assign UUID ke DepartementID
	queries.DepartementID = departementID

	attendanceHistories, err := h.Service.ListAttendanceHistory(c, queries)
	if err != nil {
		h.Log.Warn("Failed to get attendance histories: ", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "internal server error"})
		return
	}

	// Format response to include punctuality status
	response := []map[string]interface{}{}
	for _, history := range attendanceHistories {
		isLateClockIn := history.Attendance.ClockIn.After(history.Employee.Departement.MaxClockInTime)
		isLateClockOut := history.Attendance.ClockOut.After(history.Employee.Departement.MaxClockOutTime)

		response = append(response, map[string]interface{}{
			"employee_name":     history.Employee.Name,
			"department_name":   history.Employee.Departement.DepartementName,
			"date_attendance":   history.DateAttendance,
			"clock_in":          history.Attendance.ClockIn,
			"clock_out":         history.Attendance.ClockOut,
			"is_late_clock_in":  isLateClockIn,
			"is_late_clock_out": isLateClockOut,
			"description":       history.Description,
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"data": response,
	})

}
