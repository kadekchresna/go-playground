package model

import "time"

type Order struct {
	OrderID              int
	OrderDate            time.Time
	CustomerID           int
	NetAmount            float64
	Tax                  float64
	TotalAmount          float64
	OrderLineID          int
	ProdID               int
	Quantity             int
	Category             int
	Title                string
	Actor                string
	Price                float64
	Special              int
	CommonProdID         int
	CategoryName         string
	FirstName            string
	LastName             string
	Address1             string
	Address2             string
	City                 string
	State                string
	Zip                  int
	Country              string
	Region               string
	Email                string
	Phone                string
	CreditCardType       int
	CreditCard           string
	CreditCardExpiration string
	Username             string
	Password             string
	Age                  int
	Income               int
	Gender               string
}
