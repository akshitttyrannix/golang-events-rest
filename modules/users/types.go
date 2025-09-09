package users

type User struct {
	UserID    string `json:"user_id"`
	Username  string `json:"username"`
	Email     string `json:"email" binding:"required"`
	Password  string `json:"password" binding:"required"`
	CreatedAt int64  `json:"created_at"`
	UpdatedAt int64  `json:"updated_at"`
	IsDeleted bool   `json:"is_deleted"`
}
