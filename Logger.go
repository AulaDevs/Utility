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
	level  uint8
	file   *os.File
	closed bool
}

// Constructor
func NewLogger(filepath string, level loggerLevel) Logger {
	file, err := os.OpenFile(filepath, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0644)

	if err != nil {
		panic(err)
	}

	return Logger{
		level.level,
		file,
		false,
	}
}

// Standard methods
func (logger *Logger) Debug(format string, a ...any) {
	if logger.level > DEBUG.level || logger.closed {
		return
	}

	logger.file.WriteString(fmt.Sprintf("[DEBUG] %s: %s\n", time.Now().UTC().String(), fmt.Sprintf(format, a...)))
	logger.file.Sync()
}

func (logger *Logger) Info(format string, a ...any) {
	if logger.level > INFO.level || logger.closed {
		return
	}

	logger.file.WriteString(fmt.Sprintf("[INFO] %s: %s\n", time.Now().UTC().String(), fmt.Sprintf(format, a...)))
	logger.file.Sync()
}

func (logger *Logger) Warn(format string, a ...any) {
	if logger.level > WARN.level || logger.closed {
		return
	}

	logger.file.WriteString(fmt.Sprintf("[WARN] %s: %s\n", time.Now().UTC().String(), fmt.Sprintf(format, a...)))
	logger.file.Sync()
}

func (logger *Logger) Error(format string, a ...any) {
	if logger.level > ERROR.level || logger.closed {
		return
	}

	logger.file.WriteString(fmt.Sprintf("[ERROR] %s: %s\n", time.Now().UTC().String(), fmt.Sprintf(format, a...)))
	logger.file.Sync()
}

func (logger *Logger) Close(format string, a ...any) {
	logger.closed = true
	logger.file.Close()
}
