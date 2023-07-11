package models

import (
	"time"
)

type BetStatus uint
type OrderType uint

const (
	BetStatusInit           BetStatus = 1
	BetStatusPlaced         BetStatus = 2
	BetStatusFulfilled      BetStatus = 3
	BetStatusExpired        BetStatus = 4
	BetStatusrequestFulfill BetStatus = 5
)

type Bet struct {
	BetID           uint64    `gorm:"primaryKey;" json:"bet_id"`
	Trader          string    `json:"trader"`
	Deadline        time.Time `json:"deadline"`
	LiquidationTime time.Time `json:"liquidation_time"`
	IsShort         bool      `json:"is_short"`
	Amount          uint64    `json:"amount"`
	Status          BetStatus `json:"status"`
	InitPrice       uint64    `json:"price"`
}
