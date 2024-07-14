package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"rest-api-example/internal/models"
	"strconv"

	"github.com/gorilla/mux"
)

func (h *Handler) createOrder(w http.ResponseWriter, r *http.Request) {
	var order models.Order
	err := json.NewDecoder(r.Body).Decode(&order)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// check if there are enough products
	for _, item := range order.Products {
		product, err := h.services.ProductServiceList.GetProductsById(item.ProductId)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		if product.Quantity < item.QuantityInOrder {
			http.Error(w, "Not enough products", http.StatusBadRequest)
			return
		}
	}

	// create order
	id, err := h.services.OrderServiceList.CreateOrder(order)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	res := fmt.Sprintf("Order %d was created", id)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(res)
}

func (h *Handler) getOrdersList(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userId, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	list, err := h.services.OrderServiceList.GetAllOrders(userId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if len(list) == 0 {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(list)
}
