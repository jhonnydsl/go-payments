package services

import (
	"context"

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