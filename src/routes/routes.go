package routes

import (
	"github.com/go-chi/chi/v5"
	"github.com/jhonnydsl/payment-API/src/controllers"
	"github.com/jhonnydsl/payment-API/src/utils/middleware"
)

func SetupRoutes(r chi.Router, userController *controllers.UserController, paymentController *controllers.PaymentController) {
	// User routes
	r.Post("/users", userController.CreateUser)
	r.Post("/login", userController.LoginUser)

	// Payment routes
	r.Group(func(protected chi.Router) {
		protected.Use(middleware.AuthMiddleware)
		protected.Post("/payments", paymentController.CreatePayment)
		protected.Get("/payments", paymentController.GetAllPayments)
		protected.Get("/payments/{paymentID}", paymentController.GetPaymentByID)
	})
}