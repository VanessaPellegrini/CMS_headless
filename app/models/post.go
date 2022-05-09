package models

import "time"

type Post struct {
	ID        int64     `json:"id"`
	Title     string    `json:"title"`
	Body      string    `json:"body"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Autor     string    `json:"autor"`
	Category  string    `json:"category"`
	Tags      []string  `json:"tags"`
	Image     string    `json:"image"`
	Comments  []string  `json:"comments"`
}
