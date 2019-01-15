package log

import "testing"

func Test(t *testing.T) {
	defer CatchPanic()
	LogInit(Log_mode_PrintAndWrite, Log_Lv_Debug, Unix, Color_On)
	PrintError("rqwrewqrewq")
	LogError("fdhsajkfhdsagfjgsdajkfgjksa")
	panic("fhdsakfgjsdagjfdgsajfgjdsa")
}
