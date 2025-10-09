package dtos

import "time"

type PaymentInput struct {
	Amount        float64 `json:"amount"`
	Currency      string  `json:"currency"`
	PaymentMethodID int  `json:"payment_method_id"`
}

type PaymentOutput struct {
	ID            int     `json:"id"`
	UserID        int     `json:"user_id"`
	Amount        float64 `json:"amount"`
	Currency      string  `json:"currency"`
	PaymentMethodID int  `json:"payment_method_id"`
	Status        string  `json:"status"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt	  time.Time `json:"updated_at"`
}

type PSPPaymentRequest struct {
	Amount int `json:"amount"`
	Currency string `json:"currency"`
}

type PSPPaymentResponse struct {
	ID string `json:"id"`
	Status string `json:"status"`
	Amount int `json:"amount"`
}