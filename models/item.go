package models

import "time"

type Item struct {
	ID          uint       `json:"lineItemId"`
	Code        string     `json:"itemCode"`
	Description string     `json:"description"`
	Quantity    uint       `json:"quantity"`
	OrderID     uint       `json:"orderId"`
	CreatedAt   *time.Time `json:"createdAt"`
	UpdatedAt   *time.Time `json:"updatedAt"`
}
