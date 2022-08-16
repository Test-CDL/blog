package main

import (
	"chrombit/config"
	"chrombit/factory"
	"chrombit/routes"

	"chrombit/middlewares"
)

func main() {
	dbConn := config.InitDB()

	presenter := factory.InitFactory(dbConn)
	e := routes.New(presenter)
	middlewares.LogMiddleware(e)
	e.Logger.Fatal(e.Start(":80"))
}
