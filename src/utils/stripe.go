package utils

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/jhonnydsl/payment-API/src/dtos"
)

func CreateStripePayment(ctx context.Context, amount int, currency, paymentMethod string) (*dtos.PSPPaymentResponse, error) {
	secretKey := os.Getenv("STRIPE_SECRET_KEY")
	url := "https://api.stripe.com/v1/payment_intents"

	// Builds the data string in the format required by Stripe, inserting values dynamically
	data := fmt.Sprintf(
		"amount=%d&currency=%s&payment_method_types[]=%s",
		amount, currency, paymentMethod,
	)

	req, err := http.NewRequestWithContext(ctx, "POST", url, strings.NewReader(data))
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