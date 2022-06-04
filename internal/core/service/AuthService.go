package service

import (
	"go-hospital-server/internal/core/entity/models"
	"go-hospital-server/internal/core/entity/request"
	"go-hospital-server/internal/core/entity/response"
	"go-hospital-server/internal/core/repository"
	"go-hospital-server/internal/utils"
	"go-hospital-server/internal/utils/errors"
)

type AuthService struct {
	repo repository.AuthRepository
}

func NewAuthService(repo repository.AuthRepository) *AuthService {
	return &AuthService{repo: repo}
}

func (srv AuthService) Login(login request.LoginRequest) (res *response.User, err error) {
	var user models.User
	var checkPassword bool
	user, err  = srv.repo.Login(login.Email)

	err = errors.CheckError(user, err)

	if err == nil {
		checkPassword = utils.ComparePassword(login.Password, user.Password)
		if checkPassword {
			res, _ = utils.TypeConverter[response.User](&user)
		} else {
			err = errors.New(417, "Wrong Password")
		}
	}

	return
}

func (srv AuthService) RefreshToken(str models.Token) (token models.Token, err error) {
	token, err = srv.repo.RefreshToken(str)
	return
}