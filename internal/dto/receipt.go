package dto

import "github.com/google/uuid"

type GetAllByIdResponse struct {
	Title     string       `json:"title" example:"Toko Test"`
	Category  string       `json:"category" example:"Makanan"`
	TotalBill float64      `json:"total" example:"38850"`
	Receipts  []ReceiptDTO `json:"receipts"`
} // @name GetAllByIdResponse

type GetAllByIdResponseDTO struct {
	CommonBaseResponseDTO
	Data GetAllByIdResponse `json:"data"`
} // @name GetAllByIdResponseDTO

type ReceiptDTO struct {
	Id        uuid.UUID    `json:"id" example:"f5063dca-556c-4723-931b-cbade7ca139a"`
	Name      string       `json:"name" example:"adit"`
	UserTotal float64      `json:"total" example:"16650"`
	Products  []ProductDTO `json:"products"`
} // @name ReceiptDTO

type GetOneByUserIdResponse struct {
	Title     string     `json:"title" example:"Toko Test"`
	Category  string     `json:"category" example:"Makanan"`
	TotalBill float64    `json:"total" example:"38850"`
	Receipts  ReceiptDTO `json:"receipts"`
} // @name GetOneByUserIdResponse

type GetOneByUserIdResponseDTO struct {
	CommonBaseResponseDTO
	Data GetOneByUserIdResponse `json:"data"`
} // @name GetOneByUserIdResponseDTO
