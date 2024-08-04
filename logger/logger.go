package logger

import (
	"log"
	"os"
)

type Logger interface {
	Debug(v ...interface{})
	Info(v ...interface{})
	Error(v ...interface{})
	Warn(v ...interface{})
}

type ConsoleLogger struct {
	debug *log.Logger
	error *log.Logger
	warn  *log.Logger
	info  *log.Logger
}

func NewConsoleLogger() *ConsoleLogger {
	return &ConsoleLogger{
		debug: log.New(os.Stdout, "DEBUG : ", log.LstdFlags),
		info:  log.New(os.Stdout, "INFO : ", log.LstdFlags),
		warn:  log.New(os.Stdout, "WARN : ", log.LstdFlags),
		error: log.New(os.Stdout, "ERROR : ", log.LstdFlags),
	}
}

// methods defined for the ConsoleLogger type

func (l *ConsoleLogger) Debug(v ...interface{}) {
	l.debug.Println(v...)
}

func (l *ConsoleLogger) Info(v ...interface{}) {
	l.info.Println(v...)
}

func (l *ConsoleLogger) Warn(v ...interface{}) {
	l.warn.Println(v...)
}

func (l *ConsoleLogger) Error(v ...interface{}) {
	l.error.Println(v...)
}

var CustomLogger = NewConsoleLogger()
