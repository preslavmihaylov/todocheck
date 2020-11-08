package logger

import (
	"fmt"
	"sync"
)

var lock sync.Mutex
var log *Logger

// Logger provided simple logger.
type Logger struct {
	verbose bool
}

// Setup configures logger
func Setup(verbose bool) *Logger {
	lock.Lock()
	defer lock.Unlock()

	if log == nil {
		log = &Logger{
			verbose,
		}
	}
	return log
}

// Info prints in verbose mode to standart output. Format according to fmt.Println.
func (l *Logger) Info(args ...interface{}) {
	Info(args...)
}

// Info prints in verbose mode to standart output. Format according to fmt.Println.
func Info(args ...interface{}) {
	if log.verbose {
		fmt.Println(args...)
	}
}

// Infof formatted prints in verbose mode to standart output. Format according to fmt.Printf.
func (l *Logger) Infof(format string, args ...interface{}) {
	Infof(format, args...)
}

// Infof formatted prints in verbose mode to standart output. Format according to fmt.Printf.
func Infof(format string, args ...interface{}) {
	if log.verbose {
		fmt.Printf(format, args...)
	}
}
