package settlements

import "context"

type SettlementsRepo interface {
	GetSettlements(ctx context.Context, params GetSettlementsParams) ([]Settlements, error)
}
