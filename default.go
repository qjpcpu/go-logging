package logging

import (
    "code.google.com/p/go.exp/inotify"
    "fmt"
    stdlog "log"
    "os"
    "strings"
)

const (
    default_id = "default-logger"
)

var default_logger *Logger

func InitLogger(level Level) {
    default_logger = MustGetLogger(default_id)
    format := MustStringFormatter("%{level} %{message}")
    SetFormatter(format)
    // Setup one stdout and one syslog backend.
    logBackend := NewLogBackend(os.Stdout, "", stdlog.LstdFlags|stdlog.Lshortfile)
    logBackend.Color = true

    syslogBackend, err := NewSyslogBackend("")
    if err != nil {
        default_logger.Fatal(err)
    }
    // Combine them both into one logging backend.
    SetBackend(logBackend, syslogBackend)
    SetLevel(level, default_id)
}
func InitSimpleLogger(level Level) {
    default_logger = MustGetLogger(default_id)
    format := MustStringFormatter("%{level} %{message}")
    SetFormatter(format)
    SetLevel(level, default_id)
}
func InitFileLogger(filename string, level Level) {
    initFileLogger(filename, level, stdlog.LstdFlags|stdlog.Lshortfile)
}
func InitSimpleFileLogger(filename string, level Level) {
    initFileLogger(filename, level, stdlog.LstdFlags)
}
func initFileLogger(filename string, level Level, flags int) {
    default_logger = MustGetLogger(default_id)
    format := MustStringFormatter("%{level} %{message}")
    SetFormatter(format)
    writer, _ := os.OpenFile(filename, os.O_RDWR|os.O_APPEND|os.O_CREATE, 0660)
    logBackend := NewLogBackend(writer, "", flags)
    SetBackend(logBackend)
    SetLevel(level, default_id)
    go func() {
        watcher, err := inotify.NewWatcher()
        if err != nil {
            return
        }
        defer watcher.Close()
        for {
            err = watcher.AddWatch(filename, inotify.IN_ATTRIB|inotify.IN_DELETE_SELF|inotify.IN_MOVE)
            if err != nil {
                fmt.Println(err)
                return
            }
            for {
                <-watcher.Event
                if _, err = os.Stat(filename); os.IsNotExist(err) {
                    break
                }
            }
            writer.Close()
            writer, _ = os.OpenFile(filename, os.O_RDWR|os.O_APPEND|os.O_CREATE, 0660)
            logBackend := NewLogBackend(writer, "", flags)
            SetBackend(logBackend)
        }
    }()
}

func SetLogger(logger *Logger) {
    default_logger = logger
}

// Fatal is equivalent to l.Critical(fmt.Sprint()) followed by a call to os.Exit(1).
func Fatal(args ...interface{}) {
    if default_logger == nil {
        return
    }
    s := strings.TrimRight(fmt.Sprintln(args...), "\n")
    default_logger.log(CRITICAL, "%s", s)
    os.Exit(1)
}

// Fatalf is equivalent to l.Critical followed by a call to os.Exit(1).
func Fatalf(format string, args ...interface{}) {
    if default_logger == nil {
        return
    }
    default_logger.log(CRITICAL, format, args...)
    os.Exit(1)
}

// Panic is equivalent to l.Critical(fmt.Sprint()) followed by a call to panic().
func Panic(args ...interface{}) {
    if default_logger == nil {
        return
    }
    s := strings.TrimRight(fmt.Sprintln(args...), "\n")
    default_logger.log(CRITICAL, "%s", s)
    panic(s)
}

// Panicf is equivalent to l.Critical followed by a call to panic().
func Panicf(format string, args ...interface{}) {
    if default_logger == nil {
        return
    }
    s := fmt.Sprintf(format, args...)
    default_logger.log(CRITICAL, "%s", s)
    panic(s)
}

// Criticadefault_logger.logs a message using CRITICAL as log level.
func Critical(args ...interface{}) {
    if default_logger == nil {
        return
    }
    default_logger.log(CRITICAL, "%s", strings.TrimRight(fmt.Sprintln(args...), "\n"))
}
func Criticalf(format string, args ...interface{}) {
    if default_logger == nil {
        return
    }
    default_logger.log(CRITICAL, "%s", fmt.Sprintf(format, args...))
}

// Error logs a message using INFO as log level.
func Error(args ...interface{}) {
    if default_logger == nil {
        return
    }
    default_logger.log(ERROR, "%s", strings.TrimRight(fmt.Sprintln(args...), "\n"))
}
func Errorf(format string, args ...interface{}) {
    if default_logger == nil {
        return
    }
    default_logger.log(ERROR, "%s", fmt.Sprintf(format, args...))
}

// Warning logs a message using WARNING as log level.
func Warning(args ...interface{}) {
    if default_logger == nil {
        return
    }
    default_logger.log(WARNING, "%s", strings.TrimRight(fmt.Sprintln(args...), "\n"))
}
func Warningf(format string, args ...interface{}) {
    if default_logger == nil {
        return
    }
    default_logger.log(WARNING, "%s", fmt.Sprintf(format, args...))
}

// Notice logs a message using INFO as log level.
func Notice(args ...interface{}) {
    if default_logger == nil {
        return
    }
    default_logger.log(NOTICE, "%s", strings.TrimRight(fmt.Sprintln(args...), "\n"))
}
func Noticef(format string, args ...interface{}) {
    if default_logger == nil {
        return
    }
    default_logger.log(NOTICE, "%s", fmt.Sprintf(format, args...))
}

// Info logs a message using INFO as log level.
func Info(args ...interface{}) {
    if default_logger == nil {
        return
    }
    default_logger.log(INFO, "%s", strings.TrimRight(fmt.Sprintln(args...), "\n"))
}
func Infof(format string, args ...interface{}) {
    if default_logger == nil {
        return
    }
    default_logger.log(INFO, "%s", fmt.Sprintf(format, args...))
}

// Debug logs a message using DEBUG as log level.
func Debug(args ...interface{}) {
    if default_logger == nil {
        return
    }
    default_logger.log(DEBUG, "%s", strings.TrimRight(fmt.Sprintln(args...), "\n"))
}
func Debugf(format string, args ...interface{}) {
    if default_logger == nil {
        return
    }
    default_logger.log(DEBUG, "%s", fmt.Sprintf(format, args...))
}
