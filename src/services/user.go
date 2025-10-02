package services

import (
	"context"
	"errors"
	"fmt"

	"github.com/jhonnydsl/payment-API/src/dtos"
	"github.com/jhonnydsl/payment-API/src/repository"
	"github.com/jhonnydsl/payment-API/src/utils"
)

type UserService struct {
	Repo *repository.UserRepository
}

func (service *UserService) CreateUser (ctx context.Context, user dtos.UserInput) (dtos.UserOutput, error) {
	if err := utils.ValidateUserInput(user); err != nil {
		return dtos.UserOutput{}, err
	}

	return service.Repo.CreateUser(ctx, user)
}

func (service *UserService) LoginUser(ctx context.Context, login dtos.UserLogin) (string, error) {
	userLogin, password, err := service.Repo.GetUserByEmail(ctx, login.Email)
	if err != nil {
		return "", errors.New("email or password invalid")
	}

	if err := utils.CheckPassword(password, login.Password); err != nil {
		return "", errors.New("email or password invalid")
	}

	tokenStr, err := utils.GenerateJWT(userLogin.ID, userLogin.Email)
	if err != nil {
		return "", fmt.Errorf("error generating token: %w", err)
	}

	return tokenStr, nil
}