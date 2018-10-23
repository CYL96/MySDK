package util

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"
)

func ReadFile(path string) string {
	fmt.Println("path:", path)
	fi, err := os.Open(path)
	if err != nil {
		fmt.Println(err)
	}
	defer fi.Close()
	fd, err := ioutil.ReadAll(fi)
	fd = bytes.TrimPrefix(fd, []byte("\xef\xbb\xbf"))
	return string(fd)
}
func GetCurrentDirectory() string {
	if runtime.GOOS == "windows" {
		dir, err := filepath.Abs(filepath.Dir(os.Args[0])) //返回绝对路径  filepath.Dir(os.Args[0])去除最后一个元素的路径
		if err != nil {
			fmt.Println(err)
			return ""
		}
		//return strings.Replace(dir, "\\", "/", -1) //将\替换成/
		//fmt.Println(dir)
		return dir
	} else {
		return "./"
	}
}

func GetCurrentPath() string {
	file, err := exec.LookPath(os.Args[0])
	if err != nil {
		return ""
	}
	path, err := filepath.Abs(file)
	if err != nil {
		return ""
	}
	i := strings.LastIndex(path, "/")
	if i < 0 {
		i = strings.LastIndex(path, "\\")
	}
	if i < 0 {
		return ""
	}
	return string(path[0 : i+1])
}

func CurrentFile() string {
	_, file, _, ok := runtime.Caller(1)
	if !ok {
	}
	return strings.Replace(file, "/", "\\", -1)
}
