package sql

import "strconv"

type Select interface {
	GetSqlString() (string, []interface{})
	Get(tag interface{}) (interface{}, error)
	executeSelect() ([]map[string]interface{}, error)
}

type SelectInfo struct {
	engine
}

func (this *engine) Select() *SelectInfo {
	this.query.WriteString("SELECT ")
	if len(this.filed) == 0 {
		this.query.WriteString("*")
	} else {
		for i := range this.filed {
			this.query.WriteString(this.filed[i])
			if i != len(this.filed)-1 {
				this.query.WriteString(",")
			}
		}
	}
	this.writeFrom()
	this.writeJoin()
	this.writeCondition("WHERE", this.condition)
	this.writeTail()
	this.writePagination()
	return &SelectInfo{*this}
}

func (this *SelectInfo) GetMap() ([]map[string]interface{}, error) {
	return this.executeSelect()
}

func (this *SelectInfo) Get(tag interface{}) (interface{}, error) {
	data, err := this.executeSelect()
	if err != nil {
		return nil, err
	}
	return ParseStructSlice(tag, data)
}

func (this *engine) COUNT() (int64, error) {
	this.query.WriteString("SELECT count(*) as count")
	this.writeFrom()
	this.writeJoin()
	this.writeCondition("WHERE", this.condition)
	this.writeTail()
	this.writePagination()
	data, err := this.executeSelect()
	if err != nil {
		return 0, err
	}
	if len(data) == 0 {
		return 0, nil
	}
	data_2, err := interfaceToString(data)
	if err != nil {
		return 0, err
	}
	num, err := strconv.ParseInt(data_2[0]["count"], 10, 64)
	return num, err
}
