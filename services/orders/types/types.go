package types

import (
	"context"

	"github.com/jacobschwantes/order-managemnt-grpc/services/common/genproto/orders"
)

type OrderService interface {
	CreateOrder(context.Context, *orders.Order) error
}
