package service

import (
	"errors"

	"kenykendf/digital-kiosk/internal/app/model"
	"kenykendf/digital-kiosk/internal/app/schema"
	"kenykendf/digital-kiosk/internal/pkg/reason"

	"golang.org/x/crypto/bcrypt"
)

type RegistrationService struct {
	repo UserRepository
}

func NewRegistrationService(repo UserRepository) *RegistrationService {
	return &RegistrationService{repo: repo}
}

func (rs *RegistrationService) Register(req *schema.RegisterReq) error {

	existingUser, _ := rs.repo.GetByEmail(req.Email)
	if existingUser.ID > 0 {
		return errors.New(reason.UserAlreadyExist)
	}

	password, _ := rs.hashPassword(req.Password)

	var insertData model.User
	insertData.Fullname = req.Fullname
	insertData.Password = password
	insertData.Email = req.Email

	err := rs.repo.Create(insertData)
	if err != nil {
		return errors.New(reason.RegisterFailed)
	}

	return nil
}

func (rs *RegistrationService) hashPassword(password string) (string, error) {
	bytePassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

	if err != nil {
		return "", err
	}

	return string(bytePassword), nil

}
