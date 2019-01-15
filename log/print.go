package log

import (
	"fmt"
)

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

//------------------------
func LogDebug(v ...interface{}) {
	//fmt.Println(fmt.Sprintln())
	mylog.output_log(0, 1, fmt.Sprintln(v))
}
func LogDebugf(format string, v ...interface{}) {
	//fmt.Println(fmt.Sprintln())
	mylog.output_log(0, 1, fmt.Sprintf(format, v))
}

func LogPrintln(v ...interface{}) {
	//fmt.Println(fmt.Sprintln())
	mylog.output_log(1, 1, fmt.Sprintln(v))
}
func LogPrintlnlv(lv int, v ...interface{}) {
	//fmt.Println(fmt.Sprintln())
	mylog.output_log(lv, 1, fmt.Sprintln(v))
}

func LogPrintf(format string, v ...interface{}) {

	//fmt.Println(fmt.Sprintln())
	mylog.output_log(1, 1, fmt.Sprintf(format, v))

}
func LogPrintfLv(lv int, format string, v ...interface{}) {

	//fmt.Println(fmt.Sprintln())
	mylog.output_log(lv, 1, fmt.Sprintf(format, v))

}

func LogWarn(v ...interface{}) {

	mylog.output_log(2, 1, fmt.Sprintln(v))
}
func LogWarnf(format string, v ...interface{}) {

	mylog.output_log(2, 1, fmt.Sprintf(format, v))
}

func LogError(v ...interface{}) {
	mylog.output_log(3, 1, fmt.Sprintln(v))
}
func LogErrorf(format string, v ...interface{}) {

	mylog.output_log(3, 1, fmt.Sprintf(format, v))
}

func LogPanic(v ...interface{}) {
	mylog.output_log(4, 1, fmt.Sprintln(v))
}
func LogPanicf(format string, v ...interface{}) {

	mylog.output_log(4, 1, fmt.Sprintf(format, v))
}

func LogFatal(v ...interface{}) {

	mylog.output_log(5, 1, fmt.Sprintln(v))
}
func LogFatalf(format string, v ...interface{}) {

	mylog.output_log(5, 1, fmt.Sprintf(format, v))
}
