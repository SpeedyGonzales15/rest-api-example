package handler

import (
	"rest-api-example/internal/service"

	"github.com/gorilla/mux"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) InitRoutes() *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/users/create", h.createUser).Methods("POST")
	router.HandleFunc("/users/list", h.getUsersList).Methods("GET")
	router.HandleFunc("/users/{id}", h.getUserById).Methods("GET")
	router.HandleFunc("/users/{id}", h.updateUser).Methods("PUT")
	router.HandleFunc("/users/{id}", h.deleteUser).Methods("DELETE")

	router.HandleFunc("/products/list", h.getProductsList).Methods("GET")
	router.HandleFunc("/products/{id}", h.getProductById).Methods("GET")

	router.HandleFunc("/order/create", h.createOrder).Methods("POST")
	router.HandleFunc("/order/{id}", h.getOrdersList).Methods("GET")

	return router
}
