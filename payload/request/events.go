package payload

import "mime/multipart"

type CreateEventData struct {
	Name string `form:"name" binding:"required"`
	// Example: cat1,cat2
	Categories  []string `form:"categories"`
	Description string   `form:"description" binding:"required"`
	Location    string   `form:"location" binding:"required"`
	Price       float32  `form:"price" binding:"required"`
	Capacity    int32    `form:"capacity" binding:"required"`
	Organizer   string   `form:"organizer"`
	DressCode   string   `form:"dress_code"`
	AgeLimit    int      `form:"age_limit"`
	// Example: 2023-12-01T20:00:00.000Z
	StartTime string `form:"start_time" binding:"required"`
	// Example: 2023-12-01T20:00:00.000Z
	EndTime string `form:"end_time" binding:"required"`
}

type CreateEventRequest struct {
	Image *multipart.FileHeader `form:"image" binding:"required"`
	CreateEventData
}

type UpdateEventRequest struct {
	Name string `form:"name" binding:"required"`
	// Example: cat1,cat2
	Categories  []string `form:"categories" binding:"required"`
	Description string   `form:"description" binding:"required"`
	Location    string   `form:"location" binding:"required"`
	Price       float32  `form:"price" binding:"required"`
	Capacity    int32    `form:"capacity" binding:"required"`
	Organizer   string   `form:"organizer" binding:"required"`
	DressCode   string   `form:"dress_code" binding:"required"`
	AgeLimit    int      `form:"age_limit" binding:"required"`
	// Example: 2023-12-01T20:00:00.000Z
	StartTime string `form:"start_time" binding:"required"`
	// Example: 2023-12-01T20:00:00.000Z
	EndTime string `form:"end_time" binding:"required"`
}
