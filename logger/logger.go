package logger

import (
	"fmt"
	"sync"
)

var once sync.Once
var log *Logger

// Logger provided simple logger.
type Logger struct {
	verbose bool
}

// Setup configures logger
func Setup(verbose bool) {
	once.Do(func() {
		log = &Logger{
			verbose,
		}
	})
}

// Info prints in verbose mode to standart output. Format according to fmt.Println.
func Info(args ...interface{}) {
	log.Info(args...)
}

// Info prints in verbose mode to standart output. Format according to fmt.Println.
func (l *Logger) Info(args ...interface{}) {
	if l.verbose {
		fmt.Println(args...)
	}
}

// Infof formatted prints in verbose mode to standart output. Format according to fmt.Printf.
func Infof(format string, args ...interface{}) {
	log.Infof(format, args...)
}

// Infof formatted prints in verbose mode to standart output. Format according to fmt.Printf.
func (l *Logger) Infof(format string, args ...interface{}) {
	if l.verbose {
		fmt.Printf(format, args...)
	}
}
