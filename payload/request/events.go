package payload

type CreateEventRequest struct {
	Name        string   `json:"name"`
	Categories  []string `json:"categories"`
	Description string   `json:"description"`
	Location    string   `json:"location"`
	Price       float32  `json:"price"`
	Capacity    int32    `json:"capacity"`
	Organizer   string   `json:"organizer"`
	DressCode   string   `json:"dress_code"`
	AgeLimit    int      `json:"age_limit"`
	StartTime   string   `json:"start_time"`
	EndTime     string   `json:"end_time"`
}

type UpdateEventRequest struct {
	Name        string   `json:"name"`
	Categories  []string `json:"categories"`
	Description string   `json:"description"`
	Location    string   `json:"location"`
	Price       float32  `json:"price"`
	Capacity    int32    `json:"capacity"`
	Organizer   string   `json:"organizer"`
	DressCode   string   `json:"dress_code"`
	AgeLimit    int      `json:"age_limit"`
	StartTime   string   `json:"start_time"`
	EndTime     string   `json:"end_time"`
}
