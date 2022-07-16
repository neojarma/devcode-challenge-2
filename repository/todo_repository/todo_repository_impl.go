package todo_repository

import (
	"database/sql"
	"devcode_challenge/model/domain"
	"errors"
	"sync"
)

type todoRepositoryImpl struct {
	Db *sql.DB
}

func NewTodoRepository(db *sql.DB) TodoRepository {
	var doOnce sync.Once
	repo := new(todoRepositoryImpl)
	doOnce.Do(func() {
		repo = &todoRepositoryImpl{
			Db: db,
		}
	})

	return repo
}

func (repo *todoRepositoryImpl) GetAll(activityGroupId *string) ([]*domain.TodoDomain, error) {
	SQL := "SELECT id, title, is_active, priority, activity_group_id, created_at, updated_at, deleted_at FROM todos"
	SQLWithParams := "SELECT id, title, is_active, priority, activity_group_id, created_at, updated_at, deleted_at FROM todos WHERE activity_group_id = ?"

	result := make([]*domain.TodoDomain, 0)
	if activityGroupId == nil {
		rows, err := repo.Db.Query(SQL)
		if err != nil {
			return result, err
		}

		defer rows.Close()

		for rows.Next() {
			todo := new(domain.TodoDomain)
			err := rows.Scan(&todo.Id, &todo.Title, &todo.IsActive, &todo.Priority, &todo.ActivityGroupId, &todo.CreatedAt, &todo.UpdatedAt, &todo.DeletedAt)
			if err != nil {
				return result, err
			}

			result = append(result, todo)
		}
	}

	rows, err := repo.Db.Query(SQLWithParams, activityGroupId)
	if err != nil {
		return result, err
	}

	defer rows.Close()

	for rows.Next() {
		todo := new(domain.TodoDomain)
		err := rows.Scan(&todo.Id, &todo.Title, &todo.IsActive, &todo.Priority, &todo.ActivityGroupId, &todo.CreatedAt, &todo.UpdatedAt, &todo.DeletedAt)
		if err != nil {
			return result, err
		}

		result = append(result, todo)
	}

	return result, nil
}

func (repo *todoRepositoryImpl) DetailTodo(todoId int) (*domain.TodoDomain, error) {
	SQL := "SELECT id, title, is_active, priority, activity_group_id, created_at, updated_at, deleted_at FROM todos WHERE id = ?"

	rows, err := repo.Db.Query(SQL, todoId)
	result := new(domain.TodoDomain)
	if err != nil {
		return result, err
	}

	defer rows.Close()

	if rows.Next() {
		err := rows.Scan(&result.Id, &result.Title, &result.IsActive, &result.Priority, &result.ActivityGroupId, &result.CreatedAt, &result.UpdatedAt, &result.DeletedAt)
		if err != nil {
			return result, err
		}

		return result, nil
	}

	return result, errors.New("not found")
}

func (repo *todoRepositoryImpl) CreateTodo(req *domain.TodoDomain) (*domain.TodoDomain, error) {
	SQL := "INSERT INTO todos (id, title, activity_group_id, created_at, updated_at) VALUES(?,?,?,?,?)"
	result, err := repo.Db.Exec(SQL, req.Id, req.Title, req.ActivityGroupId, req.CreatedAt, req.UpdatedAt)
	if err != nil {
		return nil, err
	}

	rowNumber, err := result.LastInsertId()
	if err != nil {
		return nil, err
	}

	req.Id = rowNumber
	return req, nil
}

func (repo *todoRepositoryImpl) UpdateTodo(req *domain.TodoDomain) (*domain.TodoDomain, error) {
	updateTitleSQL := "UPDATE todos SET title = ?, updated_at = ? WHERE id = ?"
	updateIsActiveSQL := "UPDATE todos SET is_active = ?, updated_at = ? WHERE id = ?"

	if req.Title != "" {
		result, err := repo.Db.Exec(updateTitleSQL, req.Title, req.UpdatedAt, req.Id)
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

	result, err := repo.Db.Exec(updateIsActiveSQL, req.IsActive, req.UpdatedAt, req.Id)
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

func (repo *todoRepositoryImpl) DeleteTodo(todoId int) error {
	SQL := "DELETE FROM todos WHERE id = ?"
	result, err := repo.Db.Exec(SQL, todoId)
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
