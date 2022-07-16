package response

type ActivityResponse struct {
	Id        int64   `json:"id"`
	Email     string  `json:"email"`
	Title     string  `json:"title"`
	CreatedAt string  `json:"created_at"`
	UpdatedAt string  `json:"updated_at"`
	DeletedAt *string `json:"deleted_at"`
}
