package receipt_pg

import (
	"context"
	"database/sql"
	"log"

	"github.com/dinata1312/TechGP-Project/internal/entity"
	"github.com/dinata1312/TechGP-Project/internal/repositories/receipt_repo"
	"github.com/dinata1312/TechGP-Project/pkg/errs"
	"github.com/google/uuid"
)

type receiptPG struct {
	db *sql.DB
}

func NewRepo(db *sql.DB) receipt_repo.Repository {
	return &receiptPG{
		db: db,
	}
}

func (r *receiptPG) GetAllById(ctx context.Context, id uuid.UUID) ([]entity.FullReceipt, errs.MessageErr) {
	rows, err := r.db.QueryContext(ctx, GET_ALL_RECEIPTS, id)

	if err != nil {
		log.Printf("db get all receipts: %s\n", err.Error())
		return nil, errs.NewInternalServerError()
	}

	result := []entity.FullReceipt{}

	for rows.Next() {
		fullReceipt := entity.FullReceipt{}

		if err = rows.Scan(
			&fullReceipt.StoreName,
			&fullReceipt.Category,
			&fullReceipt.TotalBill,
			&fullReceipt.UserId,
			&fullReceipt.UserName,
			&fullReceipt.ProductId,
			&fullReceipt.ProductName,
			&fullReceipt.Price,
			&fullReceipt.Quantity,
			&fullReceipt.Discount,
			&fullReceipt.Tax,
			&fullReceipt.Service,
			&fullReceipt.Total,
		); err != nil {
			log.Printf("db scan get full receipts: %s\n", err.Error())
			return nil, errs.NewInternalServerError()
		}

		result = append(result, fullReceipt)
	}

	return result, nil
}

func (r *receiptPG) GetOneByUserId(ctx context.Context, billId, userId uuid.UUID) ([]entity.FullReceipt, errs.MessageErr) {
	rows, err := r.db.QueryContext(ctx, GET_ONE_RECEIPT_BY_USER_ID, billId, userId)

	if err != nil {
		log.Printf("db get receipts by user id: %s\n", err.Error())
		return nil, errs.NewInternalServerError()
	}

	result := []entity.FullReceipt{}

	for rows.Next() {
		fullReceipt := entity.FullReceipt{}

		if err = rows.Scan(
			&fullReceipt.StoreName,
			&fullReceipt.Category,
			&fullReceipt.TotalBill,
			&fullReceipt.UserId,
			&fullReceipt.UserName,
			&fullReceipt.ProductId,
			&fullReceipt.ProductName,
			&fullReceipt.Price,
			&fullReceipt.Quantity,
			&fullReceipt.Discount,
			&fullReceipt.Tax,
			&fullReceipt.Service,
			&fullReceipt.Total,
		); err != nil {
			log.Printf("db scan get receipt by user id: %s\n", err.Error())
			return nil, errs.NewInternalServerError()
		}

		result = append(result, fullReceipt)
	}

	return result, nil
}