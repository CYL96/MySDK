package sql

import (
	"errors"
	"reflect"
	"strconv"
	"strings"
)

func AllocateSlice(example interface{}) (interface{}, error) {
	if nil == example {
		return nil, errors.New("example is nil")
	}
	t := reflect.TypeOf(example)
	if t.Kind() == reflect.Ptr {
		v := reflect.ValueOf(example)
		if v.IsNil() || !v.IsValid() {
			return nil, errors.New("point is nil")
		}
		iv := reflect.Indirect(v).Interface()
		t = reflect.TypeOf(iv)
	}

	newSlice := reflect.SliceOf(t)
	sli := reflect.MakeSlice(newSlice, 0, 0).Interface()
	return sli, nil
}

func NewStructPoint(val interface{}) (interface{}, error) {
	t := reflect.TypeOf(val)

	var rslt interface{}
	if reflect.Ptr == t.Kind() {
		v := reflect.ValueOf(val) //指针类型的值;
		if !v.IsValid() || v.IsNil() {
			return nil, errors.New("point is nil")
		}
		vv := reflect.Indirect(v).Interface() //指针指向的值;
		tt := reflect.TypeOf(vv)
		rslt = reflect.New(tt).Interface()
	} else {
		rslt = reflect.New(t).Interface()
	}
	return rslt, nil
}

func ParseStructOne(ptr interface{}, data map[string]string) (interface{}, error) {
	if reflect.TypeOf(ptr).Kind() != reflect.Ptr {
		return nil, errors.New("prt is not a point")
	}
	ptrNew, err := NewStructPoint(ptr)
	if err != nil {
		return nil, err
	}
	v := reflect.ValueOf(ptrNew).Elem()
	t := reflect.TypeOf(ptrNew).Elem()
	if t.Kind() == reflect.Struct {
		for i := 0; i < t.NumField(); i++ {
			t_1 := t.Field(i)
			v_1 := v.Field(i)
			tags, ok := t_1.Tag.Lookup("json")
			if !ok {
				continue
			}
			tag := strings.Split(tags, ",")[0]
			tmp := data[tag]
			switch t_1.Type.Kind() {
			case reflect.String:
				v_1.SetString(tmp)
			case reflect.Int, reflect.Int32, reflect.Int64:
				tmp_1, err := strconv.ParseInt(tmp, 10, 64)
				if err != nil {
					return nil, errors.New("Parse field: [" + tag + "] err:" + err.Error())
				}
				v_1.SetInt(tmp_1)
			case reflect.Float64, reflect.Float32:
				tmp_1, err := strconv.ParseFloat(tmp, 64)
				if err != nil {
					return nil, errors.New("Parse field: [" + tag + "] err:" + err.Error())
				}
				v_1.SetFloat(tmp_1)
			default:
				return nil, errors.New("not support type:" + t_1.Type.Kind().String())
			}
		}
	}
	return reflect.Indirect(reflect.ValueOf(ptrNew)).Interface(), nil
}
func ParseStructSlice(example interface{}, data []map[string]interface{}) (interface{}, error) {
	result, err := AllocateSlice(example)
	if err != nil {
		return nil, err
	}
	data2, err := interfaceToString(data)
	if err != nil {
		return nil, err
	}
	for _, row := range data2 {
		v := reflect.ValueOf(result)
		value, err := ParseStructOne(example, row)
		if err != nil {
			return nil, err
		}
		t := reflect.TypeOf(value)
		if t.Kind() == reflect.Ptr {
			v2 := reflect.ValueOf(value)
			vv2 := reflect.Indirect(v2)
			result = reflect.Append(v, reflect.ValueOf(vv2.Interface())).Interface()
		} else {
			result = reflect.Append(v, reflect.ValueOf(value)).Interface()
		}
	}
	return result, nil
}
