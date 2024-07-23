package settlements

import (
	"context"

	"basic.web/internal/model"
)

type SettlementsUsecase interface {
	GetSettlements(ctx context.Context, params GetSettlementsParams) (*model.Settlements, error)
}
