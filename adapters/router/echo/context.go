package echo

import (
	"github.com/josebalius/go-crudify/adapters/router"
	"github.com/labstack/echo"
)

type context struct {
	echo.Context
}

func newContext(ctx echo.Context) router.RouteContext {
	return &context{ctx}
}

func (c *context) ResourceID() string {
	return c.Param("id")
}
