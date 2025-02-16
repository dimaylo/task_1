package userService

import "gorm.io/gorm"

type UserRepository interface {
	CreateUser(user User) (User, error)
	GetAllUsers() ([]User, error)
	GetUserByID(id uint) (User, error)
	UpdateUserByID(id uint, updated User) (User, error)
	DeleteUserByID(id uint) error
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{db: db}
}

func (r *userRepository) CreateUser(user User) (User, error) {
	err := r.db.Create(&user).Error
	return user, err
}

func (r *userRepository) GetAllUsers() ([]User, error) {
	var users []User
	err := r.db.Find(&users).Error
	return users, err
}

func (r *userRepository) GetUserByID(id uint) (User, error) {
	var user User
	err := r.db.First(&user, id).Error
	return user, err
}

func (r *userRepository) UpdateUserByID(id uint, updated User) (User, error) {
	var user User
	err := r.db.First(&user, id).Error
	if err != nil {
		return user, err
	}
	user.Email = updated.Email
	user.Password = updated.Password
	err = r.db.Save(&user).Error
	return user, err
}

func (r *userRepository) DeleteUserByID(id uint) error {
	var user User
	err := r.db.First(&user, id).Error
	if err != nil {
		return err
	}
	return r.db.Delete(&user).Error
}
