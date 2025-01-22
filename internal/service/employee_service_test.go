package service_test

import (
	"context"
	"errors"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/auliardana/fleetify-test/internal/entity"
	"github.com/auliardana/fleetify-test/internal/model"
	"github.com/auliardana/fleetify-test/internal/service"
	mockRepo "github.com/auliardana/fleetify-test/test/mock"
	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
)

func TestCreateEmployee(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repo := mockRepo.NewMockEmployeeRepository(ctrl)
	logger := logrus.New()
	service := service.NewEmployeeService(repo, logger)

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	data := &model.EmployeeRequest{
		DepartementID: uuid.New(),
		Name:          "John Doe",
		Address:       "123 Street",
	}

	repo.EXPECT().
		Create(c, gomock.AssignableToTypeOf(&entity.Employee{})).
		DoAndReturn(func(_ context.Context, emp *entity.Employee) error {
			assert.Equal(t, data.DepartementID, emp.DepartementID)
			assert.Equal(t, data.Name, emp.Name)
			assert.Equal(t, data.Address, emp.Address)
			assert.NotNil(t, emp.ID)
			assert.NotNil(t, emp.CreatedAt)
			assert.NotNil(t, emp.UpdatedAt)
			return nil
		}).
		Times(1)

	err := service.CreateEmployee(c, data)

	assert.NoError(t, err)
}

func TestCreateEmployee_FailedCreate(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repo := mockRepo.NewMockEmployeeRepository(ctrl)
	logger := logrus.New()
	service := service.NewEmployeeService(repo, logger)

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	req := &model.EmployeeRequest{
		DepartementID: uuid.New(),
		Name:          "John Doe",
		Address:       "123 Street",
	}

	expectedEmployee := &entity.Employee{
		DepartementID: req.DepartementID,
		Name:          req.Name,
		Address:       req.Address,
		CreatedAt:     time.Now(),
		UpdatedAt:     time.Now(),
	}

	errorReturn := errors.New("failed to create employee")

	repo.EXPECT().
		Create(c, gomock.AssignableToTypeOf(&entity.Employee{})).
		DoAndReturn(func(_ context.Context, emp *entity.Employee) error {
			assert.Equal(t, expectedEmployee.DepartementID, emp.DepartementID)
			assert.Equal(t, expectedEmployee.Name, emp.Name)
			assert.Equal(t, expectedEmployee.Address, emp.Address)
			return errorReturn
		}).
		Times(1)

	err := service.CreateEmployee(c, req)

	assert.Error(t, err)
	assert.Equal(t, errorReturn, err)
}

func TestListEmployee_OK(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repo := mockRepo.NewMockEmployeeRepository(ctrl)
	logger := logrus.New()
	service := service.NewEmployeeService(repo, logger)

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	data := []entity.Employee{
		{
			ID:            uuid.New(),
			Name:          "John Doe",
			Address:       "123 Street",
			DepartementID: uuid.New(),
		},
	}

	repo.EXPECT().
		GetAll(gomock.Any()).
		Return(data, nil).
		Times(1)

	result, err := service.ListEmployee(c)

	assert.NotNil(t, result)
	assert.NoError(t, err)
}

func TestListEmployee_FailedGetList(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repo := mockRepo.NewMockEmployeeRepository(ctrl)
	logger := logrus.New()
	service := service.NewEmployeeService(repo, logger)

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	errorReturn := errors.New("failed to get employees")

	repo.EXPECT().
		GetAll(c).
		Return(nil, errorReturn).
		Times(1)

	result, err := service.ListEmployee(c)

	assert.Error(t, err)
	assert.Nil(t, result)
}

func TestUpdateEmployee_OK(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repo := mockRepo.NewMockEmployeeRepository(ctrl)
	logger := logrus.New()
	service := service.NewEmployeeService(repo, logger)

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	req := &model.EmployeeUpdateRequest{
		ID:            uuid.New(),
		DepartementID: uuid.New(),
		Name:          "John Doe NEW",
		Address:       "123 Street",
	}

	employee := &entity.Employee{
		ID:            req.ID,
		DepartementID: req.DepartementID,
		Name:          "John Doe",
		Address:       "123 Street",
		CreatedAt:     time.Now(),
		UpdatedAt:     time.Now(),
	}

	repo.EXPECT().
		FindById(gomock.Any(), employee.ID).
		Return(employee, nil).
		Times(1)

	repo.EXPECT().
		Update(gomock.Any(), gomock.Any()).
		Return(nil).
		Times(1)

	err := service.UpdateEmployee(c, req)

	assert.NoError(t, err)
}

func TestUpdateEmployee_InvalidDepartementID(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repo := mockRepo.NewMockEmployeeRepository(ctrl)
	logger := logrus.New()
	service := service.NewEmployeeService(repo, logger)

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	req := &model.EmployeeUpdateRequest{
		ID:            uuid.New(),
		DepartementID: uuid.New(),
		Name:          "John Doe NEW",
		Address:       "123 Street",
	}

	errorReturn := errors.New("failed to get Departement")

	repo.EXPECT().
		FindById(gomock.Any(), req.ID).
		Return(nil, errorReturn).
		Times(1)

	err := service.UpdateEmployee(c, req)

	assert.Error(t, err)
}

func TestUpdateEmployee_FailedUpdate(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repo := mockRepo.NewMockEmployeeRepository(ctrl)
	logger := logrus.New()
	service := service.NewEmployeeService(repo, logger)

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	req := &model.EmployeeUpdateRequest{
		ID:            uuid.New(),
		DepartementID: uuid.New(),
		Name:          "John Doe NEW",
		Address:       "123 Street",
	}

	employee := &entity.Employee{
		ID:            req.ID,
		DepartementID: req.DepartementID,
		Name:          "John Doe",
		Address:       "123 Street",
		CreatedAt:     time.Now(),
		UpdatedAt:     time.Now(),
	}

	updatedEmployee := &entity.Employee{
		ID:            req.ID,
		DepartementID: req.DepartementID,
		Name:          req.Name,
		Address:       req.Address,
		CreatedAt:     employee.CreatedAt,
		UpdatedAt:     time.Now(),
	}

	repo.EXPECT().
		FindById(c, req.ID).
		Return(employee, nil).
		Times(1)

	errorReturn := errors.New("failed to update employee")
	repo.EXPECT().
		Update(c, updatedEmployee).
		Return(errorReturn).
		Times(1)

	err := service.UpdateEmployee(c, req)

	assert.Error(t, err)
	assert.Equal(t, errorReturn, err)
}

func TestDeleteEmployee(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repo := mockRepo.NewMockEmployeeRepository(ctrl)
	logger := logrus.New()
	service := service.NewEmployeeService(repo, logger)

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	id := uuid.New()

	repo.EXPECT().
		Delete(gomock.Any(), id).
		Return(nil).
		Times(1)

	err := service.DeleteEmployee(c, id)

	assert.NoError(t, err)

}

func TestDeleteEmployee_Error(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repo := mockRepo.NewMockEmployeeRepository(ctrl)
	logger := logrus.New()
	service := service.NewEmployeeService(repo, logger)

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	id := uuid.New()

	errorReturn := errors.New("failed to delete employee")

	repo.EXPECT().
		Delete(gomock.Any(), id).
		Return(errorReturn).
		Times(1)

	err := service.DeleteEmployee(c, id)
	assert.Error(t, err)

}
