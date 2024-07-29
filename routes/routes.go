package routes

import (
	"database/sql"
	"go-mysql-backend/handlers"

	"github.com/gorilla/mux"
)

func RegisterRoutes(router *mux.Router, db *sql.DB) {
	// Register product routes
	registerProductRoutes(router, db)

	// Register user routes
	registerUserRoutes(router, db)

	// Register order detail routes
	registerOrderDetailRoutes(router, db)

	// Login endpoint
	router.HandleFunc("/login", handlers.LoginHandler(db)).Methods("POST")
}

func registerProductRoutes(router *mux.Router, db *sql.DB) {
	productRouter := router.PathPrefix("/products").Subrouter()
	productRouter.HandleFunc("", handlers.GetProducts(db)).Methods("GET")
	productRouter.HandleFunc("/{id}", handlers.GetProduct(db)).Methods("GET")
	productRouter.HandleFunc("", handlers.CreateProduct(db)).Methods("POST")
	productRouter.HandleFunc("/{id}", handlers.UpdateProduct(db)).Methods("PUT")
	productRouter.HandleFunc("/{id}", handlers.DeleteProduct(db)).Methods("DELETE")
}

func registerUserRoutes(router *mux.Router, db *sql.DB) {
	userRouter := router.PathPrefix("/users").Subrouter()
	userRouter.HandleFunc("", handlers.GetUsers(db)).Methods("GET")
	userRouter.HandleFunc("/{id:[0-9]+}", handlers.GetUser(db)).Methods("GET")
	userRouter.HandleFunc("", handlers.CreateUser(db)).Methods("POST")
	userRouter.HandleFunc("/{id:[0-9]+}", handlers.UpdateUser(db)).Methods("PUT")
	userRouter.HandleFunc("/{id:[0-9]+}", handlers.DeleteUser(db)).Methods("DELETE")
}

func registerOrderDetailRoutes(router *mux.Router, db *sql.DB) {
	orderDetailRouter := router.PathPrefix("/order_details").Subrouter()
	orderDetailRouter.HandleFunc("", handlers.GetOrderDetails(db)).Methods("GET")
	orderDetailRouter.HandleFunc("/{id}", handlers.GetOrderDetail(db)).Methods("GET")
	orderDetailRouter.HandleFunc("", handlers.CreateOrderDetail(db)).Methods("POST")
	orderDetailRouter.HandleFunc("/{id}", handlers.UpdateOrderDetail(db)).Methods("PUT")
	orderDetailRouter.HandleFunc("/{id}", handlers.DeleteOrderDetail(db)).Methods("DELETE")
}
