package logger

import (
	"github.com/sirupsen/logrus"
	"os"
)

type Logger struct {
	logrus.Logger
}

func New(name string, logLevel uint32) *Logger {
	// setup the logrus global variable (just in case somebody uses it, but this should not
	// be recommended)
	logrus.SetLevel(logrus.Level(logLevel))
	logrus.SetReportCaller(true) // for having the filename and line number

	l := Logger{}
	l.SetOutput(os.Stderr)
	l.SetLevel(logrus.Level(logLevel))
	l.SetReportCaller(true)
	l.ExitFunc = os.Exit

	return &l
}
