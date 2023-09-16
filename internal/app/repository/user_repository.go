package repository

import (
	"fmt"
	"time"

	"kenykendf/digital-kiosk/internal/app/model"

	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type UserRepository struct {
	DB *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{DB: db}
}

func (ur *UserRepository) Create(user model.User) error {

	if err := ur.DB.Select("email", "password", "fullname").Create(&user).Error; err != nil {
		log.Error(fmt.Errorf("error UserRepository - Create : %w", err))
		return err
	}

	return nil
}

func (ur *UserRepository) Browse() ([]model.User, error) {
	var User []model.User

	result := ur.DB.Find(&User)
	if result.Error != nil {
		log.Error(fmt.Errorf("error UserAddressRepository - Browse : %w", result.Error))
		return User, result.Error
	}

	return User, nil
}

func (ur *UserRepository) GetByEmail(email string) (model.User, error) {

	var user model.User

	result := ur.DB.Where("email = ?", email).First(&user)
	if result.Error != nil {
		log.Error(fmt.Errorf("error UserRepository - GetByEmail : %w", result.Error))
		return user, result.Error

	}
	return user, nil

}

func (ur *UserRepository) GetByID(id int) (model.User, error) {

	var user model.User

	result := ur.DB.Find(&user, id)
	if result != nil {
		log.Error(fmt.Errorf("error UserRepository - GetByID : %w", result.Error))
		return user, result.Error

	}
	return user, nil

}

func (ur *UserRepository) UpdateByID(id int, user model.User) error {

	data := &model.User{
		Email:     user.Email,
		Fullname:  user.Fullname,
		UpdatedAt: time.Now(),
	}

	if err := ur.DB.Where("id = ?", id).Updates(&data).Error; err != nil {
		log.Error(fmt.Errorf("error UserRepository - Update : %w", err))
		return err
	}

	return nil

}

func (ur *UserRepository) DeleteByID(id int) error {
	var user model.User

	if err := ur.DB.Where("id = ?", id).Delete(&user).Error; err != nil {
		log.Error(fmt.Errorf("error UserRepository - Delete : %w", err))
		return err
	}

	return nil
}
