package handler

import (
	"context"

	"github.com/AchillesLe/kitchen/services/common/genproto/orders"
	"github.com/AchillesLe/kitchen/services/orders/types"
	"google.golang.org/grpc"
)

type OrderGrpcHandler struct {
	// service injection
	ordersService types.IOrderService
	orders.UnimplementedOrderServiceServer
}

func NewGrpcOrderService(grpc *grpc.Server, ordersService types.IOrderService) {
	grpcHandler := &OrderGrpcHandler{ordersService: ordersService}
	// register order service server
	orders.RegisterOrderServiceServer(grpc, grpcHandler)
}

func (h *OrderGrpcHandler) CreateOrder(ctx context.Context, req *orders.CreateOrderRequest) (*orders.CreateOrderResponse, error) {
	order := &orders.Order{
		OrderID:    42,
		CustomerID: req.GetCustomerID(),
		ProductID:  req.GetProductID(),
		Quantity:   req.GetQuantity(),
	}

	err := h.ordersService.CreateOrder(ctx, order)

	if err != nil {
		return nil, err
	}

	res := &orders.CreateOrderResponse{
		Status: "success",
	}

	return res, nil
}

func (h *OrderGrpcHandler) GetOrders(ctx context.Context, req *orders.GetOrdersRequest) (*orders.GetOrdersResponse, error) {
	CustomerID := req.GetCustomerID()
	orderDb := h.ordersService.GetOrders(ctx, CustomerID)
	res := &orders.GetOrdersResponse{
		Orders: orderDb,
	}
	return res, nil
}
