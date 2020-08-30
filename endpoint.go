package crudify

import (
	"reflect"

	"github.com/pkg/errors"
)

type Endpoint struct {
	Options        *Options
	ControllerName string
}

func NewEndpoint(opts ...Option) error {
	e := &Endpoint{
		Options: &Options{},
	}

	for _, o := range opts {
		if err := o(e.Options); err != nil {
			return errors.Wrap(err, "apply option")
		}
	}

	if err := e.validateOptions(); err != nil {
		return errors.Wrap(err, "validate options")
	}

	e.database().WithModel(e.Options.Model)

	if e.Options.ControllerName != "" {
		e.ControllerName = e.Options.ControllerName
	} else {
		e.ControllerName = e.database().TableName()
	}

	e.createRouterEndpoints()

	return nil
}

func (e *Endpoint) validateOptions() error {
	switch {
	case e.Options.Router == nil:
		return errors.New("A router is required to create an endpoint")
	case e.Options.DB == nil:
		return errors.New("A database is required to create an endpoint")
	case e.Options.Model == nil:
		return errors.New("A model is required to create an endpoint")
	case e.Options.Model != nil:
		// TODO: move this out of validation
		model := reflect.ValueOf(e.Options.Model)
		if model.Kind() == reflect.Ptr {
			e.Options.Model = model.Elem().Interface()
		}
	}

	return nil
}
