package services

import (
	"context"

	"github.com/jhonnydsl/payment-API/src/dtos"
	"github.com/jhonnydsl/payment-API/src/repository"
	"github.com/jhonnydsl/payment-API/src/utils"
)

type PaymentService struct {
	Repo *repository.PaymentsRepository
}

func (service *PaymentService) CreatePayment(ctx context.Context, payment dtos.PaymentInput, userID int) (dtos.PaymentOutput, error) {
	if err := utils.ValidatePaymentInput(payment); err != nil {
		return dtos.PaymentOutput{}, err
	}
	
	createdPayment, err := service.Repo.CreatePayment(ctx, payment, userID)
	if err != nil {
		return dtos.PaymentOutput{}, err
	}

	pspResponse, err := utils.CreateStripePayment(ctx, int(payment.Amount), payment.Currency, "card")	// <= Just "card" for testing purposes.
	if err != nil {
		return dtos.PaymentOutput{}, err
	}

	updatedPayment, err := service.Repo.UpdatePaymentWithPSP(ctx, createdPayment.ID, pspResponse.ID, pspResponse.Status)
	if err != nil {
		return dtos.PaymentOutput{}, err
	}

	return updatedPayment, nil
}

func (service *PaymentService) GetAllPayments(ctx context.Context, userID int) ([]dtos.PaymentOutput, error) {
	return service.Repo.GetAllPayments(ctx, userID)
}