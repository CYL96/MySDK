package sql

type WhereInfo struct {
	engine
}
type Where interface {
	//LeftBracket()Where
	//RightBracket()Where
	AND(params string, values ...interface{}) Where
	OR(params string, values ...interface{}) Where
	SET(andor string) Where
	IN(param string, values ...interface{}) Where
	NotIN(param string, values ...interface{}) Where

	LIKE(param string, value string) Where

	GroupBy(params string) Where
	OrderBy(params string) Where
	Pagination(start, offset int64) Where
	Select() *SelectInfo
	COUNT() (int64, error)
}

func (this *engine) Where(params string, values ...interface{}) Where {
	this.condition = append(this.condition, condition{params, values, "AND"})
	return &WhereInfo{*this}
}

func (this *WhereInfo) AND(params string, values ...interface{}) Where {
	this.condition = append(this.condition, condition{condition: params, value: values, andor: "AND"})
	return this
}
func (this *WhereInfo) OR(params string, values ...interface{}) Where {
	this.condition = append(this.condition, condition{condition: params, value: values, andor: "OR"})
	return this
}
func (this *WhereInfo) SET(andor string) Where {
	this.condition[len(this.condition)-1].andor = andor
	return this
}

func (this *WhereInfo) IN(param string, values ...interface{}) Where {
	param += " IN("
	for i := range values {
		param += "?"
		if i != len(values)-1 {
			param += ","
		}
	}
	param += ")"
	this.condition = append(this.condition, condition{})
	return this
}
func (this *WhereInfo) NotIN(param string, values ...interface{}) Where {
	param += " NOT IN("
	for i := range values {
		param += "?"
		if i != len(values)-1 {
			param += ","
		}
	}
	param += ")"
	this.condition = append(this.condition, condition{})
	return this
}
func (this *WhereInfo) LIKE(param string, value string) Where {
	param += " LIKE ?"
	value = "%" + value + "%"
	this.condition = append(this.condition, condition{param, value, "AND"})
	return this
}

func (this *engine) GroupBy(params string) Where {
	this.groupBy = params
	return &WhereInfo{*this}
}
func (this *engine) OrderBy(params string) Where {
	this.orderBy = params
	return &WhereInfo{*this}
}
func (this *engine) Pagination(start, offset int64) Where {
	this.start = start
	this.offset = offset
	return &WhereInfo{*this}
}

//
//func (this *WhereInfo)LeftBracket()Where{
//	this.condition[len(this.condition)-1].left_bracket = true
//	return this
//}
//func (this *WhereInfo)RightBracket()Where{
//	this.condition[len(this.condition)-1].right_bracket = true
//	return this
//}
