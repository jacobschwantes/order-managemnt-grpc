package handler

import (
	"encoding/json"
	"net/http"

	"github.com/jacobschwantes/order-managemnt-grpc/services/common/genproto/orders"
	"github.com/jacobschwantes/order-managemnt-grpc/services/orders/types"
)

type OrdersHttpHandler struct {
	ordersService types.OrderService
}

func NewOrdersHttpHandler(ordersService types.OrderService) *OrdersHttpHandler {
	handler := &OrdersHttpHandler{ordersService: ordersService}

	return handler
}

func (h *OrdersHttpHandler) RegisterRouter(router *http.ServeMux) {
	router.HandleFunc("POST /orders", h.CreateOrder)
}

func (h *OrdersHttpHandler) CreateOrder(w http.ResponseWriter, r *http.Request) {
	var req orders.CreateOrderRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	order := &orders.Order{
		OrderID:    5,
		CustomerID: req.CustomerID,
		ProductID:  req.ProductID,
		Quantity:   req.Quantity,
	}

	err = h.ordersService.CreateOrder(r.Context(), order)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	res := &orders.CreateOrderResponse{
		Status: "success",
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(res)
}

func (h *OrdersHttpHandler) GetOrders(w http.ResponseWriter, r *http.Request) {
	orders := h.ordersService.GetOrders(r.Context())

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(orders)
}
