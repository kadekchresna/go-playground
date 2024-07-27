package customers

type Customer struct {
	CustomerID           int    `gorm:"column:customerid"`
	FirstName            string `gorm:"column:firstname"`
	LastName             string `gorm:"column:lastname"`
	Address1             string `gorm:"column:address1"`
	Address2             string `gorm:"column:address2"`
	City                 string `gorm:"column:city"`
	State                string `gorm:"column:state"`
	Zip                  int    `gorm:"column:zip"`
	Country              string `gorm:"column:country"`
	Region               string `gorm:"column:region"`
	Email                string `gorm:"column:email"`
	Phone                string `gorm:"column:phone"`
	CreditCardType       int    `gorm:"column:creditcardtype"`
	CreditCard           string `gorm:"column:creditcard"`
	CreditCardExpiration string `gorm:"column:creditcardexpiration"`
	Username             string `gorm:"column:username"`
	Password             string `gorm:"column:password"`
	Age                  int    `gorm:"column:age"`
	Income               int    `gorm:"column:income"`
	Gender               string `gorm:"column:gender"`
}

func (c Customer) TableName() string {
	return `customers`
}
