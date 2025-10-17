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

func (service *PaymentService) CreatePayment(ctx context.Context, payment dtos.PaymentInput, userID int) (dtos.ConfirmResponse, error) {
	if err := utils.ValidatePaymentInput(payment); err != nil {
		return dtos.ConfirmResponse{}, err
	}

	isDuplicate, exiting, err := service.CheckDuplicatePayment(ctx, payment, userID)
	if err != nil {
		return dtos.ConfirmResponse{}, err
	}

	if isDuplicate {
		return dtos.ConfirmResponse{
			Payment: exiting,
			RequiresConfirmation: true,
			Message: "Similar payment found. Do you want to confirm creation?",
		}, nil
	}
	
	createdPayment, err := service.Repo.CreatePayment(ctx, payment, userID)
	if err != nil {
		return dtos.ConfirmResponse{}, err
	}

	pspResponse, err := utils.CreateStripePayment(ctx, int(payment.Amount), payment.Currency, "card")	// <= Just "card" for testing purposes.
	if err != nil {
		return dtos.ConfirmResponse{}, err
	}

	updatedPayment, err := service.Repo.UpdatePaymentWithPSP(ctx, createdPayment.ID, pspResponse.ID, pspResponse.Status)
	if err != nil {
		return dtos.ConfirmResponse{}, err
	}

	return dtos.ConfirmResponse{
		Payment: updatedPayment,
		RequiresConfirmation: false,
	}, nil
}

func (service *PaymentService) CheckDuplicatePayment(ctx context.Context, input dtos.PaymentInput, userID int) (bool, dtos.PaymentOutput, error) {
	similar := dtos.SimilarPayment {
		UserID: userID,
		Amount: input.Amount,
		PaymentMethodID: input.PaymentMethodID,
		Status: "pending",
	}

	existing, err := service.Repo.FindSimilarPayment(ctx, similar)
	if err != nil {
		return false, dtos.PaymentOutput{}, err
	}

	if existing.ID != 0 {
		return true, existing, nil
	}

	return false, dtos.PaymentOutput{}, nil
}

func (service *PaymentService) ForceCreatePayment(ctx context.Context, payment dtos.PaymentInput, userID int) (dtos.PaymentOutput, error) {
	if err := utils.ValidatePaymentInput(payment); err != nil {
		return dtos.PaymentOutput{}, err
	}

	createdPayment, err := service.Repo.CreatePayment(ctx, payment, userID)
	if err != nil {
		return dtos.PaymentOutput{}, err
	}

	pspResponse, err := utils.CreateStripePayment(ctx, int(payment.Amount), payment.Currency, "card")
	if err != nil {
		return dtos.PaymentOutput{}, err
	}

	return service.Repo.UpdatePaymentWithPSP(ctx, createdPayment.ID, pspResponse.ID, pspResponse.Status)
}

func (service *PaymentService) GetAllPayments(ctx context.Context, userID int) ([]dtos.PaymentOutput, error) {
	return service.Repo.GetAllPayments(ctx, userID)
}

func (service *PaymentService) GetPaymentByID(ctx context.Context, userID, paymentID int) (dtos.PaymentOutput, error) {
	return service.Repo.GetPaymentByID(ctx, userID, paymentID)
}

func (service *PaymentService) DeletePayment(ctx context.Context, userID, paymentID int) error {
	payment, err := service.Repo.GetPaymentByID(ctx, userID, paymentID)
	if err != nil {
		return err
	}

	err = utils.DeleteStripePayment(payment.PspID.String)
	if err != nil {
		return err
	}

	return service.Repo.DeletePayment(ctx, userID, paymentID)
}