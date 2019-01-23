package main

import (
	"github.com/CYL96/MySDK/log"
	"reflect"
)

func main() {
	log.LogInit(log.Log_mode_Print, log.Log_Lv_Debug, log.Unix, log.Color_On)
	var cc cccc

	t_1 := reflect.TypeOf(cc)
	log.Println(t_1)
	t_2 := reflect.SliceOf(t_1)
	log.Println(t_2)
	v_1 := reflect.MakeSlice(t_2, 0, 0)
	log.Println(v_1)

	v_2 := reflect.New(t_1)
	a := reflect.Append(v_1, v_2.Elem()).Interface()
	log.Println(a.([]cccc))

}

type cccc struct {
	A string `json:"a"`
	B string `json:"b"`
}
