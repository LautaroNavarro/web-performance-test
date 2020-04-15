package main

import (
    "log"
    "os"
    "sync"
    "io"
)

type logger struct {
    filename string
    *log.Logger
}

var l *logger
var once sync.Once

func getLogger() *logger {
    once.Do(func() {
        l = createLogger("performance-test.log")
    })
    return l
}

func createLogger(fname string) *logger {
    file, _ := os.OpenFile(fname, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0666)
    logWriter := io.MultiWriter(os.Stdout, file)

    return &logger{
        filename: fname,
        Logger:   log.New(logWriter, "[DEBUG] ", log.Lshortfile),
    }
}