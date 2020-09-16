package response

//ErrorResponseData -
type ErrorResponseData struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

//ErrorResponse -
type ErrorResponse struct {
	Success bool              `json:"success"`
	Error   ErrorResponseData `json:"data"`
}
