package request

type TodoRequest struct {
	Id              int64
	IsActive        bool   `json:"is_active"`
	ActivityGroupId *int64 `json:"activity_group_id"`
	Title           string `json:"title"`
}
