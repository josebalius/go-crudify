package crudify

import (
	"reflect"

	"github.com/josebalius/go-crudify/adapters/database"
)

func (e *endpoint) database() database.Database {
	return e.options.db
}

func (e *endpoint) model() interface{} {
	return e.options.model
}

func (e *endpoint) modelType() reflect.Type {
	return reflect.TypeOf(e.model())
}

func (e *endpoint) modelSlice() reflect.Value {
	return reflect.MakeSlice(reflect.SliceOf(e.modelType()), 0, 0)
}

func (e *endpoint) newModelPtr() interface{} {
	model := reflect.New(e.modelType())
	return model.Interface()
}

func (e *endpoint) newModelSlice() interface{} {
	modelSlice := e.modelSlice()
	slice := reflect.New(modelSlice.Type())
	slice.Elem().Set(modelSlice)
	return slice.Interface()
}
