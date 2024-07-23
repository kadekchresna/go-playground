package settlements

import (
	"context"
	"time"

	"basic.web/internal/model"
	"basic.web/internal/repository/settlements"
)

type settlementsUsecase struct {
	SettlementsRepo settlements.SettlementsRepo
}

func NewSettlementsUsecase(SettlementsRepo settlements.SettlementsRepo) SettlementsUsecase {
	return &settlementsUsecase{
		SettlementsRepo: SettlementsRepo,
	}
}

func (u *settlementsUsecase) GetSettlements(ctx context.Context, params GetSettlementsParams) (*model.Settlements, error) {
	settlements, err := u.SettlementsRepo.GetSettlements(ctx, settlements.GetSettlementsParams{})
	if err != nil {
		return nil, err
	}

	return &model.Settlements{ID: settlements.ID, Code: settlements.Code, TotalAmount: settlements.TotalAmount, CreatedAt: time.Now()}, nil
}
