package models

type Comment struct {
	ID       uint   `json:"id" gorm:"primaryKey"`
	TicketID uint   `json:"ticket_id"`
	UserID   uint   `json:"user_id"`
	Message  string `json:"message"`
}
