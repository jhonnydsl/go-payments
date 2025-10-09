package utils

import (
	"context"
	"os"

	"github.com/jhonnydsl/payment-API/src/dtos"
	"github.com/stripe/stripe-go/v76"
	"github.com/stripe/stripe-go/v76/paymentintent"
)

func CreateStripePayment(ctx context.Context, amount int, currency, paymentMethod string) (*dtos.PSPPaymentResponse, error) {
	stripe.Key = os.Getenv("STRIPE_SECRET_KEY")
	
	params := &stripe.PaymentIntentParams{
		Amount: stripe.Int64(int64(amount * 100)),
		Currency: stripe.String(currency),
		PaymentMethodTypes: stripe.StringSlice([]string{"card"}),
		PaymentMethod: stripe.String("pm_card_visa"),	// <= just testing
		Confirm: stripe.Bool(true),
	}

	intent, err := paymentintent.New(params)
	if err != nil {
		return nil, err
	}

	response := &dtos.PSPPaymentResponse{
		ID: intent.ID,
		Status: string(intent.Status),
		Amount: int(intent.Amount / 100),
	}

	return response, nil
}