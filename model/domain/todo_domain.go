package domain

type TodoDomain struct {
	Id              int64
	ActivityGroupId int64
	Title           string
	IsActive        string
	Priority        string
	CreatedAt       string
	UpdatedAt       string
	DeletedAt       *string
}
