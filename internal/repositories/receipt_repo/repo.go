package receipt_repo

import (
	"context"

	"github.com/dinata1312/TechGP-Project/internal/entity"
	"github.com/dinata1312/TechGP-Project/pkg/errs"
	"github.com/google/uuid"
)

type Repository interface {
	GetAllById(ctx context.Context, id uuid.UUID) ([]entity.FullReceipt, errs.MessageErr)
	GetOneByUserId(ctx context.Context, billId, userId uuid.UUID) ([]entity.FullReceipt, errs.MessageErr)
}
