package model

import "time"

// ExpenseType merepresentasikan kolom enum('pengeluaran','pemasukan') di MySQL.
type ExpenseType string

const (
	ExpenseTypePengeluaran ExpenseType = "pengeluaran"
	ExpenseTypePemasukan   ExpenseType = "pemasukan"
)

type Expense struct {
	ID          uint64      `db:"id"`
	UserID      uint64      `db:"user_id"`
	CategoryID  uint64      `db:"category_id"`
	Amount      uint64      `db:"amount"`
	Description *string     `db:"description"` // pointer karena kolom DEFAULT NULL (boleh kosong)
	ExpenseDate time.Time   `db:"expense_date"`
	Type        ExpenseType `db:"type"`
	CreatedAt   time.Time   `db:"created_at"`
	UpdatedAt   time.Time   `db:"updated_at"`
}
