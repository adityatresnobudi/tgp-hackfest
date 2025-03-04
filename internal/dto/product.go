package dto

import "github.com/google/uuid"

type ProductDTO struct {
	Id          uuid.UUID `json:"id" example:"f5063dca-556c-4723-931b-cbade7ca139a"`
	ProductName string    `json:"productName" example:"bakso"`
	Price       float64   `json:"price" example:"15000"`
	Quantity    int       `json:"quantity" example:"1"`
	Discount    float64   `json:"discount" example:"0"`
	Tax         float64   `json:"tax" example:"11"`
	Service     float64   `json:"service" example:"0"`
} // @name ProductDTO
