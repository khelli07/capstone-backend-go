package payload

import "mime/multipart"

type CreateEventData struct {
	Name        string   `json:"name" binding:"required"`
	Categories  []string `json:"categories"`
	Description string   `json:"description" binding:"required"`
	Location    string   `json:"location" binding:"required"`
	Price       float32  `json:"price" binding:"required"`
	Capacity    int32    `json:"capacity" binding:"required"`
	Organizer   string   `json:"organizer"`
	DressCode   string   `json:"dress_code"`
	AgeLimit    int      `json:"age_limit"`
	StartTime   string   `json:"start_time" binding:"required"`
	EndTime     string   `json:"end_time" binding:"required"`
}

type CreateEventRequest struct {
	Image *multipart.FileHeader `form:"image" binding:"required"`
	Data  CreateEventData       `form:"data" binding:"required"`
}

type UpdateEventRequest struct {
	Name        string   `json:"name" binding:"required"`
	Categories  []string `json:"categories" binding:"required"`
	Description string   `json:"description" binding:"required"`
	Location    string   `json:"location" binding:"required"`
	Price       float32  `json:"price" binding:"required"`
	Capacity    int32    `json:"capacity" binding:"required"`
	Organizer   string   `json:"organizer" binding:"required"`
	DressCode   string   `json:"dress_code" binding:"required"`
	AgeLimit    int      `json:"age_limit" binding:"required"`
	StartTime   string   `json:"start_time" binding:"required"`
	EndTime     string   `json:"end_time" binding:"required"`
}
