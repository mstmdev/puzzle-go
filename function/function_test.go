package function

import (
	"fmt"
	"io"
	"os"
	"strings"
	"testing"
)

type Logger interface {
	New(w io.Writer) Logger
	Debug(s string) Logger
	Info(s string) Logger
}

type ConsoleLogger struct {
	w io.Writer
}

func (cl ConsoleLogger) New(w io.Writer) Logger {
	cl.w = w
	return cl
}

func (cl ConsoleLogger) Debug(s string) Logger {
	fmt.Fprintln(cl.w, "[ConsoleLogger][Debug]", s)
	return cl
}

func (cl ConsoleLogger) Info(s string) Logger {
	fmt.Fprintln(cl.w, "[ConsoleLogger][Info]", s)
	return cl
}

type FileLogger struct {
	w io.Writer
}

func (fl *FileLogger) New(w io.Writer) Logger {
	fl.w = w
	return fl
}

func (fl *FileLogger) Debug(s string) Logger {
	fmt.Fprintln(fl.w, "[FileLogger][Debug]", s)
	return fl
}

func (fl *FileLogger) Info(s string) Logger {
	fmt.Fprintln(fl.w, "[FileLogger][Info]", s)
	return fl
}

func writer() (io.Writer, *strings.Builder) {
	recorder := &strings.Builder{}
	w := io.MultiWriter(os.Stdout, recorder)
	return w, recorder
}

func TestLogger(t *testing.T) {
	w, recorder := writer()
	Logger.New(ConsoleLogger{}, w).Debug("hello").Info("world")
	expect := "[ConsoleLogger][Debug] hello\n[ConsoleLogger][Info] world\n"
	actual := recorder.String()
	if actual != expect {
		t.Errorf("expect get %s but actual get %s", expect, actual)
	}

	recorder.Reset()
	Logger.New(&FileLogger{}, w).Debug("hello").Info("world")
	expect = "[FileLogger][Debug] hello\n[FileLogger][Info] world\n"
	actual = recorder.String()
	if actual != expect {
		t.Errorf("expect get %s but actual get %s", expect, actual)
	}
}

func TestConsoleLogger(t *testing.T) {
	w, recorder := writer()
	ConsoleLogger.New(ConsoleLogger{}, w).Debug("hello").Info("world")
	expect := "[ConsoleLogger][Debug] hello\n[ConsoleLogger][Info] world\n"
	actual := recorder.String()
	if actual != expect {
		t.Errorf("expect get %s but actual get %s", expect, actual)
	}
}

func TestFileLogger(t *testing.T) {
	w, recorder := writer()
	(*FileLogger).New(&FileLogger{}, w).Debug("hello").Info("world")
	expect := "[FileLogger][Debug] hello\n[FileLogger][Info] world\n"
	actual := recorder.String()
	if actual != expect {
		t.Errorf("expect get %s but actual get %s", expect, actual)
	}
}
