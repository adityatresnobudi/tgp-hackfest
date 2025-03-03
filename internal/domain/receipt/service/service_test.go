package service

import (
	"context"
	"net/http"
	"testing"

	"github.com/dinata1312/TechGP-Project/internal/entity"
	"github.com/dinata1312/TechGP-Project/internal/repositories/receipt_repo"
	"github.com/dinata1312/TechGP-Project/pkg/errs"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func Test_GetAllById_Success(t *testing.T) {
	r := receipt_repo.NewRepoMock()
	s := NewReceiptService(r)
	userId, err := uuid.Parse("f5063dca-556c-4723-931b-cbade7ca139a")
	assert.NoError(t, err)
	productId, err := uuid.Parse("516af9d0-d49e-4bac-8cab-932c6b622682")
	assert.NoError(t, err)
	fullReceipt := []entity.FullReceipt{
		{
			StoreName:   "Toko Test",
			Category:    "Makanan",
			TotalBill:   38850,
			UserId:      userId,
			UserName:    "Irfan",
			ProductId:   productId,
			ProductName: "Nasi Goreng",
			Price:       20000,
			Quantity:    1,
			Discount:    0,
			Tax:         11,
			Service:     0,
			Total:       22200,
		},
	}

	ctx := context.TODO()

	receipt_repo.GET_ALL_BY_ID = func(ctx context.Context, id uuid.UUID) ([]entity.FullReceipt, errs.MessageErr) {
		return fullReceipt, nil
	}

	result, err := s.GetAllById(ctx, "ad9197e5-cb9a-419c-b055-596962d5501e")

	assert.Nil(t, err)

	assert.NotNil(t, result)

	assert.Equal(t, http.StatusOK, result.CommonBaseResponseDTO.ResponseCode)
	assert.Equal(t, "OK", result.CommonBaseResponseDTO.ResponseMessage)
	assert.Equal(t, fullReceipt[0].StoreName, result.Data.Title)
	assert.Equal(t, fullReceipt[0].Category, result.Data.Category)
	assert.Equal(t, fullReceipt[0].TotalBill, result.Data.TotalBill)
	assert.Equal(t, fullReceipt[0].UserId, result.Data.Receipts[0].Id)
	assert.Equal(t, fullReceipt[0].UserName, result.Data.Receipts[0].Name)
	assert.Equal(t, fullReceipt[0].Total, result.Data.Receipts[0].UserTotal)
	assert.Equal(t, fullReceipt[0].ProductId, result.Data.Receipts[0].Products[0].Id)
	assert.Equal(t, fullReceipt[0].ProductName, result.Data.Receipts[0].Products[0].ProductName)
	assert.Equal(t, fullReceipt[0].Price, result.Data.Receipts[0].Products[0].Price)
	assert.Equal(t, fullReceipt[0].Quantity, result.Data.Receipts[0].Products[0].Quantity)
	assert.Equal(t, fullReceipt[0].Discount, result.Data.Receipts[0].Products[0].Discount)
	assert.Equal(t, fullReceipt[0].Tax, result.Data.Receipts[0].Products[0].Tax)
	assert.Equal(t, fullReceipt[0].Service, result.Data.Receipts[0].Products[0].Service)
}

func Test_GetAllById_ErrorIdNotValidUUID(t *testing.T) {
	s := NewReceiptService(nil)

	ctx := context.TODO()

	result, err := s.GetAllById(ctx, "1")

	assert.NotNil(t, err)

	assert.Nil(t, result)

	assert.Equal(t, "id has to be a valid uuid", err.Error())

	assert.Equal(t, http.StatusBadRequest, err.StatusCode())
}

func Test_GetAllById_ErrorInternalServerError(t *testing.T) {
	r := receipt_repo.NewRepoMock()
	s := NewReceiptService(r)

	ctx := context.TODO()

	receipt_repo.GET_ALL_BY_ID = func(ctx context.Context, id uuid.UUID) ([]entity.FullReceipt, errs.MessageErr) {
		return nil, errs.NewInternalServerError()
	}

	result, err := s.GetAllById(ctx, "f8d8e1ea-b7d6-43bc-948f-fddbc79949d9")

	assert.NotNil(t, err)

	assert.Nil(t, result)

	assert.Equal(t, "Something went wrong", err.Error())

	assert.Equal(t, http.StatusInternalServerError, err.StatusCode())
}

func Test_GetAllById_ErrorReceiptDoesNotExist(t *testing.T) {
	r := receipt_repo.NewRepoMock()
	s := NewReceiptService(r)

	ctx := context.TODO()

	receipt_repo.GET_ALL_BY_ID = func(ctx context.Context, id uuid.UUID) ([]entity.FullReceipt, errs.MessageErr) {
		return []entity.FullReceipt{}, nil
	}

	result, err := s.GetAllById(ctx, "f8d8e1ea-b7d6-43bc-948f-fddbc79949d9")

	assert.NotNil(t, err)

	assert.Nil(t, result)

	assert.Equal(t, "receipt does not exist", err.Error())

	assert.Equal(t, http.StatusBadRequest, err.StatusCode())
}

func Test_GetOneByUserId_Success(t *testing.T) {
	r := receipt_repo.NewRepoMock()
	s := NewReceiptService(r)
	userId, err := uuid.Parse("f5063dca-556c-4723-931b-cbade7ca139a")
	assert.NoError(t, err)
	productId, err := uuid.Parse("516af9d0-d49e-4bac-8cab-932c6b622682")
	assert.NoError(t, err)
	fullReceipt := []entity.FullReceipt{
		{
			StoreName:   "Toko Test",
			Category:    "Makanan",
			TotalBill:   38850,
			UserId:      userId,
			UserName:    "Irfan",
			ProductId:   productId,
			ProductName: "Nasi Goreng",
			Price:       20000,
			Quantity:    1,
			Discount:    0,
			Tax:         11,
			Service:     0,
			Total:       22200,
		},
	}

	ctx := context.TODO()

	receipt_repo.GET_ONE_BY_USER_ID = func(ctx context.Context, billId, userId uuid.UUID) ([]entity.FullReceipt, errs.MessageErr) {
		return fullReceipt, nil
	}

	result, err := s.GetOneByUserId(ctx, "ad9197e5-cb9a-419c-b055-596962d5501e", "f5063dca-556c-4723-931b-cbade7ca139a")

	assert.Nil(t, err)

	assert.NotNil(t, result)

	assert.Equal(t, http.StatusOK, result.CommonBaseResponseDTO.ResponseCode)
	assert.Equal(t, "OK", result.CommonBaseResponseDTO.ResponseMessage)
	assert.Equal(t, fullReceipt[0].StoreName, result.Data.Title)
	assert.Equal(t, fullReceipt[0].Category, result.Data.Category)
	assert.Equal(t, fullReceipt[0].TotalBill, result.Data.TotalBill)
	assert.Equal(t, fullReceipt[0].UserId, result.Data.Receipts.Id)
	assert.Equal(t, fullReceipt[0].UserName, result.Data.Receipts.Name)
	assert.Equal(t, fullReceipt[0].Total, result.Data.Receipts.UserTotal)
	assert.Equal(t, fullReceipt[0].ProductId, result.Data.Receipts.Products[0].Id)
	assert.Equal(t, fullReceipt[0].ProductName, result.Data.Receipts.Products[0].ProductName)
	assert.Equal(t, fullReceipt[0].Price, result.Data.Receipts.Products[0].Price)
	assert.Equal(t, fullReceipt[0].Quantity, result.Data.Receipts.Products[0].Quantity)
	assert.Equal(t, fullReceipt[0].Discount, result.Data.Receipts.Products[0].Discount)
	assert.Equal(t, fullReceipt[0].Tax, result.Data.Receipts.Products[0].Tax)
	assert.Equal(t, fullReceipt[0].Service, result.Data.Receipts.Products[0].Service)
}

func Test_GetOneByUserId_ErrorBillIdNotValidUUID(t *testing.T) {
	s := NewReceiptService(nil)

	ctx := context.TODO()

	result, err := s.GetOneByUserId(ctx, "1", "f5063dca-556c-4723-931b-cbade7ca139a")

	assert.NotNil(t, err)

	assert.Nil(t, result)

	assert.Equal(t, "bill id has to be a valid uuid", err.Error())

	assert.Equal(t, http.StatusBadRequest, err.StatusCode())
}

func Test_GetOneByUserId_ErrorUserIdNotValidUUID(t *testing.T) {
	s := NewReceiptService(nil)

	ctx := context.TODO()

	result, err := s.GetOneByUserId(ctx, "ad9197e5-cb9a-419c-b055-596962d5501e", "1")

	assert.NotNil(t, err)

	assert.Nil(t, result)

	assert.Equal(t, "user id has to be a valid uuid", err.Error())

	assert.Equal(t, http.StatusBadRequest, err.StatusCode())
}

func Test_GetOneByUserId_ErrorInternalServerError(t *testing.T) {
	r := receipt_repo.NewRepoMock()
	s := NewReceiptService(r)

	ctx := context.TODO()

	receipt_repo.GET_ONE_BY_USER_ID = func(ctx context.Context, billId, userId uuid.UUID) ([]entity.FullReceipt, errs.MessageErr) {
		return nil, errs.NewInternalServerError()
	}

	result, err := s.GetOneByUserId(ctx, "f8d8e1ea-b7d6-43bc-948f-fddbc79949d9", "ad9197e5-cb9a-419c-b055-596962d5501e")

	assert.NotNil(t, err)

	assert.Nil(t, result)

	assert.Equal(t, "Something went wrong", err.Error())

	assert.Equal(t, http.StatusInternalServerError, err.StatusCode())
}

func Test_GetOneByUserId_ErrorReceiptDoesNotExist(t *testing.T) {
	r := receipt_repo.NewRepoMock()
	s := NewReceiptService(r)

	ctx := context.TODO()

	receipt_repo.GET_ONE_BY_USER_ID = func(ctx context.Context, billId, userId uuid.UUID) ([]entity.FullReceipt, errs.MessageErr) {
		return []entity.FullReceipt{}, nil
	}

	result, err := s.GetOneByUserId(ctx, "f8d8e1ea-b7d6-43bc-948f-fddbc79949d9", "f8d8e1ea-b7d6-43bc-948f-fddbc79949d9")

	assert.NotNil(t, err)

	assert.Nil(t, result)

	assert.Equal(t, "receipt does not exist", err.Error())

	assert.Equal(t, http.StatusBadRequest, err.StatusCode())
}
