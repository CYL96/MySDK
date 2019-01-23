package log

import (
	"fmt"
	"runtime"
	"strconv"
	"syscall"
	"time"
)

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
	if this.color == Color_On && this.dev == Unix {
		this.buf.WriteString(fmt.Sprintf("\x1b[0;%dm ", colors[lv]))
	}
	tm := time.Now().Format("2006-01-02 15:04:05")
	this.buf.WriteString(tm)
	this.buf.WriteString(" ")
	this.buf.WriteString(level[lv])
	this.buf.WriteString("[Run]")
	this.buf.WriteString(" ")
	this.buf.WriteString(file)
	this.buf.WriteString(" ")
	this.buf.WriteString(v)
	if this.color == Color_On && this.dev == Unix {
		this.buf.WriteString("\x1b[0m")
	}
	//ColorPrint(this.buf.String(),2|8)
	//Print
	if this.mode != Log_mode_Write {
		if this.dev == Unix {
			this.unixPrint()
		} else {
			//fmt.Println(this.buf)
			this.windowsPrint(lv)
		}

	}
	//Write to file
	if this.loglv <= lv && (this.mode == Log_mode_Write || this.mode == Log_mode_PrintAndWrite || this.mode == Log_mode_All) {
		if this.tm.Day() < time.Now().Day() {
			this.createlogfile()
		}
		if this.color == Color_On && this.dev == Unix {
			this.fd_run.Write(this.buf.Bytes()[7 : this.buf.Len()-4])
		} else {
			this.fd_run.Write(this.buf.Bytes())
		}
		this.fd_run.Sync()
	}
	this.buf.Reset()
	depth--
}

func (this *std) output_log(lv, depth int, v string) {
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
	this.print_log(lv, depth, file+":"+strconv.Itoa(line), v)
	depth--
}

func (this *std) print_log(lv, depth int, file string, v string) {
	depth++
	if this.color == Color_On && this.dev == Unix {
		this.buf.WriteString(fmt.Sprintf("\x1b[0;%dm ", colors[lv]))
	}
	tm := time.Now().Format("2006-01-02 15:03:04")
	this.buf.WriteString(tm)
	this.buf.WriteString(" ")
	this.buf.WriteString(level[lv])
	this.buf.WriteString("[Log]")
	this.buf.WriteString(" ")
	this.buf.WriteString(file)
	this.buf.WriteString(" ")
	this.buf.WriteString(v)
	if this.color == Color_On && this.dev == Unix {
		this.buf.WriteString("\x1b[0m")
	}
	//ColorPrint(this.buf.String(),2|8)
	//Print
	if this.mode != Log_mode_Write {
		if this.dev == Unix {
			this.unixPrint()
		} else {
			//fmt.Println(this.buf)
			this.windowsPrint(lv)
		}

	}
	//Write to file
	if this.loglv <= lv && (this.mode == Log_mode_Write || this.mode == Log_mode_PrintAndWrite || this.mode == Log_mode_All) {
		if this.tm.Day() < time.Now().Day() {
			this.createlogfile()
		}
		if this.color == Color_On && this.dev == Unix {
			this.fd_log.Write(this.buf.Bytes()[7 : this.buf.Len()-4])
		} else {
			this.fd_log.Write(this.buf.Bytes())
		}
		this.fd_log.Sync()
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
