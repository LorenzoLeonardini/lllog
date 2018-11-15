package lllog

import (
	"fmt"
	"github.com/fatih/color"
	"os"
	"time"
)

type Logger struct {
	name           string
	format         string
	hasFormat      bool
	colors         map[string]func(...interface{}) string
	file           *os.File
	disableConsole bool
}

func New(name string) *Logger {
	l := new(Logger)
	l.colors = make(map[string]func(...interface{}) string)
	l.colors["log"] = color.New(color.FgHiWhite).SprintFunc()
	l.colors["warn"] = color.New(color.FgHiBlue, color.Bold).SprintFunc()
	l.colors["err"] = color.New(color.FgHiRed, color.Bold).SprintFunc()
	l.colors["fatal"] = color.New(color.FgHiMagenta, color.Bold).SprintFunc()
	l.name = name
	return l
}

func (l *Logger) SetLogColor(c color.Attribute) {
	l.colors["log"] = color.New(c).SprintFunc()
}

func (l *Logger) SetWarnColor(c color.Attribute) {
	l.colors["warn"] = color.New(c, color.Bold).SprintFunc()
}

func (l *Logger) SetErrColor(c color.Attribute) {
	l.colors["err"] = color.New(c, color.Bold).SprintFunc()
}

func (l *Logger) SetFatalColor(c color.Attribute) {
	l.colors["fatal"] = color.New(c, color.Bold).SprintFunc()
}

func (l *Logger) Log(format string, args ...interface{}) {
	if !l.disableConsole {
		fmt.Printf(l.colors["log"]("[Log]  ", l.getHeader(), l.name, ": ")+format, args...)
	}
}

func (l *Logger) Warn(format string, args ...interface{}) {
	if !l.disableConsole {
		fmt.Printf(l.colors["warn"]("[Warn] ", l.getHeader(), l.name, ": ")+format, args...)
	}
}

func (l *Logger) Err(format string, args ...interface{}) {
	if !l.disableConsole {
		fmt.Printf(l.colors["err"]("[Err]  ", l.getHeader(), l.name, ": ")+format, args...)
	}
}

func (l *Logger) Fatal(format string, args ...interface{}) {
	if !l.disableConsole {
		fmt.Printf(l.colors["fatal"]("[Fatal]", l.getHeader(), l.name, ": ")+format, args...)
	}
}

func (l *Logger) WriteToConsole(enabled bool) {
	l.disableConsole = !enabled
}

func (l *Logger) setFormat(format string) {
	l.hasFormat = true
	l.format = format
}

func (l *Logger) getHeader() string {
	if l.hasFormat {
		return "[" + time.Now().Format(l.format) + "] "
	}
	return ""
}
