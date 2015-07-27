package logger

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"strings"
)

var Discard = &TaggedLogger{writer: ioutil.Discard}

var DefaultWriter = os.Stderr

type Logger interface {
	Printf(format string, args ...interface{})
	Fork(tags ...string) Logger
}

func New(tags ...string) *TaggedLogger {
	return &TaggedLogger{
		writer: DefaultWriter,
		tags:   strings.Join(tags, " "),
	}
}

type TaggedLogger struct {
	writer io.Writer
	tags   string
}

func (l *TaggedLogger) Printf(format string, args ...interface{}) {
	if l.tags != "" {
		format = l.tags + " " + format
	}
	fmt.Fprintf(l.writer, format+"\n", args...)
}

func (l *TaggedLogger) Fork(tags ...string) Logger {
	newTags := ""
	if l.tags != "" {
		newTags = l.tags + " "
	}
	newTags += strings.Join(tags, " ")
	return &TaggedLogger{
		writer: l.writer,
		tags:   newTags,
	}
}
