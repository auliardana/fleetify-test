package handler

import (
	"errors"
	"net/http"
	"strings"

	"github.com/auliardana/fleetify-test/internal/model"
	"github.com/auliardana/fleetify-test/internal/service"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type EmployeeHandler interface {
	//CRUD employee
	CreateEmployee(c *gin.Context)
	ListEmployee(c *gin.Context)
	// GetEmployee(c *gin.Context)
	UpdateEmployee(c *gin.Context)
	DeleteEmployee(c *gin.Context)
}

type employeeHandler struct {
	Service  service.EmployeeService
	Log      *logrus.Logger
	Validate *validator.Validate
}

func NewEmployeeHandler(service service.EmployeeService, logger *logrus.Logger, validate *validator.Validate) EmployeeHandler {
	return &employeeHandler{
		Service:  service,
		Log:      logger,
		Validate: validate,
	}
}

func (h *employeeHandler) CreateEmployee(c *gin.Context) {
	employeeRequest := new(model.EmployeeRequest)
	err := c.ShouldBindJSON(&employeeRequest)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid format", "details": err.Error()})
		return
	}

	if err := h.Validate.Struct(employeeRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input", "details": err.Error()})
		return
	}

	err = h.Service.CreateEmployee(c, employeeRequest)
	if err != nil {
		if strings.Contains(err.Error(), "fk_employees_departement") {
			c.JSON(http.StatusBadRequest, gin.H{"error": "departement_id invalid"})
			return
		}
		h.Log.Warn("Failed to create employee: ", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "internal server error"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "employee created successfully"})

}

// func (h *employeeHandler) GetEmployee(c *gin.Context) {
// 	employees, err := h.Service.

// }

func (h *employeeHandler) ListEmployee(c *gin.Context) {
	employees, err := h.Service.ListEmployee(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "internal server error"})
		return

	}

	c.JSON(http.StatusOK, gin.H{"data": employees})
}

func (h *employeeHandler) UpdateEmployee(c *gin.Context) {
	employeeRequest := new(model.EmployeeUpdateRequest)

	idParam := c.Param("id")
	if idParam == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "id is required"})
		return
	}

	err := c.ShouldBindJSON(&employeeRequest)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid input", "details": err.Error()})
		return
	}

	parsedID, err := uuid.Parse(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}

	employeeRequest.ID = parsedID

	err = h.Service.UpdateEmployee(c, employeeRequest)
	if err != nil {
		if strings.Contains(err.Error(), "fk_employees_departement") {
			c.JSON(http.StatusBadRequest, gin.H{"error": "departement_id invalid"})
			return
		}

		c.JSON(http.StatusInternalServerError, gin.H{"error": "internal server error"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "employee updated successfully"})

}

func (h *employeeHandler) DeleteEmployee(c *gin.Context) {
	idParam := c.Param("id")
	if idParam == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "id is required"})
		return
	}

	parsedID, err := uuid.Parse(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}

	err = h.Service.DeleteEmployee(c, parsedID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"error": "employee not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "internal server error"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "employee deleted successfully"})

}
