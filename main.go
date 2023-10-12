package main

import (
	"github.com/labstack/echo/v4"
	"simple-rest/config"
	"simple-rest/controller"
)

func main() {
	e := echo.New()

	config.DbInit()
	gorm := config.Db()

	dbGorm, err := gorm.DB()
	if err != nil {
		e.Logger.Fatalf("Failed to connect to the database: %v", err)
	}

	err = dbGorm.Ping()
	if err != nil {
		e.Logger.Fatalf("Failed to ping the database: %v", err)
	}

	bookRoute := e.Group("/book")
	bookRoute.POST("/", controller.CreateBook)
	bookRoute.GET("/:id", controller.GetBook)
	bookRoute.PUT("/:id", controller.UpdateBook)
	bookRoute.DELETE("/:id", controller.DeleteBook)

	err = e.Start(":4545")
	if err != nil {
		e.Logger.Fatalf("Failed to start the server: %v", err)
	}
}
