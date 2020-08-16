package database

type Database interface {
	TableName() string
	Find(records interface{}) error
	Create(record interface{}) error
	First(record interface{}, id string) error
	Update(record map[string]interface{}, id string) error
	Delete(id string) error
}
