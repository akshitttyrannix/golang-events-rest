package events

type Event struct {
	EventID     string `json:"event_id"`
	Name        string `json:"name" binding:"required"`
	Description string `json:"description" binding:"required"`
	StartDate   int64  `json:"start_date" binding:"required"`
	EndDate     int64  `json:"end_date" binding:"required"`

	Location  string `json:"location" binding:"required"`
	CreatedAt int64  `json:"created_at"`
	UpdatedAt int64  `json:"updated_at"`
	IsDeleted bool   `json:"is_deleted"`

	UserID string `json:"user_id" binding:"required"`
}
