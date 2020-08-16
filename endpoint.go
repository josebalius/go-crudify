package crudify

import "github.com/pkg/errors"

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
		e.controllerName = e.database().NewScope(e.model()).TableName()
	}

	e.createRouterEndpoints()

	return nil
}
