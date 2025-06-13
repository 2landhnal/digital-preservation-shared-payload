package logger

type Logger interface {
	Print(args ...any)
	Println(args ...any)
	Printf(format string, args ...any)
	Warn(args ...any)
	Error(err error, args ...any)
}
