package orders

import "context"

type OrderRepo interface {
	GetAllOrders(ctx context.Context, params GetAllOrdersParams) ([]OrderWithLine, error)
}
