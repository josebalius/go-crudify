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

func (e *endpoint) router() router.Router {
	return e.options.router
}

func (e *endpoint) createRouterEndpoints() {
	endpointPath := fmt.Sprintf("/%s", strings.TrimPrefix(e.controllerName, "/"))

	e.createRouterListEndpoint(endpointPath)
	e.createRouterPostEndpoint(endpointPath)
	e.createRouterGetEndpoint(endpointPath)
	e.createRouterPutEndpoint(endpointPath)
	e.createRouterDeleteEndpoint(endpointPath)
}

func (e *endpoint) createRouterListEndpoint(endpointPath string) {
	fmt.Println("Creating endpoint: GET " + endpointPath)

	handler := func(ctx router.RouteContext) error {
		records := e.newModelSlice()

		if err := e.database().Find(records).Error; err != nil {
			return errors.Wrapf(err, "list %s", e.controllerName)
		}

		return ctx.JSON(http.StatusOK, records)
	}

	e.router().GET(endpointPath, handler)
}

func (e *endpoint) createRouterPostEndpoint(endpointPath string) {
	fmt.Println("Creating endpoint: POST " + endpointPath)

	handler := func(ctx router.RouteContext) error {
		record := e.newModelPtr()
		if err := ctx.Bind(record); err != nil {
			return errors.Wrap(err, "bind json payload to model")
		}

		if err := e.database().Create(record).Error; err != nil {
			return errors.Wrapf(err, "create %s", e.controllerName)
		}

		return ctx.JSON(http.StatusCreated, record)
	}

	e.router().POST(endpointPath, handler)
}

func (e *endpoint) createRouterGetEndpoint(endpointPath string) {
	endpointRoute := path.Join(endpointPath, ":id")
	fmt.Println("Creating endpoint: GET " + endpointRoute)

	handler := func(ctx router.RouteContext) error {
		record := e.newModelPtr()

		if err := e.database().First(record, ctx.Param("id")).Error; err != nil {
			if err == gorm.ErrRecordNotFound {
				return ctx.NoContent(http.StatusNotFound)
			}

			return errors.Wrapf(err, "get %s", e.controllerName)
		}

		return ctx.JSON(http.StatusOK, record)
	}

	e.router().GET(endpointRoute, handler)
}

func (e *endpoint) createRouterPutEndpoint(endpointPath string) {
	endpointRoute := path.Join(endpointPath, ":id")
	fmt.Println("Creating endpoint: PUT " + endpointRoute)

	handler := func(ctx router.RouteContext) error {
		record := e.newModelPtr()
		if err := ctx.Bind(record); err != nil {
			return errors.Wrap(err, "bind json payload to model")
		}

		recordMap := structs.Map(record)

		// Delete any gorm related fields
		delete(recordMap, "Model")

		if err := e.database().Model(e.model()).Where("id = ?", ctx.Param("id")).Update(recordMap).Error; err != nil {
			return errors.Wrapf(err, "update %s", e.controllerName)
		}

		updatedRecord := e.newModelPtr()
		if err := e.database().First(updatedRecord, ctx.Param("id")).Error; err != nil {
			return errors.Wrapf(err, "get %s", e.controllerName)
		}

		return ctx.JSON(http.StatusOK, updatedRecord)
	}

	e.router().PUT(endpointRoute, handler)
}

func (e *endpoint) createRouterDeleteEndpoint(endpointPath string) {
	endpointRoute := path.Join(endpointPath, ":id")
	fmt.Println("Creating endpoint: DELETE " + endpointRoute)

	handler := func(ctx router.RouteContext) error {
		if err := e.database().Where("id = ?", ctx.Param("id")).Delete(e.model()).Error; err != nil {
			return errors.Wrapf(err, "delete %s", e.controllerName)
		}

		return ctx.NoContent(http.StatusOK)
	}

	e.router().DELETE(endpointRoute, handler)
}
