package main

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"simple-rest/config"
)

type Book struct {
	Id          int    `gorm:"primaryKey" json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

func main() {
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.JSON(http.StatusOK, map[string]interface{}{
			"hello": "world",
		})
	})
	e.Logger.Fatal(e.Start(":8080"))

	config.DbInit()
	gorm := config.Db()

	dbGorm, err := gorm.DB()
	if err != nil {
		panic(err)
	}
	dbGorm.Ping()
	e.Logger.Fatal(e.Start(":8080"))
}
