package crudify

import (
	"github.com/josebalius/go-crudify/adapters/database"
	"github.com/josebalius/go-crudify/adapters/router"
)

type Option func(opts *Options) error

type Options struct {
	Router         router.Router
	DB             database.Database
	Model          interface{}
	ControllerName string
}

func WithRouter(router router.Router) Option {
	return func(opts *Options) error {
		opts.Router = router
		return nil
	}
}

func WithDatabase(db database.Database) Option {
	return func(opts *Options) error {
		opts.DB = db
		return nil
	}
}

func WithModel(model interface{}) Option {
	return func(opts *Options) error {
		opts.Model = model
		return nil
	}
}

func WithControllerName(controllerName string) Option {
	return func(opts *Options) error {
		opts.ControllerName = controllerName
		return nil
	}
}
