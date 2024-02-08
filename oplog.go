package hclogx

import (
	"github.com/hashicorp/go-hclog"
)

type UnheardEntries struct {
	Message string
	Level   hclog.Level
	Args    []any
}

type OpLogger struct {
	Rest
	MinLogger

	Original hclog.Logger
	Level    hclog.Level

	Queue []UnheardEntries
}

func NewOpLogger(log hclog.Logger) *OpLogger {
	l := &OpLogger{
		Rest:     log,
		Original: log,
		Level:    log.GetLevel(),
	}

	l.Next = l

	return l
}

// Args are alternating key, val pairs
// keys must be strings
// vals can be any type, but display is implementation specific
// Emit a message and key/value pairs at a provided log level
func (o *OpLogger) Log(level hclog.Level, msg string, args ...interface{}) {
	if level < o.Level {
		o.Queue = append(o.Queue, UnheardEntries{
			Message: msg,
			Level:   level,
			Args:    args,
		})

		return
	}

	o.Original.Log(level, msg, args...)
}

// The Is* functions always return true for us so that we can capture any
// logs that are guarded with them.

func (o *OpLogger) IsTrace() bool {
	return true
}

func (o *OpLogger) IsDebug() bool {
	return true
}

func (o *OpLogger) IsInfo() bool {
	return true
}

func (o *OpLogger) IsWarn() bool {
	return true
}

func (o *OpLogger) IsError() bool {
	return true
}

func (o *OpLogger) Flush() {
	for _, uh := range o.Queue {
		o.Original.Error(uh.Message, append([]any{"level", uh.Level}, uh.Args...)...)
	}
}

type Flusher interface {
	Flush()
}

func Flush(log hclog.Logger) {
	if f, ok := log.(Flusher); ok {
		f.Flush()
	}
}

func FlushOnError(log hclog.Logger, errP *error) {
	if *errP != nil {
		Flush(log)
	}
}
