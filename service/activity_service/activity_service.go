package activity_service

import (
	"devcode_challenge/model/request"
	"devcode_challenge/model/response"
)

type ActivityService interface {
	GetAll() ([]*response.ActivityResponse, error)
	DetailActivity(activityId int) (*response.ActivityResponse, error)
	CreateActivity(req *request.ActivityRequest) (*response.ActivityResponse, error)
	UpdateActivity(req *request.ActivityRequest) (*response.ActivityResponse, error)
	DeleteActivity(activityId int) error
}
