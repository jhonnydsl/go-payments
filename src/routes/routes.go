package routes

import (
	"net/http"

	"github.com/jhonnydsl/payment-API/src/controllers"
)

func SetupRoutes(userController *controllers.UserController) {
	http.HandleFunc("/users", userController.CreateUser)
}