package sql

import (
	"bytes"
	"database/sql"
	"sync"
)

const (
	mysql int = iota
	sqlserver
	access
	sqllite3
)

type MyDB struct {
	db      *sql.DB
	lk      *sync.RWMutex
	path    string
	delay   int64
	sqlType int
}
type engine struct {
	db      *sql.DB
	sqlType int
	result  interface{}

	filed     []string
	tableName []string
	condition []condition

	join []join

	groupBy string
	orderBy string

	start  int64
	offset int64

	query bytes.Buffer
	args  []interface{}
}
type join struct {
	tableName string
	joinName  string
	condition []condition
}
type condition struct {
	condition string
	value     interface{}
	andor     string
}
