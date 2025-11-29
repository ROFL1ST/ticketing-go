package models

type Ticket struct {
	ID       uint      `json:"id" gorm:"primaryKey"`
	Title    string    `json:"title"`
	Message  string    `json:"message"`
	Status   string    `json:"status"`
	UserID   uint      `json:"user_id"`
	Comments []Comment `json:"comments"`
}
