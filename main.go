package main

import (
	"test/mnc/config"
	"test/mnc/factory"
	"test/mnc/routes"
	// "test/mnc/middlewares"
)

func main() {
	dbConn := config.InitDB()

	presenter := factory.InitFactory(dbConn)
	e := routes.New(presenter)
	// middlewares.LogMiddleware(e)
	e.Logger.Fatal(e.Start(":80"))
}
