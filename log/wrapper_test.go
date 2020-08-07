package log

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type stdlogger struct{}

func (s *stdlogger) Print(...interface{})          {}
func (s *stdlogger) Printf(string, ...interface{}) {}
func (s *stdlogger) Println(...interface{})        {}
func (s *stdlogger) Panic(...interface{})          {}
func (s *stdlogger) Panicf(string, ...interface{}) {}
func (s *stdlogger) Panicln(...interface{})        {}
func (s *stdlogger) Fatal(...interface{})          {}
func (s *stdlogger) Fatalf(string, ...interface{}) {}
func (s *stdlogger) Fatalln(...interface{})        {}

type leveledLogger struct{}

func (l *leveledLogger) Debugf(string, ...interface{})   {}
func (l *leveledLogger) Infof(string, ...interface{})    {}
func (l *leveledLogger) Warningf(string, ...interface{}) {}
func (l *leveledLogger) Errorf(string, ...interface{})   {}
func (l *leveledLogger) Fatalf(string, ...interface{})   {}
func (l *leveledLogger) Panicf(string, ...interface{})   {}
func (l *leveledLogger) Debug(...interface{})            {}
func (l *leveledLogger) Info(...interface{})             {}
func (l *leveledLogger) Warning(...interface{})          {}
func (l *leveledLogger) Error(...interface{})            {}
func (l *leveledLogger) Fatal(...interface{})            {}
func (l *leveledLogger) Panic(...interface{})            {}

type shortLeveledLogger struct{}

func (c *shortLeveledLogger) Debugf(string, ...interface{}) {}
func (c *shortLeveledLogger) Infof(string, ...interface{})  {}
func (c *shortLeveledLogger) Warnf(string, ...interface{})  {}
func (c *shortLeveledLogger) Errorf(string, ...interface{}) {}
func (c *shortLeveledLogger) Fatalf(string, ...interface{}) {}
func (c *shortLeveledLogger) Panicf(string, ...interface{}) {}
func (c *shortLeveledLogger) Debug(...interface{})          {}
func (c *shortLeveledLogger) Info(...interface{})           {}
func (c *shortLeveledLogger) Warn(...interface{})           {}
func (c *shortLeveledLogger) Error(...interface{})          {}
func (c *shortLeveledLogger) Fatal(...interface{})          {}
func (c *shortLeveledLogger) Panic(...interface{})          {}

type extendedLogger struct {
	leveledLogger
}

func (e *extendedLogger) Print(...interface{})          {}
func (e *extendedLogger) Printf(string, ...interface{}) {}
func (e *extendedLogger) Println(...interface{})        {}
func (e *extendedLogger) Debugln(...interface{})        {}
func (e *extendedLogger) Infoln(...interface{})         {}
func (e *extendedLogger) Warningln(...interface{})      {}
func (e *extendedLogger) Errorln(...interface{})        {}
func (e *extendedLogger) Fatalln(...interface{})        {}
func (e *extendedLogger) Panicln(...interface{})        {}

type nonLogger struct{}

// TestNewLoggerWrapper tests NewLoggerWrapper function.
func TestNewLoggerWrapper(t *testing.T) {
	loggers := []interface{}{&stdlogger{}, &leveledLogger{}, &shortLeveledLogger{}, &extendedLogger{}}

	t.Run("NewWrapper", func(t *testing.T) {
		for _, logger := range loggers {
			wrapper, err := NewLoggerWrapper(logger)
			assert.IsType(t, &Wrapper{}, wrapper)
			assert.NoError(t, err)
			wrapper = MustGetLoggerWrapper(logger)
			assert.IsType(t, &Wrapper{}, wrapper)
		}

		var args []interface{}
		format := "some format"
		for _, logger := range loggers {
			wrapper := MustGetLoggerWrapper(logger)
			wrapper.Print(args)
			wrapper.Printf(format, args)
			wrapper.Println(args)

			wrapper.Debug(args)
			wrapper.Debugf(format, args...)
			wrapper.Debugln(args)

			wrapper.Info(args)
			wrapper.Infof(format, args...)
			wrapper.Infoln(args)

			wrapper.Warning(args)
			wrapper.Warningf(format, args...)
			wrapper.Warningln(args)

			wrapper.Error(args)
			wrapper.Errorf(format, args)
			wrapper.Errorln(args)

			wrapper.Fatal(args)
			wrapper.Fatalf(format, args)
			wrapper.Fatalln(args)

			wrapper.Panic(args)
			wrapper.Panicf(format, args)
			wrapper.Panicln(args)
		}
	})

	t.Run("NotImplement", func(t *testing.T) {
		unknownLogger := nonLogger{}
		wrapper, err := NewLoggerWrapper(unknownLogger)
		assert.Error(t, err)
		assert.Nil(t, wrapper)
		assert.Panics(t, func() { MustGetLoggerWrapper(unknownLogger) })
	})
}

// TestBuildLeveled tests the buildLeveled function.
func TestBuildLeveled(t *testing.T) {
	level := LevelDebug
	format := "some format"
	arguments := []interface{}{"First", "Second"}

	// Providing nil format should add level as first argument to args
	t.Run("NilFormat", func(t *testing.T) {
		args := buildLeveled(level, nil, arguments...)
		assert.Equal(t, fmt.Sprintf("%s: ", level), args[0])
	})

	t.Run("Formatted", func(t *testing.T) {
		thisFormat := format
		buildLeveled(level, &thisFormat, arguments...)
		assert.NotEqual(t, format, thisFormat)
	})
}
