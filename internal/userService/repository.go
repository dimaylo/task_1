package userService

import "gorm.io/gorm"

type UserRepository interface {
	CreateUser(task User) (User, error)
	GetAllUsers() ([]User, error)
	UpdateUserByID(id uint, updated User) (User, error)
	DeleteUserByID(id uint) error
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{db: db}
}

func (u *userRepository) CreateUser(user User) (User, error) {
	err := u.db.Create(&user).Error
	return user, err
}

func (r *userRepository) GetAllUsers() ([]User, error) {
	var tasks []User
	err := r.db.Find(&tasks).Error
	return tasks, err
}

func (r *userRepository) UpdateUserByID(id uint, updated User) (User, error) {
	var existing User
	err := r.db.First(&existing, id).Error
	if err != nil {
		return existing, err
	}
	existing.Email = updated.Email
	existing.Password = updated.Password
	err = r.db.Save(&existing).Error
	return existing, err
}

func (r *userRepository) DeleteUserByID(id uint) error {
	var existing User
	err := r.db.First(&existing, id).Error
	if err != nil {
		return err
	}
	return r.db.Delete(&existing).Error
}
