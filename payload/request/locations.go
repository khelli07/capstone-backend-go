package payload

type CreateLocationRequest struct {
	Name  string `json:"name" binding:"required"`
	Level string `json:"level" binding:"required"`
}

type UpdateLocationRequest struct {
	Name  string `json:"name"`
	Level string `json:"level"`
}
