package connection

import (
	"database/sql"
	"fmt"
	"os"
	"sync"

	_ "github.com/go-sql-driver/mysql"
)

func GetConnection() (*sql.DB, error) {
	var doOnce sync.Once

	connection := new(sql.DB)
	var err error

	doOnce.Do(func() {
		dbUsername := os.Getenv("MYSQL_USER")
		dbPassword := os.Getenv("MYSQL_PASSWORD")
		dbName := os.Getenv("MYSQL_DBNAME")
		dbHost := os.Getenv("MYSQL_HOST")

		dataSourceName := fmt.Sprintf("%s:%s@tcp(%s:3306)/%s", dbUsername, dbPassword, dbHost, dbName)
		connection, err = sql.Open("mysql", dataSourceName)
	})

	if err != nil {
		return nil, err
	}

	return connection, nil
}
