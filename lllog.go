package lllog

import (
	"encoding/json"
	"fmt"
	"github.com/fatih/color"
	"os"
	"strings"
	"sync"
	"time"
)

type Logger struct {
	name           string
	format         string
	hasFormat      bool
	colors         map[string]func(...interface{}) string
	file           *os.File
	disableConsole bool
	mutex          *sync.Mutex
	logs           []LogMessage
	formatter      func(*LogMessage) string
}

type LogMessage struct {
	Level     string
	Msg       string
	Timestamp int64
	Format    string
	Name      string
}

func ConsoleFromatter(log *LogMessage) string {
	if log.Format != "" {
		return "[" + log.Level + "]" + "[" + time.Unix(log.Timestamp, 0).Format(log.Format) + "] " + log.Name + ": " + log.Msg
	} else {
		return "[" + log.Level + "]" + log.Name + ": " + log.Msg
	}
}

func JSONFormatter(log *LogMessage) string {
	b, _ := json.Marshal(log)
	return string(b) + "\n"
}

func New(name string) *Logger {
	l := new(Logger)
	l.colors = make(map[string]func(...interface{}) string)
	l.colors["Log"] = color.New(color.FgHiWhite).SprintFunc()
	l.colors["Warn"] = color.New(color.FgHiBlue, color.Bold).SprintFunc()
	l.colors["Err"] = color.New(color.FgHiRed, color.Bold).SprintFunc()
	l.colors["Fatal"] = color.New(color.FgHiMagenta, color.Bold).SprintFunc()
	l.name = name
	l.mutex = &sync.Mutex{}
	l.logs = make([]LogMessage, 0)
	l.formatter = ConsoleFromatter
	return l
}

func (l *Logger) SetOutputFormatter(f func(*LogMessage) string) {
	l.formatter = f
}

func (l *Logger) LogToFile(path string) {
	if !strings.HasSuffix(path, "/") {
		path += "/"
	}
	l.file, _ = os.OpenFile(path+time.Now().Format("2006-01-02")+".log", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0644)
	l.file.WriteString("\n\n")
}

func (l *Logger) SetLogColor(c color.Attribute) {
	l.colors["Log"] = color.New(c).SprintFunc()
}

func (l *Logger) SetWarnColor(c color.Attribute) {
	l.colors["Warn"] = color.New(c, color.Bold).SprintFunc()
}

func (l *Logger) SetErrColor(c color.Attribute) {
	l.colors["Err"] = color.New(c, color.Bold).SprintFunc()
}

func (l *Logger) SetFatalColor(c color.Attribute) {
	l.colors["Fatal"] = color.New(c, color.Bold).SprintFunc()
}

func (l *Logger) golog(log LogMessage) string {
	l.mutex.Lock()
	l.logs = append(l.logs, log)
	if !l.disableConsole {
		fmt.Printf(l.colors[log.Level]("[", log.Level, "]", l.getHeader(), l.name, ": ") + log.Msg)
	}
	ret := l.formatter(&log)
	if l.file != nil {
		l.file.WriteString(ret)
	}
	l.mutex.Unlock()
	return ret
}

func (l *Logger) Log(format string, args ...interface{}) string {
	return l.golog(LogMessage{Level: "Log", Msg: fmt.Sprintf(format, args...), Timestamp: time.Now().Unix(), Format: l.format, Name: l.name})
}

func (l *Logger) Warn(format string, args ...interface{}) string {
	return l.golog(LogMessage{Level: "Warn", Msg: fmt.Sprintf(format, args...), Timestamp: time.Now().Unix(), Format: l.format, Name: l.name})
}

func (l *Logger) Err(format string, args ...interface{}) string {
	return l.golog(LogMessage{Level: "Err", Msg: fmt.Sprintf(format, args...), Timestamp: time.Now().Unix(), Format: l.format, Name: l.name})
}

func (l *Logger) Fatal(format string, args ...interface{}) string {
	return l.golog(LogMessage{Level: "Fatal", Msg: fmt.Sprintf(format, args...), Timestamp: time.Now().Unix(), Format: l.format, Name: l.name})
}

func (l *Logger) WriteToConsole(enabled bool) {
	l.disableConsole = !enabled
}

func (l *Logger) SetFormat(format string) {
	l.hasFormat = true
	l.format = format
}

func (l *Logger) getHeader() string {
	if l.hasFormat {
		return "[" + time.Now().Format(l.format) + "] "
	}
	return ""
}
