package tracer

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestTracer_New(t *testing.T) {

	tracer := New(LevelDebug)

	assert.NotNil(t, tracer)

}

func TestTracer_Level(t *testing.T) {
	tracer := New(LevelCritical)
	assert.Equal(t, LevelCritical, tracer.Level())

	tracer = New(LevelDebug)
	assert.Equal(t, LevelDebug, tracer.Level())

	tracer = New(LevelInfo)
	assert.Equal(t, LevelInfo, tracer.Level())
}

func TestTracer_LevelWithNil(t *testing.T) {

	var tr *Tracer = nil

	assert.Equal(t, LevelNothing, tr.Level())

}

func TestTracer_Should(t *testing.T) {

	var tr *Tracer = nil

	assert.False(t, tr.Should(LevelError))
	assert.False(t, tr.Should(LevelDebug))
	assert.False(t, tr.Should(LevelInfo))
	assert.False(t, tr.Should(LevelEverything))

	tr = New(LevelDebug)

	assert.True(t, tr.Should(LevelDebug))
	assert.False(t, tr.Should(LevelEverything))

}

func TestTracer_TraceLevels(t *testing.T) {

	tracer := New(LevelCritical)

	tracer.Trace(LevelDebug, "%s", "test")
	assert.NotEqual(t, len(tracer.data), 1)

	tracer.Trace(LevelInfo, "%s", "test")
	assert.NotEqual(t, len(tracer.data), 1)

	tracer.Trace(LevelWarning, "%s", "test")
	assert.NotEqual(t, len(tracer.data), 1)

	tracer.Trace(LevelError, "%s", "test")
	assert.NotEqual(t, len(tracer.data), 1)

	tracer.Trace(LevelCritical, "%s", "test")
	assert.Equal(t, len(tracer.data), 1)

	tracer = New(LevelDebug)

	tracer.Trace(LevelDebug, "%s", "test")
	assert.Equal(t, len(tracer.data), 1)

	tracer.Trace(LevelInfo, "%s", "test")
	assert.Equal(t, len(tracer.data), 2)

	tracer.Trace(LevelWarning, "%s", "test")
	assert.Equal(t, len(tracer.data), 3)

	tracer.Trace(LevelError, "%s", "test")
	assert.Equal(t, len(tracer.data), 4)

	tracer.Trace(LevelCritical, "%s", "test")
	assert.Equal(t, len(tracer.data), 5)

}

func TestTracer_Trace(t *testing.T) {

	tracer := New(LevelDebug)

	tracer.Trace(LevelDebug, "%s", "test")

	assert.Equal(t, tracer.data[0].Data, "test")
	assert.Equal(t, tracer.data[0].Level, LevelDebug)

}

func TestTracer_NilTracer(t *testing.T) {

	var tracer *Tracer = nil

	assert.NotPanics(t, func() {
		tracer.Trace(LevelDebug, "test")
		tracer.TraceDebug("test")
		tracer.TraceInfo("test")
		tracer.TraceWarning("test")
		tracer.TraceError("test")
		tracer.TraceCritical("test")
	})

}

func TestTracer_TraceHelpers(t *testing.T) {

	tracer := New(LevelDebug)
	tracer.TraceDebug("trace")
	tracer.TraceInfo("trace")
	tracer.TraceWarning("trace")
	tracer.TraceError("trace")
	tracer.TraceCritical("trace")
	assert.Equal(t, len(tracer.data), 5)

	tracer = New(LevelInfo)
	tracer.TraceDebug("trace")
	tracer.TraceInfo("trace")
	tracer.TraceWarning("trace")
	tracer.TraceError("trace")
	tracer.TraceCritical("trace")
	assert.Equal(t, len(tracer.data), 4)

	tracer = New(LevelWarning)
	tracer.TraceDebug("trace")
	tracer.TraceInfo("trace")
	tracer.TraceWarning("trace")
	tracer.TraceError("trace")
	tracer.TraceCritical("trace")
	assert.Equal(t, len(tracer.data), 3)

	tracer = New(LevelError)
	tracer.TraceDebug("trace")
	tracer.TraceInfo("trace")
	tracer.TraceWarning("trace")
	tracer.TraceError("trace")
	tracer.TraceCritical("trace")
	assert.Equal(t, len(tracer.data), 2)

	tracer = New(LevelCritical)
	tracer.TraceDebug("trace")
	tracer.TraceInfo("trace")
	tracer.TraceWarning("trace")
	tracer.TraceError("trace")
	tracer.TraceCritical("trace")
	assert.Equal(t, len(tracer.data), 1)

}

func TestTracer_TraceWithInvalidLevels(t *testing.T) {

	assert.Panics(t, func() {

		tracer := New(LevelDebug)

		tracer.Trace(LevelEverything, "%s", "test")

	}, "Trace LevelEverything")

	assert.Panics(t, func() {

		tracer := New(LevelDebug)

		tracer.Trace(LevelNothing, "%s", "test")

	}, "Trace LevelNothing")

	assert.Panics(t, func() {

		tracer := New(LevelDebug)

		tracer.Trace(LevelEverything-1, "%s", "test")

	}, "Trace too low")

	assert.Panics(t, func() {

		tracer := New(LevelDebug)

		tracer.Trace(LevelNothing+1, "%s", "test")

	}, "Trace too high")

}

func TestTracer_Copy(t *testing.T) {

	tracer := New(LevelDebug)

	tracer.Trace(LevelDebug, "%s", "test")

	temp := tracer.Data()

	assert.Equal(t, temp[0].Data, "test")

}

func TestTracer_Filter(t *testing.T) {

	tracer := New(LevelDebug)

	tracer.Trace(LevelDebug, "%s", "debug")
	tracer.Trace(LevelCritical, "%s", "critical")

	temp := tracer.Filter(LevelCritical)

	assert.Equal(t, temp[0].Data, "critical")

}
