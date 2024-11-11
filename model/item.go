package model

import "time"

type Items struct {
	ID int `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
	CategoryId int `json:"category_id,omitempty"`
	CategoryName string `json:"category,omitempty"`
	Image string `json:"image,omitempty"`
	Price int `json:"price,omitempty"`
	TransactionDate time.Time `json:"transaction_date"`
	Depreciation int `json:"depresiasi,omitempty"`
	DepreciatedValue  int `json:"depreciated_value"`
	TotalDaysUsage int `json:"total_days_usage,omitempty"`
	DeletedAt       *time.Time `json:"deleted_at,omitempty"`
}
