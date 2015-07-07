package logger

import (
	"fmt"
	"io"
	"os"
	"strings"
)

var DefaultWriter = os.Stderr

func New(tags ...string) *Logger {
	return &Logger{
		Writer: DefaultWriter,
		Tags:   strings.Join(tags, " "),
	}
}

type Logger struct {
	Writer io.Writer
	Tags   string
}

func (l *Logger) Printf(format string, args ...interface{}) {
	fmt.Fprintf(l.Writer, format, args...)
}

func (l *Logger) Fork(tags ...string) *Logger {
	return &Logger{
		Writer: l.Writer,
		Tags:   l.Tags + " " + strings.Join(tags, " "),
	}
}
