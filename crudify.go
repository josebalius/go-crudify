package crudify

import (
	"github.com/jinzhu/gorm"
	"github.com/labstack/echo"
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
