package repository

import (
	"context"

	"github.com/jhonnydsl/payment-API/src/dtos"
)

type PaymentsRepository struct{}

func (r *PaymentsRepository) CreatePayment(ctx context.Context, payment dtos.PaymentInput, userID int) (dtos.PaymentOutput, error) {
	query := `
	INSERT INTO payments (user_id, amount, currency, payment_method_id, status)
	VALUES ($1, $2, $3, $4, $5)
	RETURNING id, user_id, amount, currency, payment_method_id, status, created_at, updated_at;
	`
	var createdPayment dtos.PaymentOutput


	err := DB.QueryRowContext(ctx, query, userID, payment.Amount, payment.Currency, payment.PaymentMethodID, "pending").Scan(
		&createdPayment.ID,
		&createdPayment.UserID,
		&createdPayment.Amount,
		&createdPayment.Currency,
		&createdPayment.PaymentMethodID,
		&createdPayment.Status,
		&createdPayment.CreatedAt,
		&createdPayment.UpdatedAt,
	)
	if err != nil {
		return dtos.PaymentOutput{}, err
	}

	return createdPayment, nil
}