package http

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/josebalius/go-crudify/adapters/router"
	"github.com/pkg/errors"
)

type context struct {
	responseWriter http.ResponseWriter
	request        *http.Request
	resourceID     string
}

func newContext(w http.ResponseWriter, req *http.Request, resourceID string) router.RouteContext {
	return &context{w, req, resourceID}
}

func (c *context) Bind(payload interface{}) error {
	b, err := ioutil.ReadAll(c.request.Body)
	if err != nil {
		return errors.Wrap(err, "read request body")
	}

	if err := json.Unmarshal(b, payload); err != nil {
		return errors.Wrap(err, "unmarshal json")
	}

	return nil
}

func (c *context) ResourceID() string {
	return c.resourceID
}

func (c *context) JSON(status int, response interface{}) error {
	c.responseWriter.Header().Add("Content-Type", "application/json")
	c.responseWriter.WriteHeader(status)
	b, err := json.Marshal(response)
	if err != nil {
		return errors.Wrap(err, "marshal json")
	}
	c.responseWriter.Write(b)
	return nil
}

func (c *context) NoContent(status int) error {
	c.responseWriter.WriteHeader(status)
	return nil
}
