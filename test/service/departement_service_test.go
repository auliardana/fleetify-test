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
	testCases := []struct {
		name          string
		req           *model.DepartementRequest
		errorValue    error
		checkResponse func(t *testing.T, err error)
	}{
		{
			name: "OK",
			req: &model.DepartementRequest{
				DepartementName: "HR",
				MaxClockInTime:  time.Date(2025, time.January, 21, 9, 0, 0, 0, time.UTC),
				MaxClockOutTime: time.Date(2025, time.January, 21, 17, 0, 0, 0, time.UTC),
			},
			errorValue: nil,
			checkResponse: func(t *testing.T, err error) {
				assert.NoError(t, err)
			},
		},

		{
			name: "Invalid Request",
			req: &model.DepartementRequest{
				DepartementName: "HR",
				MaxClockInTime:  time.Date(2025, time.January, 21, 20, 0, 0, 0, time.UTC),
				MaxClockOutTime: time.Date(2025, time.January, 21, 17, 0, 0, 0, time.UTC),
			},
			errorValue: errors.New("invalid request"),
			checkResponse: func(t *testing.T, err error) {
				assert.Error(t, err)
			},
		},
		// Add more test cases if necessary
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			repo := mockRepo.NewMockDepartementRepository(ctrl)
			logger := logrus.New()
			service := service.NewDepartementService(repo, logger)

			w := httptest.NewRecorder()
			c, _ := gin.CreateTestContext(w)

			repo.EXPECT().
				Create(gomock.Any(), gomock.Any()).
				Return(tc.errorValue).
				Times(1)

			// Call the service method
			err := service.CreateDepartement(c, tc.req)
			// Check response
			tc.checkResponse(t, err)
		})
	}
}

func TestListDepartement(t *testing.T) {
	testCases := []struct {
		name          string
		returnValue   []entity.Departement
		errorValue    error
		checkError    func(t *testing.T, err error)
		checkResponse func(t *testing.T, departements []entity.Departement)
	}{
		{
			name:        "OK",
			returnValue: []entity.Departement{{}}, // Example departement
			errorValue:  nil,
			checkError: func(t *testing.T, err error) {
				assert.NoError(t, err)
			},
			checkResponse: func(t *testing.T, departements []entity.Departement) {
				assert.NotNil(t, departements)
			},
		},
		{
			name:        "Error",
			returnValue: nil,
			errorValue:  errors.New("failed to list departement"),
			checkError: func(t *testing.T, err error) {
				assert.Equal(t, "failed to list departement", err.Error())
			},
			checkResponse: func(t *testing.T, departements []entity.Departement) {
				assert.Nil(t, departements)
			},
		},
		// Add more test cases if necessary
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			repo := mockRepo.NewMockDepartementRepository(ctrl)
			logger := logrus.New()
			service := service.NewDepartementService(repo, logger)

			w := httptest.NewRecorder()
			c, _ := gin.CreateTestContext(w)

			// Correctly mock the repository call
			repo.EXPECT().
				GetAll(gomock.Any()).
				Return(tc.returnValue, tc.errorValue).
				Times(1)

			// Call the service method
			result, err := service.ListDepartement(c)

			// Check response
			tc.checkError(t, err)
			tc.checkResponse(t, result)
		})
	}
}

func TestUpdateDepartement(t *testing.T) {
	tests := []struct {
		name        string
		setupMocks  func(repo *mockRepo.MockDepartementRepository, req *model.DepartementUpdateRequest)
		req         *model.DepartementUpdateRequest
		expectError bool
	}{
		{
			name: "Should update department successfully",
			setupMocks: func(repo *mockRepo.MockDepartementRepository, req *model.DepartementUpdateRequest) {
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
			},
			req: &model.DepartementUpdateRequest{
				ID:              uuid.New(),
				DepartementName: "Finance",
				MaxClockInTime:  time.Date(2025, time.January, 21, 8, 0, 0, 0, time.UTC),
				MaxClockOutTime: time.Date(2025, time.January, 21, 16, 0, 0, 0, time.UTC),
			},
			expectError: false,
		},
		{
			name: "Should return error when FindById fails",
			setupMocks: func(repo *mockRepo.MockDepartementRepository, req *model.DepartementUpdateRequest) {
				repo.EXPECT().
					FindById(gomock.Any(), req.ID).
					Return(nil, errors.New("departement not found")).
					Times(1)
			},
			req: &model.DepartementUpdateRequest{
				ID:              uuid.New(),
				DepartementName: "Finance",
				MaxClockInTime:  time.Date(2025, time.January, 21, 8, 0, 0, 0, time.UTC),
				MaxClockOutTime: time.Date(2025, time.January, 21, 16, 0, 0, 0, time.UTC),
			},
			expectError: true,
		},
		{
			name: "Should return error when Update fails",
			setupMocks: func(repo *mockRepo.MockDepartementRepository, req *model.DepartementUpdateRequest) {
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
			},
			req: &model.DepartementUpdateRequest{
				ID:              uuid.New(),
				DepartementName: "Finance",
				MaxClockInTime:  time.Date(2025, time.January, 21, 8, 0, 0, 0, time.UTC),
				MaxClockOutTime: time.Date(2025, time.January, 21, 16, 0, 0, 0, time.UTC),
			},
			expectError: true,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			repo := mockRepo.NewMockDepartementRepository(ctrl)
			logger := logrus.New()
			service := service.NewDepartementService(repo, logger)

			w := httptest.NewRecorder()
			c, _ := gin.CreateTestContext(w)

			tc.setupMocks(repo, tc.req)

			err := service.UpdateDepartement(c, tc.req)

			if tc.expectError {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

func TestDeleteDepartement(t *testing.T) {
	tests := []struct {
		name        string
		id          int
		setupMocks  func(repo *mockRepo.MockDepartementRepository, id int)
		expectError bool
	}{
		{
			name: "Should delete department successfully",
			id:   1,
			setupMocks: func(repo *mockRepo.MockDepartementRepository, id int) {
				repo.EXPECT().
					Delete(gomock.Any(), id).
					Return(nil).
					Times(1)
			},
			expectError: false,
		},
		{
			name: "Should return error when Delete fails",
			id:   1,
			setupMocks: func(repo *mockRepo.MockDepartementRepository, id int) {
				repo.EXPECT().
					Delete(gomock.Any(), id).
					Return(errors.New("failed to delete departement")).
					Times(1)
			},
			expectError: true,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			repo := mockRepo.NewMockDepartementRepository(ctrl)
			logger := logrus.New()
			service := service.NewDepartementService(repo, logger)

			w := httptest.NewRecorder()
			c, _ := gin.CreateTestContext(w)

			tc.setupMocks(repo, tc.id)

			err := service.DeleteDepartement(c, tc.id)

			if tc.expectError {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}
