package model

type Task struct {
	ID          string `json:"id"`
	Title       string `json:"title" validate:"required"`
	Description string `json:"description" validate:"required"`
	Done        bool   `json:"done"`
	/*DateStart   time.Time `json:"date_start"  time_format:"0000-00-00 00:00:00"`
	DateStop    time.Time `json:"date_stop"   time_format:"0006-00-00 00:00:00"`
	DateCreate  time.Time `json:"date_create" time_format:"0000-00-00 00:00:00"`*/
}
