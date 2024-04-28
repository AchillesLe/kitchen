package service

import (
	"context"
	"fmt"

	"github.com/AchillesLe/kitchen/services/common/genproto/orders"
)

type OrderService struct {
	// rpc CreateOrder (CreateOrderRequest) returns (CreateOrderResponse) {}
}

var OrderDb = make([]*orders.Order, 0)

func NewOrderService() *OrderService {
	return &OrderService{}
}

func (h *OrderService) CreateOrder(ctx context.Context, order *orders.Order) error {
	OrderDb = append(OrderDb, order)
	fmt.Println("OrderDb => ", OrderDb)
	return nil
}

func (h *OrderService) GetOrders(ctx context.Context, customerID int32) []*orders.Order {
	return OrderDb
}
