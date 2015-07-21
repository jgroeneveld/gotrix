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
	if l.Tags != "" {
		format = l.Tags + " " + format
	}
	fmt.Fprintf(l.Writer, format+"\n", args...)
}

func (l *Logger) Fork(tags ...string) *Logger {
	newTags := ""
	if l.Tags != "" {
		newTags = l.Tags + " "
	}
	newTags += strings.Join(tags, " ")
	return &Logger{
		Writer: l.Writer,
		Tags:   newTags,
	}
}
