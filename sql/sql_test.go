package sql

import (
	"github.com/CYL96/MySDK/log"
	"testing"
)

type TTTT struct {
	AAA string `json:"aaa"`
	BBB string `json:"bbb"`
}

func TestSqlInitOLEDB(t *testing.T) {
	db, err := SqlInitMysql("root", "853051095", "127.0.0.1", "3306", "mytest", "")
	if err != nil {
		log.LogError(err)
		return
	}
	eg, err := db.NewEngine()
	if err != nil {
		log.Println(err)
		return
	}
	row, err := db.db.Query("select * from my_user_info where id LIKE ? ", 1)
	if err != nil {
		log.LogError(err)
		return
	}
	log.Println(row.Columns())
	log.LogError(eg.Table("my_user_info", "").COUNT())

	sel1 := eg.Table("my_user_info", "A").Select()
	log.Println(sel1.GetSqlString())
	var a user_info
	b, _ := sel1.Get(&a)
	c := b.([]user_info)
	log.Println(c)
}

type user_info struct {
	UserName string `json:"username"`
	Id       string `json:"id"`
	Position int64  `json:"position"`
}
