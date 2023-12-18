package payload

import "mime/multipart"

type CreateEventData struct {
	Name string `form:"name" binding:"required"`
	// Example: cat1,cat2
	Categories  []string `form:"categories"`
	Description string   `form:"description" binding:"required"`
	Price       float32  `form:"price" binding:"required"`
	Capacity    int32    `form:"capacity" binding:"required"`
	Organizer   string   `form:"organizer"`
	DressCode   string   `form:"dress_code"`
	AgeLimit    int      `form:"age_limit"`
	Lat         float64  `form:"lat"`
	Long        float64  `form:"long"`
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
	Categories  []string `form:"categories"`
	Description string   `form:"description" binding:"required"`
	Price       float32  `form:"price" binding:"required"`
	Capacity    int32    `form:"capacity" binding:"required"`
	Organizer   string   `form:"organizer"`
	DressCode   string   `form:"dress_code"`
	AgeLimit    int      `form:"age_limit"`
	Lat         float64  `form:"lat"`
	Long        float64  `form:"long"`
	// Example: 2023-12-01T20:00:00.000Z
	StartTime string `form:"start_time" binding:"required"`
	// Example: 2023-12-01T20:00:00.000Z
	EndTime string `form:"end_time" binding:"required"`
}
