package receipt_repo

import (
	"context"

	"github.com/dinata1312/TechGP-Project/internal/entity"
	"github.com/dinata1312/TechGP-Project/pkg/errs"
	"github.com/google/uuid"
)

var (
	GET_ALL_BY_ID      func(ctx context.Context, id uuid.UUID) ([]entity.FullReceipt, errs.MessageErr)
	GET_ONE_BY_USER_ID func(ctx context.Context, billId, userId uuid.UUID) ([]entity.FullReceipt, errs.MessageErr)
)

type repoMock struct {
}

func NewRepoMock() Repository {
	return &repoMock{}
}

func (r *repoMock) GetAllById(ctx context.Context, id uuid.UUID) ([]entity.FullReceipt, errs.MessageErr) {
	return GET_ALL_BY_ID(ctx, id)
}
func (r *repoMock) GetOneByUserId(ctx context.Context, billId, userId uuid.UUID) ([]entity.FullReceipt, errs.MessageErr) {
	return GET_ONE_BY_USER_ID(ctx, billId, userId)
}
