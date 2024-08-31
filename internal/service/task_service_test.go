package service

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"github.com/kumazan/task-api/internal/model"
)

type MockTaskRepository struct {
	mock.Mock
}

func (m *MockTaskRepository) Create(task model.Task) (model.Task, error) {
	args := m.Called(task)
	return args.Get(0).(model.Task), args.Error(1)
}

func (m *MockTaskRepository) GetAll() ([]model.Task, error) {
	args := m.Called()
	return args.Get(0).([]model.Task), args.Error(1)
}

func (m *MockTaskRepository) GetByID(id string) (model.Task, error) {
	args := m.Called(id)
	return args.Get(0).(model.Task), args.Error(1)
}

func (m *MockTaskRepository) Update(task model.Task) error {
	args := m.Called(task)
	return args.Error(0)
}

func (m *MockTaskRepository) Delete(id string) error {
	args := m.Called(id)
	return args.Error(0)
}

func TestTaskService(t *testing.T) {
	mockRepo := new(MockTaskRepository)
	service := NewTaskService(mockRepo)

	t.Run("CreateTask", func(t *testing.T) {
		mockRepo.On("Create", mock.AnythingOfType("model.Task")).Return(model.Task{ID: "1", Name: "Test Task", Status: model.TaskStatusIncomplete}, nil)

		task, err := service.CreateTask("Test Task", model.TaskStatusIncomplete)
		assert.NoError(t, err)
		assert.Equal(t, "Test Task", task.Name)
		assert.Equal(t, model.TaskStatusIncomplete, task.Status)

		mockRepo.AssertExpectations(t)
	})

	t.Run("GetAllTasks", func(t *testing.T) {
		mockRepo.On("GetAll").Return([]model.Task{{ID: "1", Name: "Test Task", Status: 0}}, nil)

		tasks, err := service.GetAllTasks()
		assert.NoError(t, err)
		assert.Len(t, tasks, 1)
		assert.Equal(t, "Test Task", tasks[0].Name)

		mockRepo.AssertExpectations(t)
	})

	t.Run("UpdateTask", func(t *testing.T) {
		mockRepo.On("GetByID", "1").Return(model.Task{ID: "1", Name: "Test Task", Status: 0}, nil)
		mockRepo.On("Update", mock.AnythingOfType("model.Task")).Return(nil)

		err := service.UpdateTask("1", "Updated Task", 1)
		assert.NoError(t, err)

		mockRepo.AssertExpectations(t)
	})

	t.Run("DeleteTask", func(t *testing.T) {
		mockRepo.On("Delete", "1").Return(nil)

		err := service.DeleteTask("1")
		assert.NoError(t, err)

		mockRepo.AssertExpectations(t)
	})
}
