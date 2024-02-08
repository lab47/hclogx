package hclogx

import (
	"io"
	"log"

	"github.com/hashicorp/go-hclog"
)

type Rest interface {
	// Indicate if TRACE logs would be emitted. This and the other Is* guards
	// are used to elide expensive logging code based on the current level.
	IsTrace() bool

	// Indicate if DEBUG logs would be emitted. This and the other Is* guards
	IsDebug() bool

	// Indicate if INFO logs would be emitted. This and the other Is* guards
	IsInfo() bool

	// Indicate if WARN logs would be emitted. This and the other Is* guards
	IsWarn() bool

	// Indicate if ERROR logs would be emitted. This and the other Is* guards
	IsError() bool

	// ImpliedArgs returns With key/value pairs
	ImpliedArgs() []interface{}

	// Creates a sublogger that will always have the given key/value pairs
	With(args ...interface{}) hclog.Logger

	// Returns the Name of the logger
	Name() string

	// Create a logger that will prepend the name string on the front of all messages.
	// If the logger already has a name, the new value will be appended to the current
	// name. That way, a major subsystem can use this to decorate all it's own logs
	// without losing context.
	Named(name string) hclog.Logger

	// Create a logger that will prepend the name string on the front of all messages.
	// This sets the name of the logger to the value directly, unlike Named which honor
	// the current name as well.
	ResetNamed(name string) hclog.Logger

	// Updates the level. This should affect all related loggers as well,
	// unless they were created with IndependentLevels. If an
	// implementation cannot update the level on the fly, it should no-op.
	SetLevel(level hclog.Level)

	// Returns the current level
	GetLevel() hclog.Level

	// Return a value that conforms to the stdlib log.Logger interface
	StandardLogger(opts *hclog.StandardLoggerOptions) *log.Logger

	// Return a value that conforms to io.Writer, which can be passed into log.SetOutput()
	StandardWriter(opts *hclog.StandardLoggerOptions) io.Writer
}
