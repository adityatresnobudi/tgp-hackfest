package service

import (
	"context"
	"net/http"

	"github.com/dinata1312/TechGP-Project/internal/dto"
	"github.com/dinata1312/TechGP-Project/internal/repositories/receipt_repo"
	"github.com/dinata1312/TechGP-Project/pkg/errs"
	"github.com/google/uuid"
)

type ReceiptService interface {
	GetAllById(ctx context.Context, id string) (*dto.GetAllByIdResponseDTO, errs.MessageErr)
	GetOneByUserId(ctx context.Context, billId, userId string) (*dto.GetOneByUserIdResponseDTO, errs.MessageErr)
}

type receiptServiceIMPL struct {
	receiptRepo receipt_repo.Repository
}

func NewReceiptService(receiptRepo receipt_repo.Repository) ReceiptService {
	return &receiptServiceIMPL{
		receiptRepo: receiptRepo,
	}
}

func (r *receiptServiceIMPL) GetAllById(ctx context.Context, id string) (*dto.GetAllByIdResponseDTO, errs.MessageErr) {
	resp := dto.GetAllByIdResponse{}
	respReceipts := []dto.ReceiptDTO{}
	idChecker := make(map[uuid.UUID]int)
	parsedId, errParseId := uuid.Parse(id)

	if errParseId != nil {
		return nil, errs.NewBadRequest("id has to be a valid uuid")
	}

	fullReceipt, err := r.receiptRepo.GetAllById(ctx, parsedId)
	if err != nil {
		return nil, err
	}

	if len(fullReceipt) <= 0 {
		return nil, errs.NewBadRequest("receipt does not exist")
	}

	resp.Title = fullReceipt[0].StoreName
	resp.Category = fullReceipt[0].Category
	resp.TotalBill = fullReceipt[0].TotalBill

	for idx, value := range fullReceipt {
		if val, ok := idChecker[value.UserId]; ok {
			product := dto.ProductDTO{
				Id:          value.ProductId,
				ProductName: value.ProductName,
				Price:       value.Price,
				Quantity:    value.Quantity,
				Discount:    value.Discount,
				Tax:         value.Tax,
				Service:     value.Service,
			}
			respReceipts[val].UserTotal += value.Total
			respReceipts[val].Products = append(respReceipts[val].Products, product)
		}
		idChecker[value.UserId] = idx
		receipt := dto.ReceiptDTO{
			Id:        value.UserId,
			Name:      value.UserName,
			UserTotal: value.Total,
			Products:  []dto.ProductDTO{},
		}
		product := dto.ProductDTO{
			Id:          value.ProductId,
			ProductName: value.ProductName,
			Price:       value.Price,
			Quantity:    value.Quantity,
			Discount:    value.Discount,
			Tax:         value.Tax,
			Service:     value.Service,
		}
		receipt.Products = append(receipt.Products, product)
		respReceipts = append(respReceipts, receipt)
	}
	resp.Receipts = respReceipts

	result := dto.GetAllByIdResponseDTO{
		CommonBaseResponseDTO: dto.CommonBaseResponseDTO{
			ResponseCode:    http.StatusOK,
			ResponseMessage: "OK",
		},
		Data: resp,
	}

	return &result, nil
}

func (r *receiptServiceIMPL) GetOneByUserId(ctx context.Context, billId, userId string) (*dto.GetOneByUserIdResponseDTO, errs.MessageErr) {
	resp := dto.GetOneByUserIdResponse{}
	respReceipts := dto.ReceiptDTO{}

	parsedBillId, errParseId := uuid.Parse(billId)
	if errParseId != nil {
		return nil, errs.NewBadRequest("bill id has to be a valid uuid")
	}

	parsedUserId, errParseId := uuid.Parse(userId)
	if errParseId != nil {
		return nil, errs.NewBadRequest("user id has to be a valid uuid")
	}

	fullReceipt, err := r.receiptRepo.GetOneByUserId(ctx, parsedBillId, parsedUserId)
	if err != nil {
		return nil, err
	}

	if len(fullReceipt) <= 0 {
		return nil, errs.NewBadRequest("receipt does not exist")
	}

	resp.Title = fullReceipt[0].StoreName
	resp.Category = fullReceipt[0].Category
	resp.TotalBill = fullReceipt[0].TotalBill
	respReceipts.Id = fullReceipt[0].UserId
	respReceipts.Name = fullReceipt[0].UserName
	respReceipts.UserTotal = 0

	for _, value := range fullReceipt {
		product := dto.ProductDTO{
			Id:          value.ProductId,
			ProductName: value.ProductName,
			Price:       value.Price,
			Quantity:    value.Quantity,
			Discount:    value.Discount,
			Tax:         value.Tax,
			Service:     value.Service,
		}
		respReceipts.UserTotal += value.Total
		respReceipts.Products = append(respReceipts.Products, product)
	}
	resp.Receipts = respReceipts

	result := dto.GetOneByUserIdResponseDTO{
		CommonBaseResponseDTO: dto.CommonBaseResponseDTO{
			ResponseCode:    http.StatusOK,
			ResponseMessage: "OK",
		},
		Data: resp,
	}

	return &result, nil
}
