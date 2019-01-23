package sql

import (
	"database/sql"
	"runtime"
	"sync"
	"time"

	"errors"
	_ "github.com/alexbrainman/odbc"
	_ "github.com/bennof/accessDBwE"
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/mattn/go-adodb"
)

func SqlInit(driver, path string, sqlType int) (*MyDB, error) {
	db, err := sql.Open(driver, path)
	if err != nil {
		return nil, err
	}
	err = db.Ping()
	if err != nil {
		return nil, err
	}
	db.SetConnMaxLifetime(time.Second * 10)
	db.SetMaxIdleConns(1000)
	db.SetMaxOpenConns(1000)
	s := new(MyDB)
	s.lk = new(sync.RWMutex)
	s.db = db
	s.path = path
	s.delay = time.Now().Unix()
	s.sqlType = sqlType

	return s, nil
}

//sql server
func SqlInitOLEDB(provider, ip, port, database, user_id, pwd string, window bool) (*MyDB, error) {
	var (
		driver string
		path   string
	)
	if runtime.GOOS == "windows" {
		driver = "adodb"
		if provider == "" {
			path = "Provider=SQLOLEDB"
		} else {
			path = "Provider=" + provider
		}
		if window {
			path += ";Integrated Security=SSPI"
		} else {
			path += ";user id=" + user_id + ";password=" + pwd
		}
		path += ";Data Source=" + ip + ";Initial Catalog=" + database + ";port=" + port
		//path = "Provider=SQLOLEDB;Data Source=192.168.1.237;Initial Catalog="+tableName+";"
		//path += "user id=sa3;password=zt666666;port=1433"
	} else {
		driver = "odbc"
		if provider == "" {
			path = "driver=ODBC Driver 17 for SQL Server"
		} else {
			path = "driver=" + provider
		}
		if window {
			path += ";Integrated Security=SSPI"
		} else {
			path += ";UID=" + user_id + ";PWD=" + pwd
		}
		path = ";SERVER=" + ip + ";PORT=" + port + ";DATABASE=" + database
	}
	return SqlInit(driver, path, access)
}

//ACCESS
func SqlInitACCESS(provider, ip, port, path, database, user_id, pwd string, window bool) (*MyDB, error) {
	//path := `Provider=;Data Source=` + util.WBConfig["sqlPath"].(string) + `;Jet OLEDB:Database Password=` + util.WBConfig["sqlPassword"].(string) + ";"
	driver := "adodb"
	if provider == "" {
		path = "Microsoft.ACE.OLEDB.12.0"
	} else {
		path = "driver=" + provider
	}
	if len(path) == 0 {
		return nil, errors.New("path is empty")
	}
	path += "Data Source=" + path + ";Jet OLEDB:Database Password=" + pwd + ";"
	return SqlInit(driver, path, access)
}

//Sqlite3
func SqlInitSqlite3(path string) (*MyDB, error) {
	if len(path) == 0 {
		return nil, errors.New("path is empty")
	}
	return SqlInit("sqlite3", path, sqllite3)
}

//Mysql
func SqlInitMysql(login, pwd, ip, port, database, charset string) (*MyDB, error) {
	path := login + ":" + pwd + "@tcp(" + ip + ":" + port + ")/" + database
	if len(charset) == 0 {
		path += "?charset=utf8"
	} else {
		path += "?charset=" + charset
	}
	return SqlInit("mysql", path, mysql)
}
