package logger

import (
	"fmt"
	"log"
)

type FmtLogger struct {
}

func NewFmtLogger() Logger {
	return &FmtLogger{}
}

func (logger *FmtLogger) Print(args ...any) {
	log.Print(args...)
}

func (logger *FmtLogger) Println(args ...any) {
	log.Println(args...)
}

func (logger *FmtLogger) Printf(format string, args ...any) {
	log.Printf(format, args...)
}

func (logger *FmtLogger) Warn(args ...any) {
	log.Println(args...)
}

func (logger *FmtLogger) Error(err error, args ...any) {
	msg := fmt.Sprint(args...)
	log.Println(err.Error() + " " + msg)
}
