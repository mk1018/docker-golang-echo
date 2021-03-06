package main

import (
	"net/http"

	"github.com/jinzhu/gorm"

	"github.com/labstack/echo"
)

func main() {
	db, err := gorm.Open("mysql", "testuser:testpass@tcp(127.0.0.1:3306)/testdb?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
	}

	defer db.Close()
	db.Set("gorm:table_options", "ENGINE=InnoDB CHARSET=utf8").AutoMigrate(&model.User{})

	e := echo.New()
	defer e.Close()

	http.Handle("/", e)
	g := e.Group("/api/v1")

	controller.NewUser(db).Handle(g)

	appengine.Main()
}
