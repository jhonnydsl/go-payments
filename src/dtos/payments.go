package dtos

import (
	"database/sql"
	"time"
)

type PaymentInput struct {
	Amount        	float64 `json:"amount"`
	Currency      	string  `json:"currency"`
	PaymentMethodID int  `json:"payment_method_id"`
}

type PaymentOutput struct {
	ID            	int     `json:"id"`
	UserID        	int     `json:"user_id"`
	PspID		  	sql.NullString     `json:"psp_id"`
	Amount        	float64 `json:"amount"`
	Currency      	string  `json:"currency"`
	PaymentMethodID int  `json:"payment_method_id"`
	Status        	string  `json:"status"`
	CreatedAt     	time.Time `json:"created_at"`
	UpdatedAt	  	time.Time `json:"updated_at"`
}

type PSPPaymentRequest struct {
	Amount 		int `json:"amount"`
	Currency 	string `json:"currency"`
}

type PSPPaymentResponse struct {
	ID 		string `json:"id"`
	Status 	string `json:"status"`
	Amount 	int `json:"amount"`
}

type ConfirmCreatePayment struct {
	Amount 			float64 `json:"amount"`
	Currency 		string `json:"currency"`
	PaymentMethodID int `json:"payment_method_id"`
	Confirm 		bool `json:"confirm,omitempty"`
}

type SimilarPayment struct {
	UserID        	int     `json:"user_id"`
	Amount        	float64 `json:"amount"`
	PaymentMethodID int  `json:"payment_method_id"`
	Status        	string  `json:"status"`
}

type ConfirmResponse struct {
	Payment 				PaymentOutput `json:"payment"`
	RequiresConfirmation 	bool `json:"requires_confirmation"`
	Message 				string `json:"message,omitempty"`
}