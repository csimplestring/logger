package logger

import (
	"fmt"
	"io"
	"os"
	"time"
)

type Logger interface {
	Log(level int, messages ...interface{})
}

type TextLogger struct {
	writer io.Writer
}

func (f *TextLogger) Log(level int, messages ...interface{}) {
	fmt.Fprintln(f.writer, messages...)
}

type FilterLogger struct {
	logger Logger
	filter func(level int) bool
}

func (f *FilterLogger) Log(level int, messages ...interface{}) {
	if f.filter(level) {
		f.logger.Log(level, messages...)
	}
}

type TimestampLogger struct {
	logger Logger
}

func (t *TimestampLogger) Log(level int, messages ...interface{}) {
	t.logger.Log(level, time.Now().Format(time.RFC850)+fmt.Sprint(messages))
}

type SequenceLogger struct {
	loggers []Logger
}

func (s *SequenceLogger) Log(level int, messages ...interface{}) {
	for _, logger := range s.loggers {
		logger.Log(level, messages...)
	}
}

type LoggerBuilder struct {
	logger Logger
}

func NewBuilder() *LoggerBuilder {
	return &LoggerBuilder{}
}

func (l *LoggerBuilder) WriteTo(writer io.Writer) *LoggerBuilder {
	l.logger = &TextLogger{writer: writer}
	return l
}

func (l *LoggerBuilder) WriteToConsole() *LoggerBuilder {
	l.logger = &TextLogger{writer: os.Stdout}
	return l
}

func (l *LoggerBuilder) WithTimestamp() *LoggerBuilder {
	l.logger = &TimestampLogger{
		logger: l.logger,
	}
	return l
}

func (l *LoggerBuilder) If(condition func(level int) bool) *LoggerBuilder {
	l.logger = &FilterLogger{
		logger: l.logger,
		filter: condition,
	}
	return l
}

func (l *LoggerBuilder) Build() Logger {
	return l.logger
}
