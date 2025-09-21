package services

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/jhonnydsl/payment-API/src/dtos"
	"github.com/jhonnydsl/payment-API/src/repository"
	"github.com/jhonnydsl/payment-API/src/utils"
	"golang.org/x/crypto/bcrypt"
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
		return "", fmt.Errorf("email or password invalid")
	}

	err = bcrypt.CompareHashAndPassword([]byte(password), []byte(login.Password))
	if err != nil {
		return "", fmt.Errorf("email or password invalid")
	}

	claims := jwt.MapClaims{
		"user_id": userLogin.ID,
		"email": userLogin.Email,
		"exp": time.Now().Add(time.Hour *24).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenStr, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))
	if err != nil {
		return "", fmt.Errorf("error generating token")
	}

	return tokenStr, nil
}