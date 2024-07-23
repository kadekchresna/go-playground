package settlements

import "time"

type GetSettlementsParams struct {
}

type GetSettlementsResponse struct {
	ID          int       `json:"id"`
	Code        string    `json:"code"`
	TotalAmount float64   `json:"total_amount"`
	CreatedAt   time.Time `json:"created_at"`
}
