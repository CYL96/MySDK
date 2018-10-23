package log

import (
	"bytes"
	"fmt"
	"os"
	"runtime"
	"runtime/debug"
	"strconv"
	"strings"
	"sync"
	"syscall"
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
	lk    *sync.RWMutex
	fd    *os.File
	out   *os.File
	loglv int //when lv >= loglv then write to file
	tm    time.Time
	color int //0 no ,1 yes
	dev   int //1 window, 0 unix
	buf   *bytes.Buffer
	mode  int // 0 打印 1 打印+文件 2 文件 ,3 打印+文件+ 终端=>文件
}

var mylog std = std{
	lk:    new(sync.RWMutex),
	mode:  0,
	color: 0,
	loglv: 0,
	tm:    time.Now(),
	buf:   new(bytes.Buffer),
	out:   os.Stdout,
}

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
}

func PrintDebug(v ...interface{}) {
	//fmt.Println(fmt.Sprintln())
	mylog.output(0, 1, fmt.Sprintln(v))
}
func PrintDebugf(format string, v ...interface{}) {
	//fmt.Println(fmt.Sprintln())
	mylog.output(0, 1, fmt.Sprintf(format, v))
}

func Println(v ...interface{}) {
	//fmt.Println(fmt.Sprintln())
	mylog.output(1, 1, fmt.Sprintln(v))
}
func Printlnlv(lv int, v ...interface{}) {
	//fmt.Println(fmt.Sprintln())
	mylog.output(lv, 1, fmt.Sprintln(v))
}

func Printf(format string, v ...interface{}) {

	//fmt.Println(fmt.Sprintln())
	mylog.output(1, 1, fmt.Sprintf(format, v))

}
func PrintfLv(lv int, format string, v ...interface{}) {

	//fmt.Println(fmt.Sprintln())
	mylog.output(lv, 1, fmt.Sprintf(format, v))

}

func PrintWarn(v ...interface{}) {

	mylog.output(2, 1, fmt.Sprintln(v))
}
func PrintWarnf(format string, v ...interface{}) {

	mylog.output(2, 1, fmt.Sprintf(format, v))
}

func PrintError(v ...interface{}) {

	mylog.output(3, 1, fmt.Sprintln(v))
}
func PrintErrorf(format string, v ...interface{}) {

	mylog.output(3, 1, fmt.Sprintf(format, v))
}

func PrintPanic(v ...interface{}) {

	mylog.output(4, 1, fmt.Sprintln(v))
}
func PrintPanicf(format string, v ...interface{}) {

	mylog.output(4, 1, fmt.Sprintf(format, v))
}

func PrintFatal(v ...interface{}) {

	mylog.output(5, 1, fmt.Sprintln(v))
}
func PrintFatalf(format string, v ...interface{}) {

	mylog.output(5, 1, fmt.Sprintf(format, v))
}

func CatchPanic() {
	if err := recover(); err != nil {
		mylog.dealPanic(1, err)
		//log.Print(err)
	}
}

func (this *std) output(lv, depth int, v string) {
	this.lk.Lock()
	defer this.lk.Unlock()
	if lv < 0 || lv > 5 {
		lv = 1
	}
	depth++
	var file string
	var line int
	var ok bool
	_, file, line, ok = runtime.Caller(depth)
	if !ok {
		file = "???"
		line = 0
	}
	this.print(lv, depth, file+":"+strconv.Itoa(line), v)
	depth--
}
func (this *std) dealPanic(depth int, v ...interface{}) {
	depth++
	trace := debug.Stack()
	trace_s := strings.Split(string(trace), "\n")
	for i := depth*2 + 5; i < len(trace_s)-1; i = i + 2 {
		s := strings.Split(strings.Replace(trace_s[i+1], "	", "", -1), " ")[0]
		this.print(4, 1, s, trace_s[i]+"\n")
		//PrintPanic(trace_s[i]+trace_s[i+1])
	}
	depth--
}

//func (this *std) print2(lv, depth int, file string, line int, v string) {
//	depth++
//	if this.color == 1 && this.dev == 0 {
//		this.buf.WriteString(fmt.Sprintf("\x1b[0;%dm ", colors[lv]))
//	}
//	tm := time.Now().Format("2006-01-02 15:03:04")
//	this.buf.WriteString(tm)
//	this.buf.WriteString(" ")
//	this.buf.WriteString(level[lv])
//	this.buf.WriteString(" ")
//	this.buf.WriteString(file)
//	this.buf.WriteString(":")
//	this.buf.WriteString(strconv.Itoa(line))
//	this.buf.WriteString(" ")
//		this.buf.WriteString(v)

//	if this.color == 1 && this.dev == 0 {
//		this.buf.WriteString("\x1b[0m")
//	}
//	//ColorPrint(this.buf.String(),2|8)
//	//Print
//	if this.mode == 0 || this.mode == 1 || this.mode == 3 {
//		if this.dev == 0 {
//			this.unixPrint()
//		} else {
//			//fmt.Println(this.buf)
//			this.windowsPrint(lv)
//		}
//
//	}
//	//Write to file
//	if this.loglv <= lv && (this.mode == 1 || this.mode == 2 || this.mode == 3){
//
//		if this.tm.Day() < time.Now().Day(){
//			this.createlogfile()
//		}
//		if this.color == 1 && this.dev == 0 {
//			this.fd.Write(this.buf.Bytes()[7 : this.buf.Len()-4])
//		} else {
//			this.fd.Write(this.buf.Bytes())
//		}
//
//	}
//	this.buf.Reset()
//	depth--
//}

func (this *std) print(lv, depth int, file string, v string) {
	depth++
	if this.color == 1 && this.dev == 0 {
		this.buf.WriteString(fmt.Sprintf("\x1b[0;%dm ", colors[lv]))
	}
	tm := time.Now().Format("2006-01-02 15:03:04")
	this.buf.WriteString(tm)
	this.buf.WriteString(" ")
	this.buf.WriteString(level[lv])
	this.buf.WriteString(" ")
	this.buf.WriteString(file)
	this.buf.WriteString(" ")
	this.buf.WriteString(v)
	if this.color == 1 && this.dev == 0 {
		this.buf.WriteString("\x1b[0m")
	}
	//ColorPrint(this.buf.String(),2|8)
	//Print
	if this.mode == 0 || this.mode == 1 || this.mode == 3 {
		if this.dev == 0 {
			this.unixPrint()
		} else {
			//fmt.Println(this.buf)
			this.windowsPrint(lv)
		}

	}
	//Write to file
	if this.loglv <= lv && (this.mode == 1 || this.mode == 2 || this.mode == 3) {
		if this.tm.Day() < time.Now().Day() {
			this.createlogfile()
		}
		if this.color == 1 && this.dev == 0 {
			this.fd.Write(this.buf.Bytes()[7 : this.buf.Len()-4])
		} else {
			this.fd.Write(this.buf.Bytes())
		}

	}
	this.buf.Reset()
	depth--
}
func (this *std) windowsPrint(lv int) { //设置终端字体颜色
	kernel32 := syscall.NewLazyDLL("kernel32.dll")
	proc := kernel32.NewProc("SetConsoleTextAttribute")
	handle, _, _ := proc.Call(uintptr(syscall.Stdout), uintptr(colors[lv]))
	this.out.Write(this.buf.Bytes())
	handle, _, _ = proc.Call(uintptr(syscall.Stdout), uintptr(7))
	CloseHandle := kernel32.NewProc("CloseHandle")
	CloseHandle.Call(handle)
}
func (this *std) unixPrint() {
	this.out.Write(this.buf.Bytes())
}

//创建log文件
func (this *std) createlogfile() {
	t1 := time.Now().Format("2006_01_02_15_04_05")
	path := "./log"
	pathcreate(path)
	path += "/" + t1 + ".txt"
	Println(path)
	logfile, err := os.OpenFile(path, os.O_TRUNC|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		Printf("%s\r\n", err.Error())
		os.Exit(-1)
	}
	this.fd = logfile
	if this.mode == 3 {
		os.Stdout = this.fd
		os.Stderr = this.fd
	}

	return
}

//自动创建文件夹
func pathcreate(path string) error {
	var err error
	_, err = os.Stat(path)
	if err != nil {
		Println(err)
		os.MkdirAll(path, 0711)
	}
	return err
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
