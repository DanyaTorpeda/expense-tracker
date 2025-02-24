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
	ID          int
	UserID      int
	Total       float64
	Description string
	Category    ExpenseCategory
	CreatedAt   time.Time
}
