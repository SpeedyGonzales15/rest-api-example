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
	router.HandleFunc("/users/{id}", h.updateUser).Methods("PATCH")
	router.HandleFunc("/users/{id}", h.deleteUser).Methods("DELETE")

	return router
}
