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

func (r *PaymentsRepository) UpdatePaymentWithPSP(ctx context.Context, paymentID int, pspID, status string) (dtos.PaymentOutput, error) {
	query := `
	UPDATE payments
	SET psp_id = $1, status = $2, updated_at = CURRENT_TIMESTAMP
	WHERE id = $3
	RETURNING id, user_id, amount, currency, payment_method_id, status, created_at, updated_at;
	`

	var updatedPayment dtos.PaymentOutput

	err := DB.QueryRowContext(ctx, query, pspID, status, paymentID).Scan(
		&updatedPayment.ID,
		&updatedPayment.UserID,
		&updatedPayment.Amount,
		&updatedPayment.Currency,
		&updatedPayment.PaymentMethodID,
		&updatedPayment.Status,
		&updatedPayment.CreatedAt,
		&updatedPayment.UpdatedAt,
	)
	if err != nil {
		return dtos.PaymentOutput{}, err
	}

	return updatedPayment, nil
}

func (r *PaymentsRepository) GetAllPayments(ctx context.Context, userID int) ([]dtos.PaymentOutput, error) {
	query := `SELECT id, amount, currency, payment_method_id, status, created_at, updated_at FROM payments WHERE user_id = $1`
	var list []dtos.PaymentOutput

	rows, err := DB.Query(query, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var output dtos.PaymentOutput

		err = rows.Scan(
			&output.ID, 
			&output.Amount, 
			&output.Currency, 
			&output.PaymentMethodID, 
			&output.Status, 
			&output.CreatedAt, 
			&output.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}

		list = append(list, output)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return list, nil
}