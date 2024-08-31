package settlements

import (
	"context"

	"gorm.io/gorm"
)

type settlementsRepo struct {
	DB *gorm.DB
}

func NewSettlementRepo(DB *gorm.DB) SettlementsRepo {
	return &settlementsRepo{
		DB: DB,
	}
}

func (r *settlementsRepo) GetSettlements(ctx context.Context, params GetSettlementsParams) ([]Settlements, error) {
	res := []Settlements{}

	if err := r.DB.Find(&res).Error; err != nil {
		return nil, err
	}

	return res, nil
}
