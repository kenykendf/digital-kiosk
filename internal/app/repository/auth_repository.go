package repository

import (
	"fmt"

	"kenykendf/digital-kiosk/internal/app/model"

	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type AuthRepository struct {
	DB *gorm.DB
}

func NewAuthRepository(db *gorm.DB) *AuthRepository {
	return &AuthRepository{DB: db}
}

func (ar *AuthRepository) Create(auth model.Auth) error {

	if err := ar.DB.Select("token", "auth_type", "user_id", "expires_at").Create(&auth).Error; err != nil {
		log.Error(fmt.Errorf("error authRepository - Create : %w", err))
		return err
	}

	return nil
}

func (ar *AuthRepository) Find(userID int, RefreshToken string) (model.Auth, error) {
	var (
		auth model.Auth
		err  error
	)

	result := ar.DB.Where("user_id = ? AND token = ?", userID, RefreshToken).Find(&auth)
	if result != nil {
		return auth, nil
	} else {
		log.Error(fmt.Errorf("error authRepository - Find : %w", err))
		return auth, err

	}
}
func (ar *AuthRepository) GetByUserID(userID int) (model.Auth, error) {
	var (
		auth model.Auth
		err  error
	)

	result := ar.DB.Where("user_id = ?", userID).Find(&auth)
	if result != nil {
		return auth, nil
	} else {
		log.Error(fmt.Errorf("error authRepository - GetByUserID : %w", err))
		return auth, err

	}
}
func (ar *AuthRepository) Delete(userID int) error {
	var auth model.Auth

	if err := ar.DB.Where("user_id = ?", userID).Delete(&auth).Error; err != nil {
		log.Error(fmt.Errorf("error authRepository - Delete : %w", err))
		return err
	}

	return nil
}
