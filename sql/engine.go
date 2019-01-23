package sql

import (
	"errors"
)

type Engine interface {
	GetEngine() *engine
	Field(filed ...string) Field
	Table(tableName, as string) Table
	Where(params string, values ...interface{}) Where
	Select() *SelectInfo
}

func (this *MyDB) NewEngine() (Engine, error) {
	if this.db == nil {
		return nil, errors.New("db is not open")
	}

	return &engine{db: this.db, sqlType: this.sqlType}, nil

}
func (this *engine) newAllocate() {
	this.query.Reset()
	this.start = 0
	this.offset = 0
	this.args = make([]interface{}, 0)
	this.orderBy = ""
	this.groupBy = ""
	this.condition = make([]condition, 0)
	this.tableName = make([]string, 0)
	this.join = make([]join, 0)
	this.filed = make([]string, 0)
}

func (this *engine) GetEngine() *engine {
	return this
}
