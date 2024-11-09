package model

import "time"

type Categories struct {
	ID int `json:"id"`
	Name string `json:"name"`
	Description string `json:"description"`
	DeletedAt *time.Time `json:"deleted_at,omitempty"`
}