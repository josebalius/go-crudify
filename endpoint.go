package crudify

import (
	"reflect"

	"github.com/pkg/errors"
)

type endpoint struct {
	options        *options
	controllerName string
}

func NewEndpoint(opts ...Option) error {
	e := &endpoint{
		options: &options{},
	}

	for _, o := range opts {
		if err := o(e.options); err != nil {
			return errors.Wrap(err, "apply option")
		}
	}

	if err := e.validateOptions(); err != nil {
		return errors.Wrap(err, "validate options")
	}

	if e.options.controllerName != "" {
		e.controllerName = e.options.controllerName
	} else {
		e.controllerName = e.database().TableName()
	}

	e.createRouterEndpoints()

	return nil
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
