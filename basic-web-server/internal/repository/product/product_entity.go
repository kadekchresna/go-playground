package product

type Product struct {
	ProdID       int     `gorm:"column:prod_id"`
	Category     int     `gorm:"column:category"`
	Title        string  `gorm:"column:title"`
	Actor        string  `gorm:"column:actor"`
	Price        float64 `gorm:"column:price"`
	Special      int     `gorm:"column:special"`
	CommonProdID int     `gorm:"column:common_prod_id"`
}

type Category struct {
	Category     int    `gorm:"column:category"`
	CategoryName string `gorm:"column:categoryname"`
}

func (p Product) TableName() string {
	return `products`
}
