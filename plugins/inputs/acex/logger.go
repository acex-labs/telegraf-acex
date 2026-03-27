package acex

import (
	"fmt"
	"os"

	"github.com/influxdata/telegraf"
)

type LogLevel int

const (
	DebugLevel LogLevel = iota
	InfoLevel
	WarnLevel
	ErrorLevel
	TraceLevel
)

// StderrLogger implements telegraf.Logger
type StderrLogger struct{}

func (l *StderrLogger) Level() telegraf.LogLevel {
	return telegraf.Info
}

// AddAttribute can be ignored or stored if desired
func (l *StderrLogger) AddAttribute(key string, value any) {
	// no-op for now
}

// Errorf/Error
func (l *StderrLogger) Errorf(format string, args ...any) {
	fmt.Fprintf(os.Stderr, "E! "+format+"\n", args...)
}
func (l *StderrLogger) Error(args ...any) {
	fmt.Fprintf(os.Stderr, "E! %v\n", fmt.Sprint(args...))
}

// Warnf/Warn
func (l *StderrLogger) Warnf(format string, args ...any) {
	fmt.Fprintf(os.Stderr, "W! "+format+"\n", args...)
}
func (l *StderrLogger) Warn(args ...any) {
	fmt.Fprintf(os.Stderr, "W! %v\n", fmt.Sprint(args...))
}

// Infof/Info
func (l *StderrLogger) Infof(format string, args ...any) {
	fmt.Fprintf(os.Stderr, "I! "+format+"\n", args...)
}
func (l *StderrLogger) Info(args ...any) {
	fmt.Fprintf(os.Stderr, "I! %v\n", fmt.Sprint(args...))
}

// Debugf/Debug
func (l *StderrLogger) Debugf(format string, args ...any) {
	fmt.Fprintf(os.Stderr, "D! "+format+"\n", args...)
}
func (l *StderrLogger) Debug(args ...any) {
	fmt.Fprintf(os.Stderr, "D! %v\n", fmt.Sprint(args...))
}

// Tracef/Trace
func (l *StderrLogger) Tracef(format string, args ...any) {
	fmt.Fprintf(os.Stderr, "T! "+format+"\n", args...)
}
func (l *StderrLogger) Trace(args ...any) {
	fmt.Fprintf(os.Stderr, "T! %v\n", fmt.Sprint(args...))
}
