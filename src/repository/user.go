package repository

import (
	"context"

	"github.com/jhonnydsl/payment-API/src/dtos"
	"github.com/jhonnydsl/payment-API/src/utils"
)

type UserRepository struct{}

func (r *UserRepository) CreateUser(ctx context.Context, user dtos.UserInput) (dtos.UserOutput, error) {
	query := `
	INSERT INTO users (name, email, password)
	VALUES ($1, $2, $3)
	RETURNING id, name, email, created_at;
	`
	var createdUser dtos.UserOutput

	hashedPassword, err := utils.HashPassword(user.Password)
	if err != nil {
		return dtos.UserOutput{}, err
	}

	err = DB.QueryRow(query, user.Name, user.Email, hashedPassword).Scan(
		&createdUser.ID,
		&createdUser.Name,
		&createdUser.Email,
		&createdUser.CreatedAt,
	)

	return createdUser, err
}

func (r *UserRepository) GetUserByEmail(ctx context.Context, email string) (dtos.User, string, error) {
	query := `SELECT id, name, email, password, created_at FROM users WHERE email = $1`

	var user dtos.User
	var hashedPassword string

	err := DB.QueryRow(query, email).Scan(
		&user.ID,
		&user.Name,
		&user.Email,
		&hashedPassword,
		&user.CreatedAt,
	)
	if err != nil {
		return dtos.User{}, "", err
	}

	return user, hashedPassword, nil
}