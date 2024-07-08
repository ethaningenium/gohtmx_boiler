package entities

type Todo struct {
	ID          string `json:"id" db:"id"`
	Title       string `json:"title" db:"title"`
	IsCompleted string `json:"is_completed" db:"is_completed"`
	UserID      string `json:"user_id" db:"user_id"`
	CreatedAt   string `json:"created_at" db:"created_at"`
}

type User struct {
	ID        string `json:"id" db:"id"`
	Name      string `json:"name" db:"name"`
	Email     string `json:"email" db:"email"`
	Password  string `json:"password" db:"password"`
	CreatedAt string `json:"created_at" db:"created_at"`
}
