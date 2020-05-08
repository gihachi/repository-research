package main

import (
	"gihachi/repository-research/api"
	"gihachi/repository-research/db/factory"
	"gihachi/repository-research/db/model"
	"net/http"

	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {

	autoMigration()

	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/", root)
	e.GET("/diff", api.GetDiff)
	e.Logger.Fatal(e.Start(":3000"))
}

func root(c echo.Context) error {
	db := factory.DbConnection()
	defer db.Close()
	var testArr []model.Test
	db.Find(&testArr)
	return c.JSON(http.StatusOK, testArr)
}

func autoMigration() {
	db := factory.DbConnection()
	defer db.Close()
	db.AutoMigrate(&model.Test{})
}
