package service

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type TokenMaker struct {
	AccessTokenKey       string
	RefreshTokenKey      string
	AccessTokenDuration  time.Duration
	RefreshTokenDuration time.Duration
}

type Claims struct {
	UserID string
	jwt.RegisteredClaims
}

func NewTokenMaker(
	accessTokenKey string,
	refreshTokenKey string,
	accessTokenDuration time.Duration,
	refreshTokenDuration time.Duration,
) *TokenMaker {
	return &TokenMaker{
		AccessTokenKey:       accessTokenKey,
		RefreshTokenKey:      refreshTokenKey,
		AccessTokenDuration:  accessTokenDuration,
		RefreshTokenDuration: refreshTokenDuration,
	}
}

func (maker *TokenMaker) CreateAcessToken(UserID int) (string, time.Time, error) {

	expirationTime := time.Now().Add(maker.AccessTokenDuration)
	key := maker.AccessTokenKey

	claims := &Claims{
		UserID: fmt.Sprintf("%d", UserID),
		RegisteredClaims: jwt.RegisteredClaims{
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			ExpiresAt: jwt.NewNumericDate(expirationTime),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(key))
	if err != nil {
		return "", expirationTime, err
	}

	return tokenString, expirationTime, nil

}

func (maker *TokenMaker) CreateRefreshToken(UserID int) (string, time.Time, error) {

	expirationTime := time.Now().Add(maker.RefreshTokenDuration)
	key := maker.RefreshTokenKey

	claims := &Claims{
		UserID: fmt.Sprintf("%d", UserID),
		RegisteredClaims: jwt.RegisteredClaims{
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			ExpiresAt: jwt.NewNumericDate(expirationTime),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(key))
	if err != nil {
		return "", expirationTime, err
	}

	return tokenString, expirationTime, nil

}

func (maker *TokenMaker) ValidateAccessToken(tokenString string) (string, error) {
	sub, err := maker.validate(tokenString, maker.AccessTokenKey)
	return sub, err
}
func (maker *TokenMaker) ValidateRefreshToken(tokenString string) (string, error) {
	sub, err := maker.validate(tokenString, maker.RefreshTokenKey)
	return sub, err
}

func (maker *TokenMaker) validate(tokenString, tokenkey string) (string, error) {
	token, err := jwt.ParseWithClaims(tokenString, &Claims{},
		func(token *jwt.Token) (interface{}, error) {
			return []byte(tokenkey), nil
		},
	)
	if err != nil {
		return "", err
	}
	claims, ok := token.Claims.(*Claims)
	if ok && token.Valid {
		sub := claims.UserID
		return sub, nil
	}

	return "", err
}
