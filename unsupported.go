// +build windows plan9 nacl

package gsyslog

import (
	"io"
	"log"
)

type fakeSyslogger struct {
	logger   *log.Logger
	priority Priority
}

// WriteLevel writes log message with specified priority level
func (l fakeSyslogger) WriteLevel(p Priority, s string) error {
	l.logger.Output(int(p), s)
	return nil
}

// Write writes log message with default priority level
func (l fakeSyslogger) Write(s []byte) (int, error) {
	l.logger.Output(int(l.priority), string(s))
	return len(s), nil
}

// Close does nothing
func (l fakeSyslogger) Close() error {
	return nil
}

// SetOtherWriter changes log message writer
func (l fakeSyslogger) SetOtherWriter(w io.Writer) {
	l.logger.SetOutput(w)
}

// Fake writer, just prints out an input
type printWriter struct {
}

func (w printWriter) Write(p []byte) (n int, err error) {
	println(string(p))
	n = len(p)
	err = nil
	return
}

// NewLogger is used to construct a new fakeSyslogger
func NewLogger(p Priority, facility, tag string) (Syslogger, error) {
	logger := log.New(printWriter{}, facility+" "+tag+" ", log.LstdFlags)
	return fakeSyslogger{logger, p}, nil
}

// The stub, returns a new fakeSyslogger
func DialLogger(network, raddr string, p Priority, facility, tag string) (Syslogger, error) {
	return NewLogger(p, facility, tag)
}
