package apiservice

import (
	"errors"

	"github.com/Muhfikri12/project-app-inventaris-golang-fikri/model"
	apirepository "github.com/Muhfikri12/project-app-inventaris-golang-fikri/repository/api_repository"
)

type AuthService struct {
	AuthRepo *apirepository.AuthRepo
}

func NewAuthService(repo *apirepository.AuthRepo) *AuthService {
	return &AuthService{AuthRepo: repo}
}

func (as *AuthService) LoginUser(user *model.User) (error) {
	
	err := as.AuthRepo.LoginUser(user)
	if err != nil {
		return err
	}

	return nil
}

func (as *AuthService) CheckToken(token string) error {
	valid, err := as.AuthRepo.CheckToken(token)
	if err != nil {
		return err
	}

	if !valid {
		return errors.New("token tidak valid atau sudah kadaluarsa")
	}

	return nil
}
