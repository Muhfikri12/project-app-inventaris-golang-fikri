package model

type Response struct {
	Status bool `json:"status"`
	StatusCode int `json:"status_code,omitempty"`
	Message string `json:"message,omitempty"`
	Page int `json:"page,omitempty"`
	Limit int `json:"limit,omitempty"`
	TotalItems int `json:"total_item,omitempty"`
	TotalPages int `json:"total_pages,omitempty"`
	Data any `json:"data,omitempty"`
}