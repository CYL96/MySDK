package util

import (
	"encoding/json"
	"errors"
	"github.com/CYL96/MySDK/log"
	"net/http"
	"reflect"
	"strconv"
)

func ParseRequestForm(req *http.Request, data interface{}) error {
	if reflect.TypeOf(data).Kind() != reflect.Ptr {
		return errors.New("Data is not a point.")
	}
	params := req.FormValue
	v := reflect.ValueOf(data).Elem()
	t := reflect.TypeOf(data).Elem()
	if t.Kind() == reflect.Struct {
		for i := 0; i < t.NumField(); i++ {
			t_1 := t.Field(i)
			v_1 := v.Field(i)
			var tag string
			var ok bool
			if tag, ok = t_1.Tag.Lookup("json"); !ok {
				continue
			}
			switch t_1.Type.Kind() {
			case reflect.String:
				v_1.SetString(params(tag))
			case reflect.Float64:
				tmp, err := strconv.ParseFloat(params(tag), 10)
				if err != nil {
					return errors.New(tag + " is a  not a Float64 type.")
				}
				v_1.SetFloat(tmp)
			case reflect.Int64, reflect.Int:
				tmp, err := strconv.ParseInt(params(tag), 10, 64)
				if err != nil {
					return errors.New(tag + " is a  not a Int or Int64 type.")
				}
				v_1.SetInt(tmp)
			case reflect.Struct, reflect.Slice:
				data := params(tag)
				err := json.Unmarshal([]byte(data), v_1.Addr().Interface())
				if err != nil {
					log.PrintError(err)
					return errors.New(tag + " is a  not a Struct or Slice type or nil point.")
				}
			default:
				return errors.New(tag + " is a  not a support type.")

			}
		}
	} else {
		return errors.New("Data is not a struct type.")
	}
	return nil
}
