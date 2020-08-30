package router

type Router interface {
	WithEndpointPath(path string)

	GET(path string, handler RouteHandler)
	POST(path string, handler RouteHandler)
	PUT(path string, handler RouteHandler)
	DELETE(path string, handler RouteHandler)
}

type RouteHandler func(ctx RouteContext) error

type RouteContext interface {
	Bind(payload interface{}) error
	ResourceID() string
	JSON(status int, response interface{}) error
	NoContent(status int) error
}
