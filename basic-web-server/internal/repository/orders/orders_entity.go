package orders

import (
	"time"

	"basic.web/internal/repository/customers"
	"basic.web/internal/repository/product"
)

type Order struct {
	OrderID     int       `gorm:"column:orderid"`
	OrderDate   time.Time `gorm:"column:orderdate"`
	CustomerID  int       `gorm:"column:customerid"`
	NetAmount   float64   `gorm:"column:netamount"`
	Tax         float64   `gorm:"column:tax"`
	TotalAmount float64   `gorm:"column:totalamount"`
}

type OrderLine struct {
	OrderLineID int       `gorm:"column:orderlineid"`
	OrderID     int       `gorm:"column:orderid"`
	ProdID      int       `gorm:"column:prod_id"`
	Quantity    int       `gorm:"column:quantity"`
	OrderDate   time.Time `gorm:"column:orderdate"`
}

type OrderWithLine struct {
	Order
	OrderLine
	product.Product
	product.Category
	customers.Customer
}

func (o Order) TableName() string {
	return "orders"
}
