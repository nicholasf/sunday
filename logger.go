package sunday

import(
	"log"
)

type Level int

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

func NewLogger() Logger {
	l := &SimpleLogger{}
	return Logger(l)
}

//wraps calls to golang's standard logging library
type SimpleLogger struct {
	level Level
}

func (s SimpleLogger) Trace(msg string) {
	if s.level <= TRACE {
		log.Print("TRACE - ", msg)
	}
}

func (s SimpleLogger) Debug(msg string) {
	if s.level <= DEBUG {
		log.Print("DEBUG - ", msg)
	}
}

func (s SimpleLogger) Info(msg string) {
	if s.level <= INFO {
		log.Print("INFO - ", msg)
	}
}

func (s SimpleLogger) Error(msg string) {
	if s.level <= ERROR {
		log.Print("ERROR - ", msg)
	}
}

func (s SimpleLogger) Fatal(msg string) {
	if s.level <= FATAL {
		log.Print("FATAL - ", msg)
	}
}

func (s SimpleLogger) Level() Level {
	return s.level
}

func (s SimpleLogger) SetLevel(l Level) (e error) {
	s.level = l
	return
}
