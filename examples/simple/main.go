package main

import (
	"log"

	"github.com/jinzhu/gorm"
	"github.com/josebalius/go-crudify"
	databaseAdapter "github.com/josebalius/go-crudify/adapters/database/gorm"
	routerAdapter "github.com/josebalius/go-crudify/adapters/router/echo"
	"github.com/labstack/echo"
	"github.com/pkg/errors"

	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

type User struct {
	gorm.Model
	Name string
}

func main() {
	e := echo.New()

	db, err := gorm.Open("sqlite3", "test.db")
	if err != nil {
		log.Fatal(errors.Wrap(err, "open database"))
	}
	defer db.Close()

	db.AutoMigrate(&User{})

	if err := crudify.NewEndpoint(
		crudify.WithRouter(routerAdapter.NewEcho(e)),
		crudify.WithDatabase(databaseAdapter.NewGorm(db)),
		crudify.WithModel(&User{}),
	); err != nil {
		log.Fatal(err)
	}

	e.Logger.Fatal(e.Start(":8000"))
}
