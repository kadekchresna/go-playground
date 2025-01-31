package orders

import "context"

type OrderRepo interface {
	GetAllOrders(ctx context.Context, params GetAllOrdersParams) ([]OrderWithLine, error)
	GetAllLogsByID(ctx context.Context, params GetAllOrdersParams) ([]OrderWithLine, error)
	GetAllLogsByIDChunk(ctx context.Context, params GetAllOrdersParams) ([]OrderWithLine, error)
	GetAllLogPaginate(ctx context.Context, params GetAllOrdersParams) ([]OrderWithLine, error)
}
