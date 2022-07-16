package main

import (
	"devcode_challenge/connection"
	"devcode_challenge/migration"
	"devcode_challenge/router"

	"github.com/gofiber/fiber/v2"
)

func main() {
	connection, err := connection.GetConnection()
	if err != nil {
		panic(err)
	}

	err = migration.DbMigration(connection)
	if err != nil {
		panic(err)
	}

	fiber := fiber.New()

	router.Router(fiber, connection)

	fiber.Listen(":3030")
}
