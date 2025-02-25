package expensetracker

type User struct {
	ID       int    `json:"-" db:"id"`
	Name     string `json:"name" db:"name"`
	Username string `json:"username" binding:"required" db:"username"`
	Password string `json:"password" binding:"required" db:"password_hash"`
}
