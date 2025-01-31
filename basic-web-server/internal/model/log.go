package model

import "time"

type Log struct {
	Order
	DateRetrieval time.Time
}
