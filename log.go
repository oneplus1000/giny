package giny

import "fmt"

//LogTypeErr  error
const LogTypeErr = 1

//LogTypeWarn warn
const LogTypeWarn = 2

//LogTypeInfo info
const LogTypeInfo = 3

//LogTypeDebug debuf
const LogTypeDebug = 4

var writeLog func(logtype int, msg string)

//SetWriteLog set logger
func SetWriteLog(wl func(logtype int, msg string)) {
	writeLog = wl
}

//LogTypeDesc แปลง logtype เป็น string
func LogTypeDesc(logtype int) string {
	switch logtype {
	case LogTypeErr:
		return "error"
	case LogTypeWarn:
		return "warn"
	case LogTypeInfo:
		return "info"
	case LogTypeDebug:
		return "debug"
	}
	return ""
}

//LogErr error
func LogErr(err error) {
	LogErrf(err, "")
}

//LogErrf error
func LogErrf(err error, format string, a ...interface{}) {
	msg := fmt.Sprintf(format, a)
	msg += fmt.Sprintf("\n%+v", err)
	writeLog(LogTypeErr, fmt.Sprintf("%s\n", err.Error()))
}

//LogWarn warn
func LogWarn(msg string) {
	LogWarnf("%s", msg)
}

//LogWarnf warnf
func LogWarnf(format string, a ...interface{}) {
	msg := fmt.Sprintf(format, a)
	writeLog(LogTypeWarn, fmt.Sprintf("%s\n", msg))
}

//LogDebug debug
func LogDebug(msg string) {
	LogDebugf("%s", msg)
}

//LogDebugf debugf
func LogDebugf(format string, a ...interface{}) {
	msg := fmt.Sprintf(format, a)
	writeLog(LogTypeDebug, fmt.Sprintf("%s\n", msg))
}

//LogInfo info
func LogInfo(msg string) {
	LogInfof("%s", msg)
}

//LogInfof infof
func LogInfof(format string, a ...interface{}) {
	msg := fmt.Sprintf(format, a)
	writeLog(LogTypeInfo, fmt.Sprintf("%s\n", msg))
}
