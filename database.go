package crudify

import (
	"reflect"

	"github.com/josebalius/go-crudify/adapters/database"
)

func (e *Endpoint) database() database.Database {
	return e.Options.DB
}

func (e *Endpoint) model() interface{} {
	return e.Options.Model
}

func (e *Endpoint) modelType() reflect.Type {
	return reflect.TypeOf(e.model())
}

func (e *Endpoint) modelSlice() reflect.Value {
	return reflect.MakeSlice(reflect.SliceOf(e.modelType()), 0, 0)
}

func (e *Endpoint) newModelPtr() interface{} {
	model := reflect.New(e.modelType())
	return model.Interface()
}

func (e *Endpoint) newModelSlice() interface{} {
	modelSlice := e.modelSlice()
	slice := reflect.New(modelSlice.Type())
	slice.Elem().Set(modelSlice)
	return slice.Interface()
}
