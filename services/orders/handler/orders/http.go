package handler

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/AchillesLe/kitchen/services/common/genproto/orders"
	"github.com/AchillesLe/kitchen/services/orders/types"
)

type OrdersHttpHandler struct {
	ordersService types.IOrderService
}

func NewOrdersHttpHandler(ordersService types.IOrderService) *OrdersHttpHandler {
	return &OrdersHttpHandler{ordersService: ordersService}
}

func (h *OrdersHttpHandler) RegisterRoute(router *http.ServeMux) {
	router.HandleFunc("/orders", h.CreateOrder)
}

func (h *OrdersHttpHandler) CreateOrder(w http.ResponseWriter, r *http.Request) {
	var req orders.CreateOrderRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	order := &orders.Order{
		OrderID:    42,
		CustomerID: req.GetCustomerID(),
		ProductID:  req.GetProductID(),
		Quantity:   req.GetQuantity(),
	}

	fmt.Println("order ==>", order)

	err = h.ordersService.CreateOrder(r.Context(), order)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	res := &orders.CreateOrderResponse{Status: "success"}

	json.NewEncoder(w).Encode(res)
}
