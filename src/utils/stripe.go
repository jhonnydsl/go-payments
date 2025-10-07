package utils

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/jhonnydsl/payment-API/src/dtos"
)

func CreateStripePayment(amount int, currency, paymentMethod string) (*dtos.PSPPaymentResponse, error) {
	secretKey := os.Getenv("STRIPE_SECRET_KEY")
	url := "https://api.stripe.com/v1/payment_intents"

	data := fmt.Sprintf(
		"amount=%d&currency=%s&payment_method_types[]=%s",
		amount, currency, paymentMethod,
	)

	req, err := http.NewRequest("POST", url, strings.NewReader(data))
	if err != nil {
		return nil, err
	}

	req.Header.Add("Authorization", "Bearer " + secretKey)
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var result dtos.PSPPaymentResponse
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}
	
	return &result, nil
}