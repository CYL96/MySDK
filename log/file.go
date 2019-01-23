package log

import (
	"os"
	"time"
)

//创建log文件
func (this *std) createlogfile() {
	t1 := time.Now().Format("2006_01_02")
	pathcreate("./log")
	path_run := "./log/" + t1 + "_Run.txt"
	path_log := "./log/" + t1 + "_Log.txt"
	logfile_run, err := os.OpenFile(path_run, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		LogError("path_run: %s\r\n", err.Error())
		os.Exit(-1)
	}
	logfile_log, err := os.OpenFile(path_log, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		LogError("path_run: %s\r\n", err.Error())
		os.Exit(-1)
	}
	this.fd_run = logfile_run
	this.fd_log = logfile_log
	if this.mode == Log_mode_All {
		os.Stdout = this.fd_run
		os.Stderr = this.fd_log
	}
	Println(path_run, path_log)
	this.fd_log.Sync()
	this.fd_run.Sync()
	return
}

//自动创建文件夹
func pathcreate(path string) error {
	var err error
	_, err = os.Stat(path)
	if err != nil {
		LogError(err)
		err = os.MkdirAll(path, 0777)
		if err != nil {
			LogError(err)
		}
	}
	return err
}
