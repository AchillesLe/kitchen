package types

import (
	"context"

	"github.com/AchillesLe/kitchen/services/common/genproto/orders"
)

type IOrderService interface {
	CreateOrder(ctx context.Context, req *orders.Order) error
	GetOrders(ctx context.Context, customerID int32) []*orders.Order
}
