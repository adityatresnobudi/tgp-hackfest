package handler

func (r *receiptHandler) MapRoutes() {
	r.e.Group("").
		GET("/detail-bill/:id", r.GetReceipt)
}
