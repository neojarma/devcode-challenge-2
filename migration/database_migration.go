package migration

import (
	"database/sql"
	"sync"
)

func DbMigration(connection *sql.DB) error {
	activitiesTAble := `
	CREATE TABLE IF NOT EXISTS activities (
		id bigint NOT NULL AUTO_INCREMENT,
		email longtext,
		title longtext,
		created_at longtext DEFAULT NULL,
		updated_at longtext DEFAULT NULL,
		deleted_at longtext DEFAULT NULL,
		PRIMARY KEY (id));
	`

	todosTable := `
	CREATE TABLE IF NOT EXISTS todos (
		id bigint NOT NULL AUTO_INCREMENT,
		created_at longtext DEFAULT NULL,
		updated_at longtext DEFAULT NULL,
		deleted_at longtext DEFAULT NULL,
		activity_group_id bigint DEFAULT NULL,
		title longtext,
		is_active varchar(5) DEFAULT true,
		priority longtext DEFAULT "very-high",
		PRIMARY KEY (id));
	`

	var wg sync.WaitGroup

	wg.Add(2)
	var sqlError error

	go func() {
		_, err := connection.Exec(activitiesTAble)
		if err != nil {
			sqlError = err
		}
		wg.Done()
	}()

	go func() {
		_, err := connection.Exec(todosTable)
		if err != nil {
			sqlError = err
		}
		wg.Done()
	}()

	wg.Wait()
	return sqlError
}
