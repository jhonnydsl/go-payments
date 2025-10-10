package controllers

import (
	"net/http"

	"github.com/jhonnydsl/payment-API/src/dtos"
	"github.com/jhonnydsl/payment-API/src/services"
	"github.com/jhonnydsl/payment-API/src/utils"
	"github.com/jhonnydsl/payment-API/src/utils/middleware"
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

	if !utils.DecodeJSON(w, r, &paymentInput) {
		return
	}

	userID, ok := r.Context().Value(middleware.UserIDKey).(int)
	if !ok {
		http.Error(w, "unauthorized", http.StatusUnauthorized)
		return
	}

	newPayment, err := controller.Service.CreatePayment(ctx, paymentInput, userID)
	if err != nil {
		http.Error(w, "error creating payment", http.StatusInternalServerError)
		return
	}

	utils.SendJSON(w, newPayment, http.StatusCreated)
}

func (controller *PaymentController) GetAllPayments(w http.ResponseWriter, r *http.Request) {
	if !utils.ValidateMethod(w, r, "GET") {
		return
	}

	ctx, cancel := utils.NewDBContext()
	defer cancel()

	userID, ok := r.Context().Value(middleware.UserIDKey).(int)
	if !ok {
		http.Error(w, "unauthorized", http.StatusUnauthorized)
		return
	}

	paymentsList, err := controller.Service.GetAllPayments(ctx, userID)
	if err != nil {
		http.Error(w, "error listing payments", http.StatusInternalServerError)
		return
	}

	utils.SendJSON(w, paymentsList, http.StatusOK)
}