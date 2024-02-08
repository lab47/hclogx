package hclogx

import (
	"bytes"
	"testing"

	"github.com/hashicorp/go-hclog"
	"github.com/stretchr/testify/require"
)

func TestOpLogger(t *testing.T) {
	t.Run("queues operations until flushed", func(t *testing.T) {
		var out bytes.Buffer

		log := hclog.New(&hclog.LoggerOptions{
			Name:   "test",
			Level:  hclog.Info,
			Output: &out,
		})

		op := NewOpLogger(log)

		op.Debug("this is debug", "hole", 10)

		r := require.New(t)

		r.Equal(0, out.Len())

		op.Warn("this is a warning")

		r.NotEqual(0, out.Len())

		//x := out.Len()

		op.Flush()

		t.Log(out.String())

		r.Contains(out.String(), "this is debug: level=debug hole=10")
	})
}
