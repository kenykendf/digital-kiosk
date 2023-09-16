package service

import (
	"errors"

	"kenykendf/digital-kiosk/internal/app/model"
	"kenykendf/digital-kiosk/internal/app/schema"
	"kenykendf/digital-kiosk/internal/pkg/reason"
)

type UserRepository interface {
	Create(user model.User) error
	Browse() ([]model.User, error)
	GetByEmail(email string) (model.User, error)
	GetByID(id int) (model.User, error)
	UpdateByID(id int, user model.User) error
	DeleteByID(id int) error
}

type UserService struct {
	repo UserRepository
}

func NewUserService(repo UserRepository) *UserService {
	return &UserService{repo: repo}
}

func (us *UserService) BrowseAll() ([]schema.GetUsersResp, error) {
	var resp []schema.GetUsersResp

	users, err := us.repo.Browse()
	if err != nil {
		return nil, errors.New(reason.UserCannotBrowse)
	}

	for _, value := range users {
		var respData schema.GetUsersResp
		respData.ID = value.ID
		respData.Fullname = value.Fullname
		respData.Email = value.Email
		resp = append(resp, respData)
	}

	return resp, nil
}

func (us *UserService) GetByID(id int) (schema.GetUsersResp, error) {
	var resp schema.GetUsersResp

	user, err := us.repo.GetByID(id)
	if err != nil {
		return resp, errors.New(reason.UserCannotGetDetail)
	}
	if user.ID == 0 {
		return resp, errors.New(reason.UserNotFound)
	}

	resp.ID = user.ID
	resp.Fullname = user.Fullname
	resp.Email = user.Email

	return resp, nil
}

func (us *UserService) DeleteByID(id int) error {

	check, err := us.repo.GetByID(id)
	if check.ID == 0 || err != nil {
		return errors.New(reason.UserNotFound)
	}

	err = us.repo.DeleteByID(id)
	if err != nil {
		return errors.New(reason.UserCannotDelete)
	}

	return nil
}
