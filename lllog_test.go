package lllog

import (
	"fmt"
	"testing"
)

func TestLog(t *testing.T) {
	logger := New("Log Message")
	logger.Log("This is a simple log test\n")
	logger.Log("As you can see, %s works beautifully! %d * %d = %d\n", "print format", 3, 4, 3*4)
}

func TestWarn(t *testing.T) {
	logger := New("Warn Message")
	logger.Warn("This is a simple warn test\n")
	logger.Warn("As you can see, different log levels have different colors\n")
}

func TestErr(t *testing.T) {
	logger := New("Err Message")
	logger.Err("This is a simple warn test\n")
	logger.Err("As you can see, different log levels have different colors\n")
}

func TestFatal(t *testing.T) {
	logger := New("Fatal Message")
	logger.Fatal("This is a simple warn test\n")
	logger.Fatal("As you can see, different log levels have different colors\n")
}

func TestFormat(t *testing.T) {
	logger := New("Format Example")
	logger.setFormat("02 January 2006 - 15:04")
	logger.Log("Loggers can also have time/date prefix\n")
	logger.Warn("It is formatted using the golang time format\n")
}

func TestDisableConsole(t *testing.T) {
	logger := New("Disable console")
	logger.Log("You can disable writing to console. You won't see the next log\n")
	logger.WriteToConsole(false)
	logger.Log("You can't see me!\n")
	logger.WriteToConsole(true)
	logger.Log("Would you believe there was something before me?\n")
}

// Just want some whitespaces before the "PASS" print
func TestWhiteSpaces(t *testing.T) {
	fmt.Printf("\n\n")
}
