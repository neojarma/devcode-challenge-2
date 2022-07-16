package main

import (
	"devcode_challenge/connection"
	"devcode_challenge/migration"
	"devcode_challenge/router"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cache"
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

	app := fiber.New()
	app.Use(cache.New())

	router.Router(app, connection)

	app.Listen(":3030")
}
