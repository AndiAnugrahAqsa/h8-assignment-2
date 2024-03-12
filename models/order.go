package models

import (
	"time"
)

type Order struct {
	ID           uint       `json:"id"`
	CustomerName string     `json:"customerName"`
	OrderedAt    *time.Time `json:"orderedAt"`
	Items        []Item     `json:"items"`
	CreatedAt    *time.Time `json:"createdAt"`
	UpdatedAt    *time.Time `json:"updatedAt"`
}
