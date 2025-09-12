package controllers

import (
	"net/http"

	"github.com/jhonnydsl/payment-API/src/dtos"
	"github.com/jhonnydsl/payment-API/src/services"
	"github.com/jhonnydsl/payment-API/src/utils"
)

type PaymentController struct {
	Service *services.PaymentService
}

func (controller *PaymentController) CreatePayment(w http.ResponseWriter, r *http.Request) {
	if !utils.ValidateMethod(w, r, "POST") {
		return
	}

	ctx, cancel := utils.NewDBContext()
	defer cancel()

	var paymentInput dtos.PaymentInput
	userID := 1		// <= Temporary userID for testing.

	if !utils.DecodeJSON(w, r, &paymentInput) {
		return
	}

	newPayment, err := controller.Service.CreatePayment(ctx, paymentInput, userID)
	if err != nil {
		http.Error(w, "error creating payment", http.StatusInternalServerError)
		return
	}

	utils.SendJSON(w, newPayment, http.StatusCreated)
}