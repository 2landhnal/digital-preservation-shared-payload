package logger

var GlobalLogger Logger

func InitGlobalLogger() {
	GlobalLogger = &FmtLogger{}
}
