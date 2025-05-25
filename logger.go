package luxe

import (
	"fmt"
	"log"
	"os"
)

type Logger struct {
	errorLogger *log.Logger
	infoLogger  *log.Logger
	warnLogger  *log.Logger
}

func NewLogger() *Logger {
	return &Logger{
		errorLogger: log.New(os.Stderr, "\033[31m[ERROR]\033[0m ", log.Ldate|log.Ltime|log.Lshortfile),
		infoLogger:  log.New(os.Stdout, "\033[32m[INFO]\033[0m ", log.Ldate|log.Ltime),
		warnLogger:  log.New(os.Stdout, "\033[33m[WARN]\033[0m ", log.Ldate|log.Ltime),
	}
}

func (l *Logger) Error(format string, v ...any) {
	l.errorLogger.Output(2, fmt.Sprintf(format, v...))
}

func (l *Logger) Info(format string, v ...any) {
	l.infoLogger.Output(2, fmt.Sprintf(format, v...))
}

func (l *Logger) Warn(format string, v ...any) {
	l.warnLogger.Output(2, fmt.Sprintf(format, v...))
}
