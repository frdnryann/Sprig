package model

import "time"
 
type Category struct {
	ID        uint64    `db:"id"`
	UserID    uint64    `db:"user_id"`
	Name      string    `db:"name"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
}