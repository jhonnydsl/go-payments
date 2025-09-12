package dtos

import "time"

type PaymentInput struct {
	Amount        float64 `json:"amount" binding:"required"`
	Currency      string  `json:"currency" binding:"required"`
	PaymentMethodID int  `json:"payment_method_id" binding:"required"`
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