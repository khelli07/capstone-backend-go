package models

type Event struct {
	Name        string
	Category    string
	Description string `datastore:",noindex"`
	// TimeStart   time.Time
	// TimeEnd     time.Time
}
