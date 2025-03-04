package dto

type CommonBaseResponseDTO struct {
	ResponseCode    int    `json:"responseCode" example:"200"`
	ResponseMessage string `json:"responseMessage" example:"Success"`
}
