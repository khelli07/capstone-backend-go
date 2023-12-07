package payload

type GeneralResponse struct {
	Message string `json:"message"`
}

type CreateResponse struct {
	ID string `json:"id"`
}
