package repository

import (
	"context"
	"database/sql"
	"errors"

	"github.com/jhonnydsl/payment-API/src/dtos"
)

type PaymentsRepository struct{}

func (r *PaymentsRepository) CreatePayment(ctx context.Context, payment dtos.PaymentInput, userID int) (dtos.PaymentOutput, error) {
	query := `
	INSERT INTO payments (user_id, amount, currency, payment_method_id, status)
	VALUES ($1, $2, $3, $4, $5)
	RETURNING id, psp_id, user_id, amount, currency, payment_method_id, status, created_at, updated_at;
	`
	var createdPayment dtos.PaymentOutput


	err := DB.QueryRowContext(ctx, query, userID, payment.Amount, payment.Currency, payment.PaymentMethodID, "pending").Scan(
		&createdPayment.ID,
		&createdPayment.PspID,
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
	RETURNING id, psp_id, user_id, amount, currency, payment_method_id, status, created_at, updated_at;
	`

	var updatedPayment dtos.PaymentOutput

	err := DB.QueryRowContext(ctx, query, pspID, status, paymentID).Scan(
		&updatedPayment.ID,
		&updatedPayment.PspID,
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
	query := `SELECT id, COALESCE(psp_id, ''), user_id, amount, currency, payment_method_id, status, created_at, updated_at FROM payments WHERE user_id = $1`
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
			&output.PspID,
			&output.UserID,
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

func (r *PaymentsRepository) GetPaymentByID(ctx context.Context, userID, paymentID int) (dtos.PaymentOutput, error) {
	query := `
	SELECT id, COALESCE(psp_id, ''), user_id, amount, currency, payment_method_id, status, created_at, updated_at 
	FROM payments 
	WHERE user_id = $1 AND id = $2
	`
	var payment dtos.PaymentOutput

	err := DB.QueryRow(query, userID, paymentID).Scan(
		&payment.ID,
		&payment.PspID,
		&payment.UserID,
		&payment.Amount,
		&payment.Currency,
		&payment.PaymentMethodID,
		&payment.Status,
		&payment.CreatedAt,
		&payment.UpdatedAt,
	)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {		// <= checking error if payment does not exist
			return dtos.PaymentOutput{}, nil
		}
		return dtos.PaymentOutput{}, err
	}


	return payment, nil
}

func (r *PaymentsRepository) DeletePayment(ctx context.Context, userID, paymentID int) error {
	query := `
	DELETE FROM payments WHERE id = $1 AND user_id = $2
	`

	res, err := DB.Exec(query, paymentID, userID)
	if err != nil {
		return err
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return errors.New("no payments found to delete")
	}

	return nil
}

func (r *PaymentsRepository) FindSimilarPayment(ctx context.Context, similar dtos.SimilarPayment) (dtos.PaymentOutput, error) {
	query := `
	SELECT id, COALESCE(psp_id, ''), user_id, amount, currency, payment_method_id, status, created_at, updated_at
	FROM payments
	WHERE user_id = $1 AND amount = $2 AND payment_method_id = $3 AND status = $4
	LIMIT 1
	`
	var payment dtos.PaymentOutput

	err := DB.QueryRowContext(ctx, query, similar.UserID, similar.Amount, similar.PaymentMethodID, similar.Status).Scan(
		&payment.ID,
		&payment.PspID,
		&payment.UserID,
		&payment.Amount,
		&payment.Currency,
		&payment.PaymentMethodID,
		&payment.Status,
		&payment.CreatedAt,
		&payment.UpdatedAt,
	)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return dtos.PaymentOutput{}, nil
		}
		return dtos.PaymentOutput{}, err
	}

	return payment, nil
}