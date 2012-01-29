package main

import (
    "log"
)

type Syslog struct {}

func NewSyslog() *Syslog {
    return &Syslog{}
}

func (l *Syslog) Std(format string, v ...interface{}) {
    log.Printf(FgCyan + format + Reset)
}
