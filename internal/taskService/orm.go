package taskService

import "gorm.io/gorm"

type Task struct {
	gorm.Model
	Task   string `json:"task_id"`
	IsDone bool   `json:"is_done"`
}
