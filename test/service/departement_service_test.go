package service_test

import (
	"errors"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/auliardana/fleetify-test/internal/entity"
	"github.com/auliardana/fleetify-test/internal/model"
	service "github.com/auliardana/fleetify-test/internal/service"
	mockRepo "github.com/auliardana/fleetify-test/test/mock"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func TestCreateDepartement(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repo := mockRepo.NewMockDepartementRepository(ctrl)
	logger := logrus.New()
	service := service.NewDepartementService(repo, logger)

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	req := &model.DepartementRequest{
		DepartementName: "HR",
		MaxClockInTime:  time.Date(2025, time.January, 21, 9, 0, 0, 0, time.UTC),
		MaxClockOutTime: time.Date(2025, time.January, 21, 17, 0, 0, 0, time.UTC),
	}

	repo.EXPECT().
		Create(gomock.Any(), gomock.Any()).
		Return(nil).
		Times(1)

	err := service.CreateDepartement(c, req)

	assert.NoError(t, err)
	assert.Equal(t, 200, w.Code)
}

func TestListDepartement(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repo := mockRepo.NewMockDepartementRepository(ctrl)
	logger := logrus.New()
	service := service.NewDepartementService(repo, logger)

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	repo.EXPECT().
		GetAll(gomock.Any()).
		Return([]entity.Departement{
			{
				DepartementName: "HR",
				MaxClockInTime:  time.Date(2025, time.January, 21, 9, 0, 0, 0, time.UTC),
				MaxClockOutTime: time.Date(2025, time.January, 21, 17, 0, 0, 0, time.UTC),
			},
		}, nil).
		Times(1)

	departments, err := service.ListDepartement(c)

	assert.NoError(t, err)
	assert.Equal(t, 1, len(departments)) // Expect 1 department returned
	assert.Equal(t, "HR", departments[0].DepartementName)
}

func TestListDepartement_ShouldReturnError_WhenRepositoryFails(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	// Mock repository dan logger
	repo := mockRepo.NewMockDepartementRepository(ctrl)
	logger := logrus.New()
	service := service.NewDepartementService(repo, logger)

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	repo.EXPECT().
		GetAll(gomock.Any()).
		Return(nil, errors.New("failed to fetch departements")).
		Times(1)

	departments, err := service.ListDepartement(c)

	assert.Error(t, err)
	assert.Nil(t, departments) // Expect no departments returned
}

func TestUpdateDepartement(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	// Mock repository dan logger
	repo := mockRepo.NewMockDepartementRepository(ctrl)
	logger := logrus.New()
	service := service.NewDepartementService(repo, logger)

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	req := &model.DepartementUpdateRequest{
		ID:              uuid.New(),
		DepartementName: "Finance",
		MaxClockInTime:  time.Date(2025, time.January, 21, 8, 0, 0, 0, time.UTC),
		MaxClockOutTime: time.Date(2025, time.January, 21, 16, 0, 0, 0, time.UTC),
	}

	repo.EXPECT().
		FindById(gomock.Any(), req.ID).
		Return(&entity.Departement{
			ID:              req.ID,
			DepartementName: "HR",
			MaxClockInTime:  time.Date(2025, time.January, 21, 9, 0, 0, 0, time.UTC),
			MaxClockOutTime: time.Date(2025, time.January, 21, 17, 0, 0, 0, time.UTC),
		}, nil).
		Times(1)

	repo.EXPECT().
		Update(gomock.Any(), gomock.Any()).
		Return(nil).
		Times(1)

	err := service.UpdateDepartement(c, req)

	assert.NoError(t, err)
}

func TestUpdateDepartement_ShouldReturnError_WhenFindByIdFails(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repo := mockRepo.NewMockDepartementRepository(ctrl)
	logger := logrus.New()
	service := service.NewDepartementService(repo, logger)

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	req := &model.DepartementUpdateRequest{
		ID:              uuid.New(),
		DepartementName: "Finance",
		MaxClockInTime:  time.Date(2025, time.January, 21, 8, 0, 0, 0, time.UTC),
		MaxClockOutTime: time.Date(2025, time.January, 21, 16, 0, 0, 0, time.UTC),
	}

	repo.EXPECT().
		FindById(gomock.Any(), req.ID).
		Return(nil, errors.New("departement not found")).
		Times(1)

	err := service.UpdateDepartement(c, req)

	assert.Error(t, err)
}

func TestUpdateDepartement_ShouldReturnError_WhenUpdateFails(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repo := mockRepo.NewMockDepartementRepository(ctrl)
	logger := logrus.New()
	service := service.NewDepartementService(repo, logger)

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	req := &model.DepartementUpdateRequest{
		ID:              uuid.New(),
		DepartementName: "Finance",
		MaxClockInTime:  time.Date(2025, time.January, 21, 8, 0, 0, 0, time.UTC),
		MaxClockOutTime: time.Date(2025, time.January, 21, 16, 0, 0, 0, time.UTC),
	}

	repo.EXPECT().
		FindById(gomock.Any(), req.ID).
		Return(&entity.Departement{
			ID:              req.ID,
			DepartementName: "HR",
			MaxClockInTime:  time.Date(2025, time.January, 21, 9, 0, 0, 0, time.UTC),
			MaxClockOutTime: time.Date(2025, time.January, 21, 17, 0, 0, 0, time.UTC),
		}, nil).
		Times(1)

	repo.EXPECT().
		Update(gomock.Any(), gomock.Any()).
		Return(errors.New("failed to update departement")).
		Times(1)

	err := service.UpdateDepartement(c, req)

	assert.Error(t, err)
}


func TestDeleteDepartement(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	// Mock repository dan logger
	repo := mockRepo.NewMockDepartementRepository(ctrl)
	logger := logrus.New()
	service := service.NewDepartementService(repo, logger)

	// Gunakan context dengan recorder
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	// Data ID departemen yang akan dihapus
	id := 1

	// Ekspektasi panggilan ke repository untuk Delete
	repo.EXPECT().
		Delete(gomock.Any(), id).
		Return(nil).
		Times(1)

	// Jalankan layanan
	err := service.DeleteDepartement(c, id)

	// Validasi hasil
	assert.NoError(t, err)
}

func TestDeleteDepartement_ShouldReturnError_WhenDeleteFails(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	// Mock repository dan logger
	repo := mockRepo.NewMockDepartementRepository(ctrl)
	logger := logrus.New()
	service := service.NewDepartementService(repo, logger)

	// Gunakan context dengan recorder
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	// Data ID departemen yang akan dihapus
	id := 1

	// Simulasi error pada Delete
	repo.EXPECT().
		Delete(gomock.Any(), id).
		Return(errors.New("failed to delete departement")).
		Times(1)

	// Jalankan layanan
	err := service.DeleteDepartement(c, id)

	// Validasi hasil error
	assert.Error(t, err)
}
