# go-crudify

Easily create CRUD endpoints with a given MUX and Database. The perfect package to prototype and build applications with.

## Example

```go
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
		crudify.WithRouter(routerAdapter.NewEchoRouter(e)),
		crudify.WithDatabase(databaseAdapter.NewGormAdapter(db, &User{})),
		crudify.WithModel(&User{}),
	); err != nil {
		log.Fatal(err)
	}

	e.Logger.Fatal(e.Start(":8000"))
}
```

This creates the following endpoints:

```
GET /users
POST /users
GET /users/:id
PUT /users/:id
DELETE /users/:id
```

## TODOs

- [ ] Setup go.mod
- [ ] Cleanup TODOs in the code
- [ ] Support non-integer IDs from gorm / database
- [ ] Introduce adapters for router & support standard lib mux
- [ ] Introduce adapters for database & support standard lib dbo
- [ ] Tests
- [ ] Permission behavior
- [ ] Instrumentation support
- [ ] Logging support
- [ ] Support for middlewares
- [ ] Improve documentation in code and docs
- [ ] Your feature! Submit an issue or PR
