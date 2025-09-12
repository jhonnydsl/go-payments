package utils

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/jhonnydsl/payment-API/src/dtos"
)

func ValidatePaymentInput(payment dtos.PaymentInput) error {
	if payment.Amount < 1 {
		return fmt.Errorf("amount must be greater than 0")
	} 

	if payment.Currency == "" {
		return fmt.Errorf("currency is required")
	}

	if payment.PaymentMethodID == 0 {
		return fmt.Errorf("payment method is required")
	}

	return nil
}

func ValidateMethod(w http.ResponseWriter, r *http.Request, method string) bool {
	if r.Method != method {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return false
	}

	return true
}

func DecodeJSON(w http.ResponseWriter, r *http.Request, destination interface{}) bool {
	if err := json.NewDecoder(r.Body).Decode(destination); err != nil {
		http.Error(w, "error decoding JSON", http.StatusBadRequest)
		return false
	}
	
	return true
}