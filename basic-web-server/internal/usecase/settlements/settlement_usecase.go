package settlements

import (
	"context"
	"time"

	"basic.web/internal/model"
	"basic.web/internal/repository/orders"
	"basic.web/internal/repository/settlements"
)

type settlementsUsecase struct {
	SettlementsRepo settlements.SettlementsRepo
	OrderRepo       orders.OrderRepo
}

func NewSettlementsUsecase(
	SettlementsRepo settlements.SettlementsRepo,
	OrderRepo orders.OrderRepo,
) SettlementsUsecase {
	return &settlementsUsecase{
		SettlementsRepo: SettlementsRepo,
		OrderRepo:       OrderRepo,
	}
}

func (u *settlementsUsecase) GetSettlements(ctx context.Context, params GetSettlementsParams) ([]model.Settlements, error) {
	settlements, err := u.SettlementsRepo.GetSettlements(ctx, settlements.GetSettlementsParams{})
	if err != nil {
		return nil, err
	}

	res := []model.Settlements{}
	for _, s := range settlements {
		res = append(res, model.Settlements{ID: s.ID, Code: s.Code, TotalAmount: s.TotalAmount, CreatedAt: time.Now()})
	}

	return res, nil
}

func (u *settlementsUsecase) GetSettlementsOne(ctx context.Context, params GetSettlementsParams) (*model.Settlements, error) {
	settlements, err := u.SettlementsRepo.GetSettlements(ctx, settlements.GetSettlementsParams{})
	if err != nil {
		return nil, err
	}

	return &model.Settlements{ID: settlements[0].ID, Code: settlements[0].Code, TotalAmount: settlements[0].TotalAmount, CreatedAt: time.Now()}, nil
}

func (u *settlementsUsecase) GetOrders(ctx context.Context, params GetAllOrdersParams) ([]model.Order, error) {
	res := []model.Order{}
	orders, err := u.OrderRepo.GetAllOrders(ctx, orders.GetAllOrdersParams{})
	if err != nil {
		return nil, err
	}

	for _, o := range orders {
		res = append(res, model.Order{
			OrderID:              o.Order.OrderID,
			OrderDate:            o.Order.OrderDate,
			CustomerID:           o.Order.CustomerID,
			NetAmount:            o.NetAmount,
			Tax:                  o.Tax,
			TotalAmount:          o.TotalAmount,
			OrderLineID:          o.OrderLineID,
			ProdID:               o.Product.ProdID,
			Quantity:             o.Quantity,
			Category:             o.Product.Category,
			Title:                o.Title,
			Actor:                o.Actor,
			Price:                o.Price,
			Special:              o.Special,
			CommonProdID:         o.CommonProdID,
			CategoryName:         o.CategoryName,
			FirstName:            o.FirstName,
			LastName:             o.LastName,
			Address1:             o.Address1,
			Address2:             o.Address2,
			City:                 o.City,
			State:                o.State,
			Zip:                  o.Zip,
			Country:              o.Country,
			Region:               o.Region,
			Email:                o.Email,
			Phone:                o.Phone,
			CreditCardType:       o.CreditCardType,
			CreditCard:           o.CreditCard,
			CreditCardExpiration: o.CreditCardExpiration,
			Username:             o.Username,
			Password:             o.Password,
			Age:                  o.Age,
			Income:               o.Income,
			Gender:               o.Gender,
		})
	}

	return res, nil
}
