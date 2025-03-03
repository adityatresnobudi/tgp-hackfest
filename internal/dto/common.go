package dto

type CommonBaseResponseDTO struct {
	ResponseCode    int    `json:"response_code" example:"200"`
	ResponseMessage string `json:"responseMessage" example:"Success"`
	Message         string `json:"message,omitempty" example:"OK"`
}
