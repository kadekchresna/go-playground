package model

import "time"

type Settlements struct {
	ID          int
	Code        string
	TotalAmount float64
	CreatedAt   time.Time
}
