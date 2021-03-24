package zaplogger

import (
	"os"

	"github.com/hopex-hk/go_sdk/core/logging"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var sugarLogger *zap.SugaredLogger
var atomicLevel zap.AtomicLevel

type ZapLogger struct {
}

func init() {
	encoderCfg := zapcore.EncoderConfig{
		TimeKey:     "time",
		MessageKey:  "msg",
		LevelKey:    "level",
		EncodeLevel: zapcore.CapitalColorLevelEncoder,
		EncodeTime:  zapcore.ISO8601TimeEncoder,
	}

	// define default level as debug level
	atomicLevel = zap.NewAtomicLevel()
	atomicLevel.SetLevel(zapcore.DebugLevel)

	core := zapcore.NewCore(zapcore.NewConsoleEncoder(encoderCfg), os.Stdout, atomicLevel)
	sugarLogger = zap.New(core).Sugar()
}

func toZaplevel(level logging.LogLevel) zapcore.Level {
	var zaplevel zapcore.Level
	switch level {
	case logging.DEBUG:
		zaplevel = zapcore.DebugLevel
	case logging.INFO:
		zaplevel = zapcore.InfoLevel
	case logging.WARN:
		zaplevel = zapcore.WarnLevel
	case logging.ERROR:
		zaplevel = zapcore.ErrorLevel
	case logging.FATAL:
		zaplevel = zapcore.FatalLevel
	case logging.PANIC:
		zaplevel = zapcore.PanicLevel
	default:
		zaplevel = zapcore.DebugLevel
	}
	return zaplevel
}

func SetLevel(level logging.LogLevel) {
	atomicLevel.SetLevel(toZaplevel(level))
}

func (logger *ZapLogger) Enable(level logging.LogLevel) bool {
	return atomicLevel.Enabled(toZaplevel(level))
}

func (logger *ZapLogger) SetLevel(level logging.LogLevel) {
	atomicLevel.SetLevel(toZaplevel(level))
}

func (logger *ZapLogger) Fatal(template string, args ...interface{}) {
	sugarLogger.Fatalf(template, args...)
}

func (logger *ZapLogger) Error(template string, args ...interface{}) {
	sugarLogger.Errorf(template, args...)
}

func (logger *ZapLogger) Panic(template string, args ...interface{}) {
	sugarLogger.Panicf(template, args...)
}

func (logger *ZapLogger) Warn(template string, args ...interface{}) {
	sugarLogger.Warnf(template, args...)
}

func (logger *ZapLogger) Info(template string, args ...interface{}) {
	sugarLogger.Infof(template, args...)
}

func (logger *ZapLogger) Debug(template string, args ...interface{}) {
	sugarLogger.Debugf(template, args...)
}
