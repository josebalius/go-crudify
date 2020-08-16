package crudify

import (
	"reflect"

	"github.com/jinzhu/gorm"
	"github.com/labstack/echo"
	"github.com/pkg/errors"
)

type Option func(opts *options) error

type options struct {
	router         *echo.Echo
	db             *gorm.DB
	model          interface{}
	controllerName string
}

func WithRouter(router *echo.Echo) Option {
	return func(opts *options) error {
		opts.router = router
		return nil
	}
}

func WithDatabase(db *gorm.DB) Option {
	return func(opts *options) error {
		opts.db = db
		return nil
	}
}

func WithModel(model interface{}) Option {
	return func(opts *options) error {
		opts.model = model
		return nil
	}
}

func WithControllerName(controllerName string) Option {
	return func(opts *options) error {
		opts.controllerName = controllerName
		return nil
	}
}

func (e *endpoint) validateOptions() error {
	switch {
	case e.options.router == nil:
		return errors.New("A router is required to create an endpoint")
	case e.options.db == nil:
		return errors.New("A database is required to create an endpoint")
	case e.options.model == nil:
		return errors.New("A model is required to create an endpoint")
	case e.options.model != nil:
		// TODO: move this out of validation
		model := reflect.ValueOf(e.options.model)
		if model.Kind() == reflect.Ptr {
			e.options.model = model.Elem().Interface()
		}
	}

	return nil
}
