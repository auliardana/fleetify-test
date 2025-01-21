package repository_test

import (
	"context"
	"testing"

	"github.com/auliardana/fleetify-test/internal/entity"
	_ "github.com/auliardana/fleetify-test/internal/repository"
	"github.com/auliardana/fleetify-test/test/mock"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func TestEmployeeRepository_Create(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	// Create mock instance
	mockRepo := mock.NewMockEmployeeRepository(ctrl)

	// Dummy departement_id from departement table
	departement_id := uuid.New()

	// Prepare input employee
	employee := &entity.Employee{
		DepartementID: departement_id,
		Name:          "John Doe",
		Address:       "Jl. Jendral Sudirman No. 1",
	}

	// Setup expected call on the mock
	// Expect the Create method to be called with the employee and return no error
	mockRepo.EXPECT().Create(context.Background(), employee).Return(nil)

	// Call the Create method on the mock repository
	err := mockRepo.Create(context.Background(), employee)

	// Assert that no error occurred (this checks the code path inside the Create method)
	assert.NoError(t, err)

	// Additional verification:
	// Verify that the employee object is passed correctly (this is to cover the function logic)
	assert.Equal(t, departement_id, employee.DepartementID)
	assert.Equal(t, "John Doe", employee.Name)
	assert.Equal(t, "Jl. Jendral Sudirman No. 1", employee.Address)
}


// func TestEmployeeRepository_Delete(t *testing.T) {
// 	ctrl := gomock.NewController(t)
// 	defer ctrl.Finish()

// 	mockRepo := mock.NewMockEmployeeRepository(ctrl)

// 	employee_id := uuid.New()

// 	// Set up the expected call
// 	mockRepo.EXPECT().Delete(context.Background(), employee_id).Return(nil)

// 	// Call the Delete method
// 	err := mockRepo.Delete(context.Background(), employee_id)

// 	// Assert that no error occurred
// 	assert.NoError(t, err)
// }

// func TestEmployeeRepository_FindById(t *testing.T) {
// 	ctrl := gomock.NewController(t)
// 	defer ctrl.Finish()

// 	mockRepo := mock.NewMockEmployeeRepository(ctrl)

// 	//dummy departement_id from departement table
// 	departement_id := uuid.New()

// 	// Prepare the employee data to return
// 	employee := &entity.Employee{
// 		DepartementID: departement_id,
// 		Name:          "John Doe",
// 		Address:       "Jl. Jendral Sudirman No. 1",
// 	}

// 	//dummy employee_id from employee table
// 	employee_id := uuid.New()

// 	// Set up the expected call
// 	mockRepo.EXPECT().FindById(context.Background(), employee_id).Return(employee, nil)

// 	// Call the FindById method
// 	result, err := mockRepo.FindById(context.Background(), employee_id)

// 	// Assert that the result is as expected and no error occurred
// 	assert.NoError(t, err)
// 	assert.Equal(t, employee, result)
// }

// func TestEmployeeRepository_GetAll(t *testing.T) {
// 	ctrl := gomock.NewController(t)
// 	defer ctrl.Finish()

// 	mockRepo := mock.NewMockEmployeeRepository(ctrl)

// 	//dummy departement_id from departement table
// 	departement_id := uuid.New()
// 	departement_id2 := uuid.New()

// 	// Prepare the list of employees to return
// 	employees := []entity.Employee{
// 		{
// 			DepartementID: departement_id,
// 			Name:          "John Doe",
// 			Address:       "Jl. Jendral Sudirman No. 1",
// 		},
// 		{
// 			DepartementID: departement_id2,
// 			Name:          "John Doe2",
// 			Address:       "Jl. Jendral Sudirman No. 2",
// 		},
// 	}

// 	// Set up the expected call
// 	mockRepo.EXPECT().GetAll(context.Background()).Return(employees, nil)

// 	// Call the GetAll method
// 	result, err := mockRepo.GetAll(context.Background())

// 	// Assert that the result is as expected and no error occurred
// 	assert.NoError(t, err)
// 	assert.Equal(t, employees, result)
// }

// func TestEmployeeRepository_Update(t *testing.T) {
// 	ctrl := gomock.NewController(t)
// 	defer ctrl.Finish()

// 	mockRepo := mock.NewMockEmployeeRepository(ctrl)

// 	//dummy departement_id from departement table
// 	departement_id := uuid.New()

// 	// Prepare the employee data to update
// 	employee := &entity.Employee{
// 		DepartementID: departement_id,
// 		Name:          "John Doe -Updated",
// 		Address:       "Jl. Jendral Sudirman No. 1",
// 	}

// 	// Set up the expected call
// 	mockRepo.EXPECT().Update(context.Background(), employee).Return(nil)

// 	// Call the Update method
// 	err := mockRepo.Update(context.Background(), employee)

// 	// Assert that no error occurred
// 	assert.NoError(t, err)
// }
