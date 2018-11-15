package lllog

import (
	"fmt"
	"github.com/fatih/color"
)

type Logger struct {
	name   string
	format string
}

// IGNORE THIS IS A TEST FOR THE VERY FIRST COMMIT
func (l *Logger) Log(format string, args ...interface{}) color.Attribute {
	fmt.Printf(format, args...)
	return color.FgHiRed
}

func New(name string) *Logger {
	l := new(Logger)
	l.name = name
	return l
}
