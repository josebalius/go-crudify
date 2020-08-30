package crudify

import (
	"fmt"
	"net/http"
	"path"
	"strings"

	"github.com/fatih/structs"
	"github.com/jinzhu/gorm"
	"github.com/josebalius/go-crudify/adapters/router"
	"github.com/pkg/errors"
)

func (e *Endpoint) router() router.Router {
	return e.Options.Router
}

func (e *Endpoint) createRouterEndpoints() {
	endpointPath := fmt.Sprintf("/%s", strings.TrimPrefix(e.ControllerName, "/"))

	e.router().WithEndpointPath(endpointPath)

	e.createRouterListEndpoint(endpointPath)
	e.createRouterPostEndpoint(endpointPath)
	e.createRouterGetEndpoint(endpointPath)
	e.createRouterPutEndpoint(endpointPath)
	e.createRouterDeleteEndpoint(endpointPath)
}

func (e *Endpoint) createRouterListEndpoint(endpointPath string) {
	fmt.Println("Creating Endpoint: GET " + endpointPath)

	handler := func(ctx router.RouteContext) error {
		records := e.newModelSlice()

		if err := e.database().Find(records); err != nil {
			return errors.Wrapf(err, "list %s", e.ControllerName)
		}

		return ctx.JSON(http.StatusOK, records)
	}

	e.router().GET(endpointPath, handler)
}

func (e *Endpoint) createRouterPostEndpoint(endpointPath string) {
	fmt.Println("Creating Endpoint: POST " + endpointPath)

	handler := func(ctx router.RouteContext) error {
		record := e.newModelPtr()
		if err := ctx.Bind(record); err != nil {
			return errors.Wrap(err, "bind json payload to model")
		}

		if err := e.database().Create(record); err != nil {
			return errors.Wrapf(err, "create %s", e.ControllerName)
		}

		return ctx.JSON(http.StatusCreated, record)
	}

	e.router().POST(endpointPath, handler)
}

func (e *Endpoint) createRouterGetEndpoint(endpointPath string) {
	endpointRoute := path.Join(endpointPath, ":id")
	fmt.Println("Creating Endpoint: GET " + endpointRoute)

	handler := func(ctx router.RouteContext) error {
		record := e.newModelPtr()

		if err := e.database().First(record, ctx.ResourceID()); err != nil {
			if err == gorm.ErrRecordNotFound {
				return ctx.NoContent(http.StatusNotFound)
			}

			return errors.Wrapf(err, "get %s", e.ControllerName)
		}

		return ctx.JSON(http.StatusOK, record)
	}

	e.router().GET(endpointRoute, handler)
}

func (e *Endpoint) createRouterPutEndpoint(endpointPath string) {
	endpointRoute := path.Join(endpointPath, ":id")
	fmt.Println("Creating Endpoint: PUT " + endpointRoute)

	handler := func(ctx router.RouteContext) error {
		record := e.newModelPtr()
		if err := ctx.Bind(record); err != nil {
			return errors.Wrap(err, "bind json payload to model")
		}

		recordMap := structs.Map(record)

		// Delete any gorm related fields
		delete(recordMap, "Model")

		if err := e.database().Update(recordMap, ctx.ResourceID()); err != nil {
			return errors.Wrapf(err, "update %s", e.ControllerName)
		}

		updatedRecord := e.newModelPtr()
		if err := e.database().First(updatedRecord, ctx.ResourceID()); err != nil {
			return errors.Wrapf(err, "get %s", e.ControllerName)
		}

		return ctx.JSON(http.StatusOK, updatedRecord)
	}

	e.router().PUT(endpointRoute, handler)
}

func (e *Endpoint) createRouterDeleteEndpoint(endpointPath string) {
	endpointRoute := path.Join(endpointPath, ":id")
	fmt.Println("Creating Endpoint: DELETE " + endpointRoute)

	handler := func(ctx router.RouteContext) error {
		if err := e.database().Delete(ctx.ResourceID()); err != nil {
			return errors.Wrapf(err, "delete %s", e.ControllerName)
		}

		return ctx.NoContent(http.StatusOK)
	}

	e.router().DELETE(endpointRoute, handler)
}
