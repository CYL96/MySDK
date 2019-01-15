package log

import (
	"fmt"
	"runtime/debug"
	"strings"
)

func CatchPanic() {
	if err := recover(); err != nil {
		mylog.dealPanic(1, err)
		//log.Print(err)
	}
}

func (this *std) dealPanic(depth int, v ...interface{}) {
	depth++
	trace := debug.Stack()
	trace_s := strings.Split(string(trace), "\n")
	this.print_log(4, 1, "", fmt.Sprintln("---------------------------[panic start]-----------------------"))
	this.print_log(4, 1, "", fmt.Sprintln("ErrInfo:", v))
	for i := depth*2 + 5; i < len(trace_s)-1; i = i + 2 {
		s := strings.Split(strings.Replace(trace_s[i+1], "	", "", -1), " ")[0]
		this.print_log(4, 1, s, trace_s[i]+"\n")
		//PrintPanic(trace_s[i]+trace_s[i+1])
	}
	this.print_log(4, 1, "", fmt.Sprintln("---------------------------[panic end]-----------------------"))
	depth--
}
