package logging

type LogLevel int32

const (
	DEBUG LogLevel = 1
	INFO  LogLevel = 2
	WARN  LogLevel = 3
	ERROR LogLevel = 4
	FATAL LogLevel = 5
	PANIC LogLevel = 6
)

type Logger interface {
	SetLevel(level LogLevel)
	Enable(level LogLevel) bool
	Debug(template string, args ...interface{})
	Info(template string, args ...interface{})
	Warn(template string, args ...interface{})
	Error(template string, args ...interface{})
	Fatal(template string, args ...interface{})
	Panic(template string, args ...interface{})
}
