package model

import "time"
 
type Budget struct {
	ID                uint64    `db:"id"`
	UserID            uint64    `db:"user_id"`
	TotalIncome       uint64    `db:"total_income"`
	NeedsPercentage   uint8     `db:"needs_percentage"`
	WantsPercentage   uint8     `db:"wants_percentage"`
	SavingsPercentage uint8     `db:"savings_percentage"`
	Year              uint16    `db:"year"`
	Month             uint8     `db:"month"`
	CreatedAt         time.Time `db:"created_at"`
	UpdatedAt         time.Time `db:"updated_at"`
}