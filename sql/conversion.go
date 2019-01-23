package sql

import (
	"errors"
	"reflect"
	"strconv"
)

func interfaceToString(data []map[string]interface{}) ([]map[string]string, error) {
	data_s := make([]map[string]string, 0)
	ref := make(map[string]reflect.Kind)
	for _, v := range data {
		tmp := make(map[string]string)
		for j, v2 := range v {
			t, ok := ref[j]
			if !ok {
				t = reflect.TypeOf(v2).Kind()
				ref[j] = t
			}
			switch t {
			case reflect.String:
				tmp[j] = v2.(string)
			case reflect.Int:
				tmp[j] = strconv.Itoa(v2.(int))
			case reflect.Int64:
				tmp[j] = strconv.FormatInt(v2.(int64), 10)
			case reflect.Int32:
				tmp[j] = strconv.FormatInt(int64(v2.(int32)), 10)
			case reflect.Float64:
				tmp[j] = strconv.FormatFloat(v2.(float64), 'f', -1, 64)
			case reflect.Float32:
				tmp[j] = strconv.FormatFloat(float64(v2.(float32)), 'f', -1, 32)
			default:
				if v3, ok := v2.([]byte); !ok {
					return nil, errors.New("not support type:" + t.String())
				} else {
					tmp[j] = string(v3)
				}

			}

		}
		data_s = append(data_s, tmp)
	}
	return data_s, nil
}

func stringToType(value string, kind reflect.Kind) (interface{}, error) {
	switch kind {
	case reflect.String:
		return value, nil
	case reflect.Int:
		return strconv.Atoi(value)
	case reflect.Int64:
		return strconv.ParseInt(value, 10, 64)
	case reflect.Float64:
		return strconv.ParseFloat(value, 64)
	default:
		return nil, errors.New("not support type:" + kind.String())
	}
}
