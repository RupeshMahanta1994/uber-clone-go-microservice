package models

import "time"

type Profile struct {
	UserID    string    `json:"userId" db:"user_id"`
	Name      string    `json:"name" db:"name"`
	Phone     string    `json:"phone" db:"phone"`
	Rating    float64   `json:"rating,omitempty" db:"rating"`
	CreatedAt time.Time `json:"created_at,omitempty" db:"created_at"`
}
