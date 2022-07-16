package response

type TodoResponse struct {
	Id              int64   `json:"id"`
	ActivityGroupId any     `json:"activity_group_id"`
	Title           string  `json:"title"`
	IsActive        any     `json:"is_active"`
	Priority        string  `json:"priority"`
	CreatedAt       string  `json:"created_at"`
	UpdatedAt       string  `json:"updated_at"`
	DeletedAt       *string `json:"deleted_at"`
}

type TodoResponsePost struct {
	Id              int64  `json:"id"`
	ActivityGroupId any    `json:"activity_group_id"`
	Title           string `json:"title"`
	IsActive        any    `json:"is_active"`
	Priority        string `json:"priority"`
	CreatedAt       string `json:"created_at"`
	UpdatedAt       string `json:"updated_at"`
}
