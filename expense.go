package expensetracker

import "time"

type ExpenseCategory string

const (
	Groceries   ExpenseCategory = "Groceries"
	Leisure     ExpenseCategory = "Leisure"
	Electronics ExpenseCategory = "Electronics"
	Utilities   ExpenseCategory = "Utilities"
	Clothing    ExpenseCategory = "Clothing"
	Health      ExpenseCategory = "Health"
)

func (c ExpenseCategory) Validate() bool {
	switch c {
	case Groceries, Leisure, Electronics, Utilities, Clothing, Health:
		return true
	}

	return false
}

type Expense struct {
	ID          int             `json:"id" db:"id"`
	UserID      int             `json:"user_id" db:"user_id"`
	Total       float64         `json:"total" db:"total"`
	Description string          `json:"description" db:"description"`
	Category    ExpenseCategory `json:"category" db:"category"`
	CreatedAt   time.Time       `json:"created_at" db:"created_at"`
}
