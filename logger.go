package sunday

import(
	"log"
)

type Level int

var Log Logger

const(
	TRACE Level = iota
	DEBUG
	INFO
	ERROR
	FATAL 
)

type Logger interface {
	Trace(msg string)
	Debug(msg string)
	Info(msg string)
	Error(msg string)
	Fatal(msg string)
	SetLevel(level Level) error
	Level() Level
}

//wraps calls to golang's standard logging library
type logger struct {
	level Level
}

func NewLogger() Logger {
	l := &logger{}
	return Logger(l)
}

func (l logger) Trace(msg string) {
	if l.level <= TRACE {
		log.Print("TRACE - ", msg)
	}
}

func (l logger) Debug(msg string) {
	if l.level <= DEBUG {
		log.Print("DEBUG - ", msg)
	}
}

func (l logger) Info(msg string) {
	if l.level <= INFO {
		log.Print("INFO - ", msg)
	}
}

func (l logger) Error(msg string) {
	if l.level <= ERROR {
		log.Print("ERROR - ", msg)
	}
}

func (l logger) Fatal(msg string) {
	if l.level <= FATAL {
		log.Print("FATAL - ", msg)
	}
}

func (l logger) Level() Level {
	return l.level
}

func (l logger) SetLevel(le Level) (e error) {
	l.level = le
	return
}