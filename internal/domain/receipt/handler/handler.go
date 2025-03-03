package handler

import (
	"context"
	"net/http"

	"github.com/dinata1312/TechGP-Project/internal/domain/receipt/service"
	"github.com/gin-gonic/gin"
)

type receiptHandler struct {
	e       *gin.Engine
	ctx     context.Context
	service service.ReceiptService
}

func NewReceiptHandler(
	e *gin.Engine,
	ctx context.Context,
	service service.ReceiptService,
) *receiptHandler {
	return &receiptHandler{
		e:       e,
		ctx:     ctx,
		service: service,
	}
}

// @Summary Get Receipts By User ID
// @Tags receipts
// @Produce json
// @Success 200 {object}  GetAllByIdResponse
// @Router /detail-bill/{bill-id} [get]
func (r *receiptHandler) GetReceipt(c *gin.Context) {
	billId := c.Param("id")
	userId := c.Query("userId")

	if userId != "" {
		result, err := r.service.GetOneByUserId(r.ctx, billId, userId)
		if err != nil {
			c.JSON(err.StatusCode(), err)
			return
		}
		c.JSON(http.StatusOK, result)
		return
	}

	result, err := r.service.GetAllById(r.ctx, billId)
	if err != nil {
		c.JSON(err.StatusCode(), err)
		return
	}

	c.JSON(http.StatusOK, result)
}
