package logger

import (
    "fmt"
    "io"
    "log"
    "os"
    "sync"
)

const (
    prefixInfo = "[I] "
    prefixWarn = "[W] "
    prefixErr = "[E] "
)

type Logger struct {
    *log.Logger
    mu sync.Mutex
}

func New() *Logger {
    return To(os.Stdout)
}

func To(w io.Writer) *Logger {
    return &Logger{
        // don't include log.Lshortfile or log.Llongfile - they are expensive and will hog the lock
        log.New(w, "", log.Ldate | log.Ltime),
        sync.Mutex{},
    }
}

func (l *Logger) Info(msg string) {
    l.mu.Lock()
    defer l.mu.Unlock()
    l.SetPrefix(prefixInfo)
    l.Println(msg)
}

func (l *Logger) Infof(format string, a ...interface{}) {
    l.mu.Lock()
    defer l.mu.Unlock()
    l.SetPrefix(prefixInfo)
    l.Println(fmt.Sprintf(format, a...))
}

func (l *Logger) Warn(msg string) {
    l.mu.Lock()
    defer l.mu.Unlock()
    l.SetPrefix(prefixWarn)
    l.Println(msg)
}

func (l *Logger) Warnf(format string, a ...interface{}) {
    l.mu.Lock()
    defer l.mu.Unlock()
    l.SetPrefix(prefixWarn)
    l.Println(fmt.Sprintf(format, a...))
}

func (l *Logger) Err(msg string) {
    l.mu.Lock()
    defer l.mu.Unlock()
    l.SetPrefix(prefixErr)
    l.Println(msg)
}

func (l *Logger) Errf(format string, a ...interface{}) {
    l.mu.Lock()
    defer l.mu.Unlock()
    l.SetPrefix(prefixErr)
    l.Println(fmt.Sprintf(format, a...))
}
