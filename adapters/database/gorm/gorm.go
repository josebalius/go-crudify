package gorm

import (
	"github.com/jinzhu/gorm"
	"github.com/josebalius/go-crudify/adapters/database"
)

type gormAdapter struct {
	db    *gorm.DB
	model interface{}
}

func NewGorm(db *gorm.DB) database.Database {
	return &gormAdapter{db: db}
}

func (g *gormAdapter) WithModel(model interface{}) {
	g.model = model
}

func (g *gormAdapter) TableName() string {
	return g.db.NewScope(g.model).TableName()
}

func (g *gormAdapter) Find(records interface{}) error {
	return g.db.Find(records).Error
}

func (g *gormAdapter) Create(record interface{}) error {
	return g.db.Create(record).Error
}

func (g *gormAdapter) First(record interface{}, id string) error {
	return g.db.First(record, id).Error
}

func (g *gormAdapter) Update(record map[string]interface{}, id string) error {
	return g.db.Model(g.model).Where("id = ?", id).Update(record).Error
}

func (g *gormAdapter) Delete(id string) error {
	return g.db.Where("id = ?", id).Delete(g.model).Error
}
