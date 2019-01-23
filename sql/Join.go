package sql

type Join interface {
	Join(name, table, as string) Join
	LeftJoin(table, as string) Join
	RightJoin(table, as string) Join
	InnerJoin(table, as string) Join
	OnSuf(suf string) Join
	ON(params string, value ...interface{}) Join
	AND(param string, value ...interface{}) Join
	OR(param string, value ...interface{}) Join
	//In_AND(param string, value ...interface{})Join
	//In_OR(param string, value ...interface{})Join
	//LeftBracket()Join
	//RightBracket()Join
	Where(params string, values ...interface{}) Where
	COUNT() (int64, error)
	GroupBy(params string) Where
	OrderBy(params string) Where
	Pagination(start, offset int64) Where
}

type JoinInfo struct {
	engine
}

//func (this *Join)Join(name string,table string)Join{
//	this.join=append(this.join,join{tableName:table,joinName:name})
//	return this
//}
func (this *engine) Join(name string, table string, as string) Join {
	if len(as) > 0 {
		table += " AS " + as
	}
	this.join = append(this.join, join{tableName: table, joinName: name})
	return &JoinInfo{*this}
}
func (this *engine) LeftJoin(table, as string) Join {
	if len(as) > 0 {
		table += " AS " + as
	}
	this.join = append(this.join, join{tableName: table, joinName: "LEFT JOIN"})
	return &JoinInfo{*this}
}
func (this *engine) RightJoin(table, as string) Join {
	if len(as) > 0 {
		table += " AS " + as
	}
	this.join = append(this.join, join{tableName: table, joinName: "RIGHT JOIN"})
	return &JoinInfo{*this}
}
func (this *engine) InnerJoin(table, as string) Join {
	if len(as) > 0 {
		table += " AS " + as
	}
	this.join = append(this.join, join{tableName: table, joinName: "INNER JOIN"})
	return &JoinInfo{*this}
}

//use On("A.id","=",1) for set value
//use On("A.id = B.id") for set condition
func (this *JoinInfo) ON(params string, value ...interface{}) Join {
	this.join[len(this.join)-1].condition = append(this.join[len(this.join)-1].condition, condition{condition: params, value: value, andor: "AND"})
	return this
}
func (this *JoinInfo) AND(params string, value ...interface{}) Join {
	this.join[len(this.join)-1].condition = append(this.join[len(this.join)-1].condition, condition{condition: params, value: value, andor: "AND"})
	return this
}
func (this *JoinInfo) OR(param string, value ...interface{}) Join {
	this.join[len(this.join)-1].condition = append(this.join[len(this.join)-1].condition, condition{condition: param, value: value, andor: "OR"})
	return this
}

//func (this *JoinInfo)In_AND(param string, value ...interface{})Join{
//	this.join[len(this.join)-1].condition = append(this.join[len(this.join)-1].condition,condition{condition:param,value:value,operation:"IN",andor:"AND"})
//	return this
//}
//func (this *JoinInfo)In_OR(param string, value ...interface{})Join{
//	this.join[len(this.join)-1].condition = append(this.join[len(this.join)-1].condition,condition{condition:param,value:value,operation:"IN",andor:"OR"})
//	return this
//}
func (this *JoinInfo) OnSuf(suf string) Join {
	this.join[len(this.join)-1].condition = append(this.join[len(this.join)-1].condition, condition{condition: suf, andor: "AND"})
	return this
}

//
//func (this *JoinInfo)LeftBracket()Join{
//	//this := this2.GetEngine()
//	this.join[len(this.join)-1].condition[len(this.join[len(this.join)-1].condition)-1].left_bracket = true
//	return this
//}
//
//func (this *JoinInfo)RightBracket()Join{
//	//this := this2.GetEngine()
//	this.join[len(this.join)-1].condition[len(this.join[len(this.join)-1].condition)-1].right_bracket = true
//	return this
//}

func (this *JoinInfo) SET(andor string) Join {
	//this := this2.GetEngine()
	this.join[len(this.join)-1].condition[len(this.join[len(this.join)-1].condition)-1].andor = andor
	return this
}
