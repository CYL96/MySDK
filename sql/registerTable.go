package sql

import "errors"

const (
	Engine_InnoDB   = "InnoDB"
	Charset_utf8mb4 = "utf8mb4"
)

type TableInfo struct {
	TableName  string
	Engine     string
	Charset    string
	Comment    string
	PrimaryKey []string
	UniqueKey  map[string][]string
	IndexKey   map[string][]string
	Table      interface{}
}

func RigisterTable(tableName, engine, charset, comment string, table interface{}) TableInfo {
	var db TableInfo
	db.TableName = tableName
	db.Engine = engine
	db.Charset = charset
	db.Comment = comment
	db.Table = table
	db.IndexKey = make(map[string][]string)
	db.IndexKey = make(map[string][]string)
	return db
}

func (this *TableInfo) AddUniqueKey(name string, keys ...string) error {
	if len(name) == 0 || len(keys) == 0 {
		return errors.New("name or keys is empty")
	}
	this.UniqueKey[name] = append(this.UniqueKey[name], keys...)
	return nil
}

func (this *TableInfo) AddIndexKey(name string, keys ...string) error {
	if len(name) == 0 || len(keys) == 0 {
		return errors.New("name or keys is empty")
	}
	this.UniqueKey[name] = append(this.UniqueKey[name], keys...)
	return nil
}
