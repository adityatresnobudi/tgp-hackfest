package entity

import "github.com/google/uuid"

type FullReceipt struct {
	StoreName   string
	Category    string
	TotalBill   float64
	UserId      uuid.UUID
	UserName    string
	ProductId   uuid.UUID
	ProductName string
	Price       float64
	Quantity    int
	Discount    float64
	Tax         float64
	Service     float64
	Total       float64
}
