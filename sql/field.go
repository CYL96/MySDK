package sql

type Field interface {
	Table(tableName, as string) Table
}

type Table interface {
	Where(params string, values ...interface{}) Where
	Table(tableName, as string) Table
	Join(name, table, as string) Join
	LeftJoin(table, as string) Join
	RightJoin(table, as string) Join
	InnerJoin(table, as string) Join
	Select() *SelectInfo
	COUNT() (int64, error)
}

func (this *engine) Field(filed ...string) Field {
	this.filed = append(this.filed, filed...)
	return this
}

func (this *engine) Table(tableName, as string) Table {
	this.newAllocate()
	if len(as) > 0 {
		tableName += " AS " + as
	}
	this.tableName = append(this.tableName, tableName)
	return this
}
