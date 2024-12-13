package settlements

import (
	"context"

	"basic.web/internal/model"
)

type SettlementsUsecase interface {
	GetSettlements(ctx context.Context, params GetSettlementsParams) ([]model.Settlements, error)
	GetOrders(ctx context.Context, params GetAllOrdersParams) ([]model.Order, error)
	GetLogs(ctx context.Context, params GetAllOrdersParams) ([]model.Log, error)
	GetLogsChan(ctx context.Context, params GetAllOrdersParams) (chan model.Log, error)
	GetLogsPaginate(ctx context.Context, params GetAllOrdersParams) ([]model.Log, error)
}
