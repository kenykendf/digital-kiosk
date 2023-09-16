package service

import (
	"errors"
	"fmt"
	"time"

	"kenykendf/digital-kiosk/internal/app/model"
	"kenykendf/digital-kiosk/internal/app/schema"
	"kenykendf/digital-kiosk/internal/pkg/reason"

	log "github.com/sirupsen/logrus"
	"golang.org/x/crypto/bcrypt"
)

type TokenGenerator interface {
	CreateAcessToken(UserID int) (string, time.Time, error)
	CreateRefreshToken(UserID int) (string, time.Time, error)
}

type AuthRepository interface {
	Create(auth model.Auth) error
	Find(userID int, RefreshToken string) (model.Auth, error)
	GetByUserID(userID int) (model.Auth, error)
	Delete(userID int) error
}

type SessionService struct {
	userRepo   UserRepository
	authRepo   AuthRepository
	tokenMaker TokenGenerator
}

func NewSessionService(userRepo UserRepository, authRepo AuthRepository, tokenMaker TokenGenerator) *SessionService {
	return &SessionService{userRepo: userRepo, authRepo: authRepo, tokenMaker: tokenMaker}
}

func (ss *SessionService) Login(req *schema.LoginReq) (schema.LoginResp, error) {
	var resp schema.LoginResp

	existingUser, _ := ss.userRepo.GetByEmail(req.Email)
	if existingUser.ID == 0 {
		return resp, errors.New(reason.UserNotFound)
	}

	match := ss.VerifyPassword(existingUser.Password, req.Password)
	if !match {
		return resp, errors.New(reason.LoginFailed)
	}

	accessToken, _, err := ss.tokenMaker.CreateAcessToken(existingUser.ID)
	if err != nil {
		log.Error(fmt.Errorf("error SessionService - Access Token : %w", err))
		return resp, errors.New(reason.LoginFailed)
	}
	refreshToken, expireAt, err := ss.tokenMaker.CreateRefreshToken(existingUser.ID)
	if err != nil {
		log.Error(fmt.Errorf("error SessionService - Refresh Token : %w", err))
		return resp, errors.New(reason.LoginFailed)
	}

	resp.AccessToken = accessToken
	resp.RefreshToken = refreshToken

	err = ss.SaveToken(model.Auth{
		Token:    refreshToken,
		AuthType: "refresh_token",
		UserID:   existingUser.ID,
		Expiry:   expireAt,
	})
	if err != nil {
		return resp, errors.New(reason.LoginFailed)
	}

	return resp, nil
}

func (ss *SessionService) VerifyPassword(hashedPassword, password string) bool {

	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	return err == nil
}

func (ss *SessionService) SaveToken(data model.Auth) error {

	err := ss.authRepo.Create(data)
	if err != nil {
		return errors.New(reason.SaveToken)
	}
	return nil
}

func (ss *SessionService) RefreshToken(req *schema.RefreshTokenReq) (schema.RefreshTokenResp, error) {
	var resp schema.RefreshTokenResp

	existingUser, _ := ss.userRepo.GetByID(req.UserID)
	if existingUser.ID == 0 {
		return resp, errors.New(reason.UserNotFound)
	}

	find, _ := ss.authRepo.Find(existingUser.ID, req.RefreshToken)
	if find.ID == 0 {
		return resp, errors.New(reason.InvalidRefreshToken)
	}

	token, _, err := ss.tokenMaker.CreateAcessToken(existingUser.ID)
	if err != nil {
		return resp, errors.New(reason.CannotCreateAccessToken)
	}

	resp.AccessToken = token

	return resp, nil
}

func (ss *SessionService) Logout(req *schema.LogoutReq) error {

	err := ss.authRepo.Delete(req.UserID)

	if err != nil {
		log.Error(fmt.Errorf("error LoginService - Delete Session : %w", err))

	}

	return nil
}
