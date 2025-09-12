package repository

import (
	"context"

	"github.com/jhonnydsl/payment-API/src/dtos"
)

type PaymentsRepository struct{}

func (r *PaymentsRepository) CreatePayment(ctx context.Context, payment dtos.PaymentInput, userID int) (dtos.PaymentOutput, error) {
	query := `
	INSERT INTO payments (user_id, amount, currency, payment_method)
	VALUES ($1, $2, $3, $4)
	RETURNING id, user_id, amount, currency, payment_method, status, created_at, updated_at;
	`
	var createdPayment dtos.PaymentOutput

	err := DB.QueryRowContext(ctx, query, userID, payment.Amount, payment.Currency, payment.PaymentMethod).Scan(
		&createdPayment.ID,
		&createdPayment.UserID,
		&createdPayment.Amount,
		&createdPayment.Currency,
		&createdPayment.PaymentMethod,
		&createdPayment.Status,
		&createdPayment.CreatedAt,
		&createdPayment.UpdatedAt,
	)
	if err != nil {
		return dtos.PaymentOutput{}, err
	}

	return createdPayment, nil
}