package settlements

import "time"

type Settlements struct {
	ID          int       `gorm:"primaryKey;column:s_id"`
	Code        string    `gorm:"column:s_code"`
	TotalAmount float64   `gorm:"column:s_total_amount"`
	CreatedAt   time.Time `gorm:"column:s_created_at"`
}

func (t Settlements) TableName() string {
	return "settlements"
}
