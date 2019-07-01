package log

import (
	"fmt"
	"io"
	"log"
	"os"

	"github.com/neuronlabs/uni-logger"

	"github.com/neuronlabs/neuron/errors"
	"github.com/neuronlabs/neuron/errors/class"
)

const (
	// LDEBUG3 is the logger DEBUG3 level.
	LDEBUG3 = unilogger.DEBUG3

	// LDEBUG2 is the logger DEBUG2 level.
	LDEBUG2 = unilogger.DEBUG2

	// LDEBUG is the logger DEBUG level.
	LDEBUG = unilogger.DEBUG

	// LINFO is the logger INFO level.
	LINFO = unilogger.INFO

	// LWARNING is the logger WARNING level.
	LWARNING = unilogger.WARNING

	// LERROR is the logger ERROR level.
	LERROR = unilogger.ERROR

	// LCRITICAL is the logger CRITICAL level.
	LCRITICAL = unilogger.CRITICAL

	// LPRINT is the logger PRINT level.
	LPRINT = unilogger.PRINT

	// LUNKNOWN is the unspecified logger level.
	LUNKNOWN = unilogger.UNKNOWN
)

var (
	logger         unilogger.LeveledLogger
	currentLevel   = LINFO
	debugLeveled   unilogger.DebugLeveledLogger
	isDebugLeveled bool
)

// Level returns current logger Level.
func Level() unilogger.Level {
	return currentLevel
}

// Logger returns default logger.
func Logger() unilogger.LeveledLogger {
	return logger
}

// SetLevel sets the level if possible for the logger file.
func SetLevel(level unilogger.Level) error {
	currentLevel = level
	if logger == nil {
		return nil
	}

	lvl, ok := logger.(unilogger.LevelSetter)
	if !ok {
		return errors.New(class.CommonLoggerNotImplement, "logger doesn't implement LevelSetter interface")
	}

	lvl.SetLevel(currentLevel)

	for _, module := range modules {
		module.SetLevel(level)
	}

	return nil
}

// SetLogger sets the logger for the provided 'log' logger.
func SetLogger(log unilogger.LeveledLogger) {
	logger = log

	depth, ok := log.(unilogger.OutputDepthGetter)
	if ok {
		setter, ok := log.(unilogger.OutputDepthSetter)
		if ok {
			setter.SetOutputDepth(depth.GetOutputDepth() + 1)
		}
	}

	if lvlSetter, ok := log.(unilogger.LevelSetter); ok {
		lvlSetter.SetLevel(currentLevel)
	}

	debugLeveled, isDebugLeveled = log.(unilogger.DebugLeveledLogger)

	subLogger, isSubLogger := log.(unilogger.SubLogger)
	for _, m := range modules {
		if m.logger == nil && isSubLogger {
			m.logger = subLogger.SubLogger()
			m.initializeLogger()
		}
		m.SetLevel(currentLevel)
	}
}

// New creates new logger on the base of the provided values.
func New(out io.Writer, prefix string, flags int) {
	basic := unilogger.NewBasicLogger(out, prefix, flags)
	basic.SetOutputDepth(4)
	SetLogger(basic)
}

// Default creates default logger.
func Default() {
	basic := unilogger.NewBasicLogger(os.Stderr, "", log.Ldate|log.Ltime|log.Lshortfile)
	basic.SetOutputDepth(4)
	SetLogger(basic)
}

// Debug3f writes the formated debug log.
func Debug3f(format string, args ...interface{}) {
	if !isDebugLeveled {
		if logger != nil {
			logger.Debugf(format, args...)
		}
	} else {
		if debugLeveled != nil {
			debugLeveled.Debug3f(format, args...)
		}
	}
}

// Debug2f writes the formated debug log.
func Debug2f(format string, args ...interface{}) {
	if !isDebugLeveled {
		if logger != nil {
			logger.Debugf(format, args...)
		}
	} else {
		if debugLeveled != nil {
			debugLeveled.Debug2f(format, args...)
		}
	}
}

// Debugf writes the formated debug log.
func Debugf(format string, args ...interface{}) {
	if logger != nil {
		logger.Debugf(format, args...)
	}
}

// Infof writes the formated info log.
func Infof(format string, args ...interface{}) {
	if logger != nil {
		logger.Infof(format, args...)
	}
}

// Warningf writes the formated warning log.
func Warningf(format string, args ...interface{}) {
	if logger != nil {
		logger.Warningf(format, args...)
	}
}

// Errorf writes the formated error log.
func Errorf(format string, args ...interface{}) {
	if logger != nil {
		logger.Errorf(format, args...)
	}
}

// Fatalf writes the formated fatal log.
func Fatalf(format string, args ...interface{}) {
	if logger != nil {
		logger.Fatalf(format, args...)
	} else {
		fmt.Printf(format, args...)
		os.Exit(1)
	}
}

// Panicf writes the formated panic log.
func Panicf(format string, args ...interface{}) {
	if logger != nil {
		logger.Panicf(format, args...)
	} else {
		panic(fmt.Sprintf(format, args...))
	}
}

// Debug3 writes the debug3 level logs.
func Debug3(args ...interface{}) {
	if !isDebugLeveled {
		if logger != nil {
			logger.Debug(args...)
		}
	} else {
		if debugLeveled != nil {
			debugLeveled.Debug3(args...)
		}
	}
}

// Debug2 writes the debug2 level logs.
func Debug2(args ...interface{}) {
	if !isDebugLeveled {
		if logger != nil {
			logger.Debug(args...)
		}
	} else {
		if debugLeveled != nil {
			debugLeveled.Debug2(args...)
		}
	}
}

// Debug writes the  debug log.
func Debug(args ...interface{}) {
	if logger != nil {
		logger.Debug(args...)
	}
}

// Info writes the info log.
func Info(args ...interface{}) {
	if logger != nil {
		logger.Info(args...)
	}
}

// Warning writes the warning log.
func Warning(args ...interface{}) {
	if logger != nil {
		logger.Warning(args...)
	}
}

// Error writes the error log.
func Error(args ...interface{}) {
	if logger != nil {
		logger.Error(args...)
	}
}

// Fatal writes the fatal log.
func Fatal(args ...interface{}) {
	if logger != nil {
		logger.Fatal(args...)
	}
}

// Panic writes the panic log.
func Panic(args ...interface{}) {
	if logger != nil {
		logger.Panic(args...)
	} else {
		panic(fmt.Sprint(args...))
	}
}
