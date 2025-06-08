package entity

import "time"

type User struct {
	ID        uint      `json:"id"`
	Phone     string    `json:"phone"`
	Password  string    `json:"-"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
