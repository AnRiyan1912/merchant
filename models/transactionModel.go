package models

import "time"

type Transaction struct {
	ID int64 `gorm:"primary_key" json:"id"`
	SenderID int64 `json:"sender_id"`
	ReceiverID int64 `json:"receiver_id"`
	Amount int64 `json:"amount"`
	CreatedAt time.Time `json:"created_at"`
}