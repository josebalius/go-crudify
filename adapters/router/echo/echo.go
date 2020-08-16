package echo

import (
	"github.com/josebalius/go-crudify/adapters/router"
	"github.com/labstack/echo"
)

type echoRouter struct {
	echo *echo.Echo
}

func NewEcho(e *echo.Echo) router.Router {
	return &echoRouter{e}
}

func (e *echoRouter) GET(path string, handler router.RouteHandler) {
	e.echo.GET(path, func(ctx echo.Context) error {
		return handler(ctx)
	})
}

func (e *echoRouter) POST(path string, handler router.RouteHandler) {
	e.echo.POST(path, func(ctx echo.Context) error {
		return handler(ctx)
	})
}

func (e *echoRouter) PUT(path string, handler router.RouteHandler) {
	e.echo.PUT(path, func(ctx echo.Context) error {
		return handler(ctx)
	})
}

func (e *echoRouter) DELETE(path string, handler router.RouteHandler) {
	e.echo.DELETE(path, func(ctx echo.Context) error {
		return handler(ctx)
	})
}
