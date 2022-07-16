package activity_repository

import "devcode_challenge/model/domain"

type ActivityRepository interface {
	GetAll() ([]*domain.ActivityDomain, error)
	DetailActivity(activityId int) (*domain.ActivityDomain, error)
	CreateActivity(req *domain.ActivityDomain) (*domain.ActivityDomain, error)
	UpdateActivity(req *domain.ActivityDomain) (*domain.ActivityDomain, error)
	DeleteActivity(activityId int) error
}
