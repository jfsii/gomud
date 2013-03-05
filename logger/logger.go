package logger

import (
	"github.com/jfsherman/gomud/color"
	"log"
)

type Syslog struct{}

func New() *Syslog {
	return &Syslog{}
}

func (l *Syslog) Std(format string, v ...interface{}) {
	log.Printf(color.FgCyan+format+color.Reset, v...)
}
