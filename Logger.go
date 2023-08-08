package Utility

import (
	"fmt"
	"os"
	"time"
)

type loggerLevel struct {
	level uint8
}

var (
	DEBUG = loggerLevel{1}
	INFO  = loggerLevel{2}
	WARN  = loggerLevel{3}
	ERROR = loggerLevel{4}
)

type Logger struct {
	level   uint8
	file    *os.File
	closed  bool
	console bool
}

// Constructor
func NewLogger(filepath string, level loggerLevel, console bool) Logger {
	file, err := os.OpenFile(filepath, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0644)

	if err != nil {
		panic(err)
	}

	return Logger{
		level.level,
		file,
		false,
		console,
	}
}

// Standard methods
func (logger *Logger) write(level loggerLevel, name string, format string, a ...any) {
	if logger.level > level.level || logger.closed {
		return
	}

	message := fmt.Sprintf("[%s] %s: %s\n", name, time.Now().UTC().String(), fmt.Sprintf(format, a...))

	if logger.console {
		fmt.Println(message)
	}

	logger.file.WriteString(message)
	logger.file.Sync()
}

func (logger *Logger) Debug(format string, a ...any) {
	logger.write(DEBUG, "DEBUG", format, a...)
}

func (logger *Logger) Info(format string, a ...any) {
	logger.write(INFO, "INFO", format, a...)
}

func (logger *Logger) Warn(format string, a ...any) {
	logger.write(WARN, "WARN", format, a...)
}

func (logger *Logger) Error(format string, a ...any) {
	logger.write(ERROR, "ERROR", format, a...)
}

func (logger *Logger) Close(format string, a ...any) {
	logger.closed = true
	logger.file.Close()
}
