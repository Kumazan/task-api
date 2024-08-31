package service

import (
	"errors"

	"github.com/google/uuid"

	"github.com/kumazan/task-api/internal/model"
	"github.com/kumazan/task-api/internal/repository"
)

type TaskService struct {
	repo repository.TaskRepository
}

func NewTaskService(repo repository.TaskRepository) *TaskService {
	return &TaskService{repo: repo}
}

func (s *TaskService) CreateTask(name string, status model.TaskStatus) (model.Task, error) {
	if name == "" {
		return model.Task{}, errors.New("task name is required")
	}
	if status != model.TaskStatusIncomplete && status != model.TaskStatusCompleted {
		return model.Task{}, errors.New("invalid status value")
	}

	task := model.Task{
		ID:     uuid.New().String(),
		Name:   name,
		Status: status,
	}

	return s.repo.Create(task)
}

func (s *TaskService) GetAllTasks() ([]model.Task, error) {
	return s.repo.GetAll()
}

func (s *TaskService) UpdateTask(id string, name string, status model.TaskStatus) error {
	task, err := s.repo.GetByID(id)
	if err != nil {
		return err
	}

	if name != "" {
		task.Name = name
	}
	if status == model.TaskStatusIncomplete || status == model.TaskStatusCompleted {
		task.Status = status
	}

	return s.repo.Update(task)
}

func (s *TaskService) DeleteTask(id string) error {
	return s.repo.Delete(id)
}
