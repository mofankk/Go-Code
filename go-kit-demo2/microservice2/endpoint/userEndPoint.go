package endpoint

type UserRequest struct {
	Id     int    `json:"id"`
	Method string `json:"method"`
}

type UserResponse struct {
	Data string `json:"data"`
}
