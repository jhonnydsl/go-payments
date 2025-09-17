package controllers

import (
	"net/http"

	"github.com/jhonnydsl/payment-API/src/dtos"
	"github.com/jhonnydsl/payment-API/src/services"
	"github.com/jhonnydsl/payment-API/src/utils"
)

type UserController struct {
	Service *services.UserService
}

func (controller *UserController) CreateUser(w http.ResponseWriter, r *http.Request) {
	if !utils.ValidateMethod(w, r, "POST") {
		return
	}

	ctx, cancel := utils.NewDBContext()
	defer cancel()

	var userInput dtos.UserInput

	if !utils.DecodeJSON(w, r, &userInput) {
		return
	}

	newUser, err := controller.Service.CreateUser(ctx, userInput)
	if err != nil {
		http.Error(w, "error creating user", http.StatusInternalServerError)
		return
	}

	utils.SendJSON(w, newUser, http.StatusCreated)
}