package routes

import (
	"go-mysql-backend/handlers"

	"github.com/gorilla/mux"
)

func RegisterRoutes(router *mux.Router) {
	// Register product routes
	registerProductRoutes(router)

	// Register user routes
	registerUserRoutes(router)
	// Login endpoint
	router.HandleFunc("/login", handlers.LoginHandler).Methods("POST")
}

func registerProductRoutes(router *mux.Router) {
	productRouter := router.PathPrefix("/products").Subrouter()
	productRouter.HandleFunc("", handlers.GetProducts).Methods("GET")
	productRouter.HandleFunc("/{id}", handlers.GetProduct).Methods("GET")
	productRouter.HandleFunc("", handlers.CreateProduct).Methods("POST")
	productRouter.HandleFunc("/{id}", handlers.UpdateProduct).Methods("PUT")
	productRouter.HandleFunc("/{id}", handlers.DeleteProduct).Methods("DELETE")
}

func registerUserRoutes(router *mux.Router) {
	userRouter := router.PathPrefix("/users").Subrouter()
	userRouter.HandleFunc("", handlers.GetUsers).Methods("GET")
	userRouter.HandleFunc("/{id}", handlers.GetUser).Methods("GET")
	userRouter.HandleFunc("", handlers.CreateUser).Methods("POST")
	userRouter.HandleFunc("/{id}", handlers.UpdateUser).Methods("PUT")
	userRouter.HandleFunc("/{id}", handlers.DeleteUser).Methods("DELETE")
}
