package activity_service

import (
	"devcode_challenge/model/domain"
	"devcode_challenge/model/request"
	"devcode_challenge/model/response"
	"devcode_challenge/repository/activity_repository"
	"errors"
	"sync"
	"time"
)

type activityServiceImpl struct {
	Repo activity_repository.ActivityRepository
}

func NewActivityService(repo activity_repository.ActivityRepository) ActivityService {
	var doOnce sync.Once
	service := new(activityServiceImpl)
	doOnce.Do(func() {
		service = &activityServiceImpl{
			Repo: repo,
		}
	})
	return service
}

func (service *activityServiceImpl) GetAll() ([]*response.ActivityResponse, error) {
	domainResult, err := service.Repo.GetAll()
	result := make([]*response.ActivityResponse, 0)
	if err != nil {
		return result, err
	}

	if len(domainResult) == 0 {
		return result, nil
	}

	for _, each := range domainResult {
		activity := &response.ActivityResponse{
			Id:        each.Id,
			Email:     each.Email,
			Title:     each.Title,
			CreatedAt: each.CreatedAt,
			UpdatedAt: each.UpdatedAt,
			DeletedAt: each.DeletedAt,
		}

		result = append(result, activity)
	}

	return result, nil
}

func (service *activityServiceImpl) DetailActivity(activityId int) (*response.ActivityResponse, error) {
	domainResult, err := service.Repo.DetailActivity(activityId)
	if err != nil {
		return nil, err
	}

	return &response.ActivityResponse{
		Id:        domainResult.Id,
		Email:     domainResult.Email,
		Title:     domainResult.Title,
		CreatedAt: domainResult.CreatedAt,
		UpdatedAt: domainResult.UpdatedAt,
		DeletedAt: domainResult.DeletedAt,
	}, nil
}

func (service *activityServiceImpl) CreateActivity(req *request.ActivityRequest) (*response.ActivityResponse, error) {
	if req.Title == "" {
		return nil, errors.New("null title")
	}

	domainReq := &domain.ActivityDomain{
		Email:     req.Email,
		Title:     req.Title,
		CreatedAt: time.Now().String(),
		UpdatedAt: time.Now().String(),
	}

	domainRes, err := service.Repo.CreateActivity(domainReq)
	if err != nil {
		return nil, err
	}

	return &response.ActivityResponse{
		Id:        domainRes.Id,
		Email:     domainRes.Email,
		Title:     domainRes.Title,
		CreatedAt: domainRes.CreatedAt,
		UpdatedAt: domainRes.UpdatedAt,
		DeletedAt: domainRes.DeletedAt,
	}, nil
}

func (service *activityServiceImpl) UpdateActivity(req *request.ActivityRequest) (*response.ActivityResponse, error) {
	if req.Title == "" {
		return nil, errors.New("null title")
	}

	domainReq := &domain.ActivityDomain{
		Id:        req.Id,
		Title:     req.Title,
		UpdatedAt: time.Now().String(),
	}
	_, err := service.Repo.UpdateActivity(domainReq)
	if err != nil {
		return nil, err
	}

	updatedData, err := service.DetailActivity(int(req.Id))
	if err != nil {
		return nil, err
	}

	return &response.ActivityResponse{
		Id:        updatedData.Id,
		Title:     updatedData.Title,
		Email:     updatedData.Email,
		CreatedAt: updatedData.CreatedAt,
		DeletedAt: updatedData.DeletedAt,
		UpdatedAt: updatedData.UpdatedAt,
	}, nil
}

func (service *activityServiceImpl) DeleteActivity(activityId int) error {
	err := service.Repo.DeleteActivity(activityId)
	if err != nil {
		return err
	}

	return nil
}
