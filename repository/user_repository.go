package repository

import (
	"seleksi-javan/model"

	"gorm.io/gorm"
)

type (
	UserRepository interface {
		CreateUser(user model.User) error
		RestoreAndUpdateUser(username string, newUser model.User) error
		FindUserByUsername(username string) (model.User, error)
		FindUserByEmail(email string) (model.User, error)
		FindUserByID(userId uint) (model.User, error)
		GetAllUser() ([]model.User, error)
		ChangePassword(userId uint, newPassword string) error
		UpdateUser(userId uint, updateUser model.User) error
		DeleteUser(userId uint) error
	}

	userRepository struct {
		db *gorm.DB
	}
)

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{
		db: db,
	}
}

func (ur *userRepository) CreateUser(user model.User) error {
	err := ur.db.Create(&user).Error

	return err
}

func (ur *userRepository) RestoreAndUpdateUser(username string, newUser model.User) error {
	err := ur.db.Unscoped().Model(&model.User{}).Where("username = ?", username).
		Updates(map[string]interface{}{
			"deleted_at": nil,
			"email":      newUser.Email,
			"password":   newUser.Password,
		}).Error

	return err
}

func (ur *userRepository) FindUserByUsername(username string) (model.User, error) {
	person := model.User{}
	err := ur.db.Where("username = ?", username).First(&person).Error

	return person, err
}

func (ur *userRepository) FindUserByEmail(email string) (model.User, error) {
	person := model.User{}
	err := ur.db.Where("email = ?", email).First(&person).Error

	return person, err
}

func (ur *userRepository) FindUserByID(userId uint) (model.User, error) {
	person := model.User{}
	err := ur.db.Preload("Tasks").Where("id = ?", userId).First(&person).Error

	return person, err
}

func (ur *userRepository) GetAllUser() ([]model.User, error) {
	var users []model.User
	err := ur.db.Preload("Tasks").Find(&users).Error

	return users, err
}

func (ur *userRepository) ChangePassword(userId uint, newPassword string) error {
	var x model.User

	err := ur.db.Where("id = ?", userId).First(&x).Error
	if err != nil {
		return err
	}

	err = ur.db.Model(&x).Update("password", newPassword).Error
	return err
}

func (ur *userRepository) UpdateUser(userId uint, updateUser model.User) error {
	err := ur.db.Model(&model.User{}).Where("id = ?", userId).Updates(updateUser).Error

	return err
}

func (ur *userRepository) DeleteUser(userId uint) error {
	var user model.User

	result := ur.db.First(&user, userId)
	if result.Error != nil {
		return result.Error
	}

	if err := ur.db.Delete(&user).Error; err != nil {
		return err
	}

	return nil
}
