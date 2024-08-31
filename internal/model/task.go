package model

type TaskStatus int

const (
	TaskStatusIncomplete TaskStatus = 0
	TaskStatusCompleted  TaskStatus = 1
)

type Task struct {
	ID     string     `json:"id"`
	Name   string     `json:"name"`
	Status TaskStatus `json:"status"`
}