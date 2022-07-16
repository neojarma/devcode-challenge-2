package activity_repository

import (
	"database/sql"
	"devcode_challenge/model/domain"
	"errors"
	"sync"
)

type activityRepositoryImpl struct {
	Db *sql.DB
}

func NewActivityRepository(db *sql.DB) ActivityRepository {
	var doOnce sync.Once
	repo := new(activityRepositoryImpl)
	doOnce.Do(func() {
		repo = &activityRepositoryImpl{
			Db: db,
		}
	})

	return repo
}

func (repo *activityRepositoryImpl) GetAll() ([]*domain.ActivityDomain, error) {
	SQL := "SELECT id, email, title, created_at, updated_at, deleted_at FROM activities"
	result := make([]*domain.ActivityDomain, 0)

	rows, err := repo.Db.Query(SQL)
	if err != nil {
		return result, err
	}
	defer rows.Close()

	for rows.Next() {
		activity := new(domain.ActivityDomain)
		err := rows.Scan(&activity.Id, &activity.Email, &activity.Title, &activity.CreatedAt, &activity.UpdatedAt, &activity.DeletedAt)
		if err != nil {
			return result, err
		}

		result = append(result, activity)
	}

	return result, nil
}

func (repo *activityRepositoryImpl) DetailActivity(activityId int) (*domain.ActivityDomain, error) {
	SQL := "SELECT id, email, title, created_at, updated_at, deleted_at FROM activities WHERE id = ?"
	rows, err := repo.Db.Query(SQL, activityId)
	result := new(domain.ActivityDomain)
	if err != nil {
		return result, err
	}
	defer rows.Close()

	if rows.Next() {
		err := rows.Scan(&result.Id, &result.Email, &result.Title, &result.CreatedAt, &result.UpdatedAt, &result.DeletedAt)
		if err != nil {
			return result, err
		}

		return result, nil
	}

	return result, errors.New("not found")
}

func (repo *activityRepositoryImpl) CreateActivity(req *domain.ActivityDomain) (*domain.ActivityDomain, error) {
	SQL := "INSERT INTO activities (id, email, title, created_at, updated_at) VALUES(?,?,?,?,?)"
	res, err := repo.Db.Exec(SQL, req.Id, req.Email, req.Title, req.CreatedAt, req.UpdatedAt)
	result := new(domain.ActivityDomain)
	if err != nil {
		return result, err
	}

	rowNumber, err := res.LastInsertId()
	if err != nil {
		return result, err
	}

	req.Id = rowNumber
	return req, nil
}

func (repo *activityRepositoryImpl) UpdateActivity(req *domain.ActivityDomain) (*domain.ActivityDomain, error) {
	SQL := "UPDATE activities SET title = ?, updated_at = ? WHERE id = ?"
	result, err := repo.Db.Exec(SQL, req.Title, req.UpdatedAt, req.Id)
	if err != nil {
		return nil, err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return nil, err
	}

	if rowsAffected == 0 {
		return nil, errors.New("not found")
	}

	return req, nil
}

func (repo *activityRepositoryImpl) DeleteActivity(activityId int) error {
	SQL := "DELETE FROM activities WHERE id = ?"
	result, err := repo.Db.Exec(SQL, activityId)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return errors.New("not found")
	}

	return nil
}
