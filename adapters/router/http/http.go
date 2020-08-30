package http

import (
	"net/http"
	"path"
	"strings"

	"github.com/josebalius/go-crudify/adapters/router"
)

type httpRouter struct {
	mux          *http.ServeMux
	endpointPath string
	handlers     map[string]map[string]router.RouteHandler
}

func NewHTTP(mux *http.ServeMux) router.Router {
	router := &httpRouter{
		mux:      mux,
		handlers: make(map[string]map[string]router.RouteHandler),
	}
	return router
}

func (h *httpRouter) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	if handler, resourceID := h.findHandler(req.Method, req.URL.Path); handler != nil {
		handler(newContext(w, req, resourceID))
	}
}

func (h *httpRouter) findHandler(method, requestPath string) (router.RouteHandler, string) {
	if handlers, ok := h.handlers[method]; ok {
		if handler, exists := handlers[requestPath]; exists {
			return handler, ""
		}

		// no exact match, is there an id in the url?
		parts := strings.Split(requestPath, "/")
		if len(parts) > 0 {
			pathWithIDParam := path.Join(strings.Join(parts[:len(parts)-1], "/"), ":id")
			if handler, exists := handlers[pathWithIDParam]; exists {
				return handler, parts[len(parts)-1]
			}
		}
	}

	return nil, ""
}

func (h *httpRouter) WithEndpointPath(endpointPath string) {
	h.endpointPath = endpointPath
	h.mux.Handle(endpointPath, h)
	h.mux.Handle(endpointPath+"/", h)
}

func (h *httpRouter) registerHandler(method, path string, handler router.RouteHandler) {
	if _, ok := h.handlers[method]; !ok {
		h.handlers[method] = make(map[string]router.RouteHandler)
	}

	h.handlers[method][path] = handler
}

func (h *httpRouter) GET(path string, handler router.RouteHandler) {
	h.registerHandler("GET", path, handler)
}

func (h *httpRouter) POST(path string, handler router.RouteHandler) {
	h.registerHandler("POST", path, handler)
}

func (h *httpRouter) PUT(path string, handler router.RouteHandler) {
	h.registerHandler("PUT", path, handler)
}

func (h *httpRouter) DELETE(path string, handler router.RouteHandler) {
	h.registerHandler("DELETE", path, handler)
}
