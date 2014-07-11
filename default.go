package logging

import (
    stdlog "log"
    "os"
)

const (
    default_id = "default-logger"
)

func InitDefault(level Level) *Logger {
    default_logger := MustGetLogger(default_id)
    format := MustStringFormatter("%{level} %{message}")
    SetFormatter(format)
    // Setup one stdout and one syslog backend.
    logBackend := NewLogBackend(os.Stdout, "", stdlog.LstdFlags|stdlog.Lshortfile)
    logBackend.Color = true

    syslogBackend, err := NewSyslogBackend("")
    if err != nil {
        Default.Fatal(err)
    }

    // Combine them both into one logging backend.
    SetBackend(logBackend, syslogBackend)
    SetLevel(level, default_id)
    return default_logger
}