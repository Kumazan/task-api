package repository

import (
	"testing"

	"github.com/kumazan/task-api/internal/model"
	"github.com/stretchr/testify/assert"
)

func TestInMemoryTaskRepository(t *testing.T) {
	repo := NewInMemoryTaskRepository()

	t.Run("Create and GetByID", func(t *testing.T) {
		task := model.Task{ID: "1", Name: "Test Task", Status: 0}
		createdTask, err := repo.Create(task)
		assert.NoError(t, err)
		assert.Equal(t, task, createdTask)

		fetchedTask, err := repo.GetByID("1")
		assert.NoError(t, err)
		assert.Equal(t, task, fetchedTask)
	})

	t.Run("GetAll", func(t *testing.T) {
		tasks, err := repo.GetAll()
		assert.NoError(t, err)
		assert.Len(t, tasks, 1)
		assert.Equal(t, "Test Task", tasks[0].Name)
	})

	t.Run("Update", func(t *testing.T) {
		updatedTask := model.Task{ID: "1", Name: "Updated Task", Status: 1}
		err := repo.Update(updatedTask)
		assert.NoError(t, err)

		fetchedTask, err := repo.GetByID("1")
		assert.NoError(t, err)
		assert.Equal(t, updatedTask, fetchedTask)
	})

	t.Run("Delete", func(t *testing.T) {
		err := repo.Delete("1")
		assert.NoError(t, err)

		_, err = repo.GetByID("1")
		assert.Error(t, err)

		tasks, err := repo.GetAll()
		assert.NoError(t, err)
		assert.Len(t, tasks, 0)
	})
}