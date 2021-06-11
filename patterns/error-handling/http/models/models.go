package models

import (
	"time"
)

// Note represents the Note model, everyone operates with
type Note struct {
	ID          string    `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"created_at"`
}
