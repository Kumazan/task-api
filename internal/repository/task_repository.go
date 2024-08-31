package repository

import (
	"errors"
	"sync"

	"github.com/kumazan/task-api/internal/model"
)

type TaskRepository interface {
	Create(task model.Task) (model.Task, error)
	GetAll() ([]model.Task, error)
	GetByID(id string) (model.Task, error)
	Update(task model.Task) error
	Delete(id string) error
}

type InMemoryTaskRepository struct {
	tasks map[string]model.Task
	mutex sync.RWMutex
}

func NewInMemoryTaskRepository() *InMemoryTaskRepository {
	return &InMemoryTaskRepository{
		tasks: make(map[string]model.Task),
	}
}

func (r *InMemoryTaskRepository) Create(task model.Task) (model.Task, error) {
	r.mutex.Lock()
	defer r.mutex.Unlock()

	r.tasks[task.ID] = task
	return task, nil
}

func (r *InMemoryTaskRepository) GetAll() ([]model.Task, error) {
	r.mutex.RLock()
	defer r.mutex.RUnlock()

	tasks := make([]model.Task, 0, len(r.tasks))
	for _, task := range r.tasks {
		tasks = append(tasks, task)
	}
	return tasks, nil
}

func (r *InMemoryTaskRepository) GetByID(id string) (model.Task, error) {
	r.mutex.RLock()
	defer r.mutex.RUnlock()

	task, ok := r.tasks[id]
	if !ok {
		return model.Task{}, errors.New("task not found")
	}
	return task, nil
}

func (r *InMemoryTaskRepository) Update(task model.Task) error {
	r.mutex.Lock()
	defer r.mutex.Unlock()

	if _, ok := r.tasks[task.ID]; !ok {
		return errors.New("task not found")
	}
	r.tasks[task.ID] = task
	return nil
}

func (r *InMemoryTaskRepository) Delete(id string) error {
	r.mutex.Lock()
	defer r.mutex.Unlock()

	if _, ok := r.tasks[id]; !ok {
		return errors.New("task not found")
	}
	delete(r.tasks, id)
	return nil
}
