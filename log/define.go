package log

import (
	"bytes"
	"os"
	"sync"
	"time"
)

var level = []string{
	"[Debug]", //0
	"[Info]",  //1
	"[Warn]",  //2
	"[Error]", //3
	"[Panic]", //4
	"[Fatal]", //5
}

var colors = []int{
	36, //天蓝 0
	32, //绿 1
	93, //黄 2
	31, //红 3
	95, //紫 4
	34, //深蓝 5

}

type std struct {
	lk     *sync.RWMutex
	fd_run *os.File
	fd_log *os.File
	out    *os.File
	loglv  int //when lv >= loglv then write to file
	tm     time.Time
	color  int //0 no ,1 yes
	dev    int //1 window, 0 unix
	buf    *bytes.Buffer
	mode   int // 0 打印 1 打印+文件 2 文件 ,3 打印+文件+ 终端=>文件
}

var mylog std

func init() {
	mylog = std{
		lk:    new(sync.RWMutex),
		mode:  0,
		color: 0,
		loglv: 0,
		tm:    time.Now(),
		buf:   new(bytes.Buffer),
		out:   os.Stdout,
	}
}

const (
	Log_mode_Print         = 0
	Log_mode_Write         = 1
	Log_mode_PrintAndWrite = 2
	Log_mode_All           = 3

	Log_Lv_Debug = 0
	Log_Lv_Info  = 1
	Log_Lv_Warn  = 2
	Log_Lv_Error = 3
	Log_Lv_Panic = 4
	Log_Lv_Fatal = 5

	Windows = 1
	Unix    = 0

	Color_On  = 1
	Color_Off = 0
)

func LogInit(mode, loglv, dev, color int) {
	mylog.mode = mode
	mylog.color = color
	mylog.dev = dev
	mylog.loglv = loglv
	if dev == 1 && color == 1 {
		setWindowsColorInfo()
	}
	if mode > 0 {
		mylog.createlogfile()
	}
	Println("----------------------Log-Start----------------------")
}
func LogClose() {
	mylog.lk.RLock()
	defer mylog.lk.RUnlock()
	if mylog.fd_log != nil {
		mylog.fd_log.Close()
	}
	if mylog.fd_run != nil {
		mylog.fd_run.Close()
	}
	mylog.out.Close()
}

func setWindowsColorInfo() {
	colors = []int{
		3,  //天蓝 0
		2,  //绿 1
		14, //黄 2
		12, //红 3
		13, //紫 4
		1,  //深蓝 5
	}
}
