package service

import (
	"fyndr.com/api/src/config"
	"gorm.io/gorm"
)

type AuthService interface {
	Login(code string) (string, error)
}

type AuthServiceImpl struct {
	env *config.AppEnv
	db  *gorm.DB
}

func NewAuthService(env *config.AppEnv, db *gorm.DB) AuthService {
	return &AuthServiceImpl{
		env: env,
		db:  db,
	}
}

func (a AuthServiceImpl) Login(code string) (string, error) {
	return "", nil
}
