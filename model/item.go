package model

import "time"

type Items struct {
	ID int `json:"id"`
	Name string `json:"name"`
	CategoryId int `json:"category_id,omitempty"`
	CategoryName string `json:"category,omitempty"`
	Image string `json:"image"`
	Price int `json:"price"`
	TransactionDate time.Time `json:"transaction_date"`
	Deprisiasi int `json:"deprisiasi,omitempty"`
	TotalDaysUsage int `json:"total_days_usage,omitempty"`
}
