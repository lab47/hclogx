package hclogx

import (
	"github.com/hashicorp/go-hclog"
)

type Minimum interface {
	Log(level hclog.Level, msg string, args ...any)
}

type MinLogger struct {
	Next Minimum
}

// Args are alternating key, val pairs
// keys must be strings
// vals can be any type, but display is implementation specific
// Emit a message and key/value pairs at a provided log level
func (m *MinLogger) Log(level hclog.Level, msg string, args ...interface{}) {
	m.Next.Log(level, msg, args...)
}

// Emit a message and key/value pairs at the TRACE level
func (m *MinLogger) Trace(msg string, args ...interface{}) {
	m.Next.Log(hclog.Trace, msg, args...)
}

// Emit a message and key/value pairs at the DEBUG level
func (m *MinLogger) Debug(msg string, args ...interface{}) {
	m.Next.Log(hclog.Debug, msg, args...)
}

// Emit a message and key/value pairs at the INFO level
func (m *MinLogger) Info(msg string, args ...interface{}) {
	m.Next.Log(hclog.Info, msg, args...)
}

// Emit a message and key/value pairs at the WARN level
func (m *MinLogger) Warn(msg string, args ...interface{}) {
	m.Next.Log(hclog.Warn, msg, args...)
}

// Emit a message and key/value pairs at the ERROR level
func (m *MinLogger) Error(msg string, args ...interface{}) {
	m.Next.Log(hclog.Error, msg, args...)
}
