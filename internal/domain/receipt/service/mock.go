package service

import (
	"context"

	"github.com/dinata1312/TechGP-Project/internal/dto"
	"github.com/dinata1312/TechGP-Project/pkg/errs"
)

var (
	GET_ALL_BY_ID      func(ctx context.Context, id string) (*dto.GetAllByIdResponseDTO, errs.MessageErr)
	GET_ONE_BY_USER_ID func(ctx context.Context, billId, userId string) (*dto.GetOneByUserIdResponseDTO, errs.MessageErr)
)

type serviceMock struct {
}

func NewServiceMock() ReceiptService {
	return &serviceMock{}
}

func (s *serviceMock) GetAllById(ctx context.Context, id string) (*dto.GetAllByIdResponseDTO, errs.MessageErr) {
	return GET_ALL_BY_ID(ctx, id)
}
func (s *serviceMock) GetOneByUserId(ctx context.Context, billId, userId string) (*dto.GetOneByUserIdResponseDTO, errs.MessageErr) {
	return GET_ONE_BY_USER_ID(ctx, billId, userId)
}
