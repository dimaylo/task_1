package taskService

import "gorm.io/gorm"

type TaskRepository interface {
	CreateTask(task Task) (Task, error)
	GetAllTasks() ([]Task, error)
	UpdateTaskByID(id uint, updated Task) (Task, error)
	DeleteTaskByID(id uint) error
}

type taskRepository struct {
	db *gorm.DB
}

func NewTaskRepository(db *gorm.DB) TaskRepository {
	return &taskRepository{db: db}
}

func (r *taskRepository) CreateTask(task Task) (Task, error) {
	err := r.db.Create(&task).Error
	return task, err
}

func (r *taskRepository) GetAllTasks() ([]Task, error) {
	var tasks []Task
	err := r.db.Find(&tasks).Error
	return tasks, err
}

func (r *taskRepository) UpdateTaskByID(id uint, updated Task) (Task, error) {
	var existing Task
	err := r.db.First(&existing, id).Error
	if err != nil {
		return existing, err
	}
	existing.Task = updated.Task
	existing.IsDone = updated.IsDone
	err = r.db.Save(&existing).Error
	return existing, err
}

func (r *taskRepository) DeleteTaskByID(id uint) error {
	var existing Task
	err := r.db.First(&existing, id).Error
	if err != nil {
		return err
	}
	return r.db.Delete(&existing).Error
}
