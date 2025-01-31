package orders

import (
	"context"

	"gorm.io/gorm"
)

type orderRepo struct {
	DB *gorm.DB
}

func NewOrderRepo(DB *gorm.DB) OrderRepo {
	return &orderRepo{
		DB: DB,
	}
}

func (r *orderRepo) GetAllOrders(ctx context.Context, params GetAllOrdersParams) ([]OrderWithLine, error) {

	res := []OrderWithLine{}

	if err := r.DB.Raw(`
	select * from orders o 
	inner join customers c on o.customerid = c.customerid 
	inner join orderlines o2 on o.orderid = o2.orderid 
	inner join products p on o2.prod_id = p.prod_id 
	inner join categories c2 on c2.category = p.category 
	`).Scan(&res).Error; err != nil {
		return nil, err
	}

	return res, nil
}

func (r *orderRepo) GetAllLogsByID(ctx context.Context, params GetAllOrdersParams) ([]OrderWithLine, error) {

	res := []OrderWithLine{}

	if err := r.DB.Raw(`
	select * from orders o 
		inner join customers c on o.customerid = c.customerid 
		inner join orderlines o2 on o.orderid = o2.orderid 
		inner join products p on o2.prod_id = p.prod_id 
		inner join categories c2 on c2.category = p.category 
		left join inventory i on i.prod_id  = p.prod_id 
		left join cust_hist ch on ch.customerid = c.customerid 
	`).Scan(&res).Error; err != nil {
		return nil, err
	}

	return res, nil
}

func (r *orderRepo) GetAllLogsByIDChunk(ctx context.Context, params GetAllOrdersParams) ([]OrderWithLine, error) {

	res := []OrderWithLine{}

	if err := r.DB.Raw(`
	select * from orders o 
		inner join customers c on o.customerid = c.customerid 
		inner join orderlines o2 on o.orderid = o2.orderid 
		inner join products p on o2.prod_id = p.prod_id 
		inner join categories c2 on c2.category = p.category 
		left join inventory i on i.prod_id  = p.prod_id 
		left join cust_hist ch on ch.customerid = c.customerid 
		limit ? offset ?
	`, params.Limit, params.Offset).Scan(&res).Error; err != nil {
		return nil, err
	}

	return res, nil
}

func (r *orderRepo) GetAllLogPaginate(ctx context.Context, params GetAllOrdersParams) ([]OrderWithLine, error) {

	res := []OrderWithLine{}

	if err := r.DB.Raw(`
	select * from orders o 
		inner join customers c on o.customerid = c.customerid 
		inner join orderlines o2 on o.orderid = o2.orderid 
		inner join products p on o2.prod_id = p.prod_id 
		inner join categories c2 on c2.category = p.category 
		left join inventory i on i.prod_id  = p.prod_id 
		left join cust_hist ch on ch.customerid = c.customerid 
		limit ? offset ?
	`, params.Limit, params.Offset).Scan(&res).Error; err != nil {
		return nil, err
	}

	return res, nil
}
