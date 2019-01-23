package sql

import (
	"fmt"
	"github.com/CYL96/MySDK/log"
)

func (this *engine) selectBuild() {

}

func (this *engine) writeFrom() {
	this.query.WriteString(" FROM ")
	for i := range this.tableName {
		this.query.WriteString(this.tableName[i])
		if i != len(this.tableName)-1 {
			this.query.WriteString(",")
		}
	}
}

func (this *engine) writeJoin() {
	for _, v := range this.join {
		this.query.WriteString(" ")
		this.query.WriteString(v.joinName)
		this.query.WriteString(" ")
		this.query.WriteString(v.tableName)
		this.writeCondition("ON", v.condition)
	}
}

func (this *engine) writeCondition(op string, conditions []condition) {
	if len(conditions) == 0 {
		return
	}
	this.query.WriteString(" ")
	this.query.WriteString(op)
	for i, v_1 := range conditions {
		this.query.WriteString(" ")
		this.query.WriteString(v_1.condition)
		if i != 0 {
			if v_1.andor == "" {
				this.query.WriteString(" AND")
			} else {
				this.query.WriteString(v_1.andor)
			}
		}
		log.Println(this.args)
		if v_2, ok := v_1.value.([]interface{}); ok {
			this.args = append(this.args, v_2...)
		} else if v_1.value != nil {
			this.args = append(this.args, v_2)
		}
		log.Println(this.args)

	}
	return
}
func (this *engine) writeTail() {
	if len(this.groupBy) > 0 {
		this.query.WriteString(" GROUP BY ")
		this.query.WriteString(this.groupBy)
	}
	if len(this.orderBy) > 0 {
		this.query.WriteString(" ORDER BY ")
		this.query.WriteString(this.orderBy)
	}

}
func (this *engine) writePagination() {
	if this.offset > 0 {
		switch this.sqlType {
		case sqllite3, sqlserver, mysql:
			this.query.WriteString(" LIMIT")
			this.query.WriteString(fmt.Sprintf(" %d,%d", this.start, this.offset))

		}
	}
}

func (this *engine) GetSqlString() (string, []interface{}) {
	return this.query.String(), this.args

}
