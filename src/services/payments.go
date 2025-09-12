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
	
	return service.Repo.CreatePayment(ctx, payment, userID)
}