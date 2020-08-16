package crudify

import (
	"github.com/jinzhu/gorm"
	"github.com/josebalius/go-crudify/adapters/router"
)

type Option func(opts *options) error

type options struct {
	router         router.Router
	db             *gorm.DB
	model          interface{}
	controllerName string
}

func WithRouter(router router.Router) Option {
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