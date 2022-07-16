package request

type ActivityRequest struct {
	Id    int64
	Email string `json:"email"`
	Title string `json:"title"`
}
