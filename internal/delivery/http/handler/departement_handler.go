package handler

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/auliardana/fleetify-test/internal/model"
	"github.com/auliardana/fleetify-test/internal/service"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type DepartementHandler interface {
	//CRUD employee
	CreateDepartement(c *gin.Context)
	ListDepartement(c *gin.Context)
	// GetDepartement(c *gin.Context)
	UpdateDepartement(c *gin.Context)
	DeleteDepartement(c *gin.Context)
}

type departementHandler struct {
	Service  service.DepartementService
	Log      *logrus.Logger
	Validate *validator.Validate
}

func NewDepartementHandler(service service.DepartementService, logger *logrus.Logger, validate *validator.Validate) DepartementHandler {
	return &departementHandler{
		Service:  service,
		Log:      logger,
		Validate: validate,
	}
}

func (h *departementHandler) CreateDepartement(c *gin.Context) {
	departementRequest := new(model.DepartementRequest)
	err := c.ShouldBindJSON(&departementRequest)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid input", "details": err.Error()})
		return
	}

	err = h.Service.CreateDepartement(c, departementRequest)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "internal server error"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "departement created successfully"})
}

func (h *departementHandler) ListDepartement(c *gin.Context) {
	departements, err := h.Service.ListDepartement(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "internal server error"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": departements})
}

func (h *departementHandler) UpdateDepartement(c *gin.Context) {
	departementRequest := new(model.DepartementUpdateRequest)

	idParam := c.Param("id")
	if idParam == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "id is required"})
		return
	}

	err := c.ShouldBindJSON(&departementRequest)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid input", "details": err.Error()})
		return
	}

	departementRequest.ID, err = strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "id must be a number"})
		return
	}

	err = h.Service.UpdateDepartement(c, departementRequest)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "internal server error"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "departement updated successfully"})
}

func (h *departementHandler) DeleteDepartement(c *gin.Context) {
	idParam := c.Param("id")
	if idParam == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "id is required"})
		return
	}

	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "id must be a number"})
		return
	}

	err = h.Service.DeleteDepartement(c, id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"error": "departement not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "internal server error"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "departement deleted successfully"})
}
