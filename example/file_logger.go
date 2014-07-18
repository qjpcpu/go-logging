package main

import (
    "fmt"
    "github.com/qjpcpu/go-logging"
    "io/ioutil"
    "os"
)

func main() {
    filename := "/tmp/log-file"
    logging.InitSimpleFileLogger(filename, logging.DEBUG)
    logging.Debug("hello")
    logging.Info("hi")
    logging.Errorf("hello %s", "jason")
    logging.Criticalf("hi %s", "jason")

    data, _ := ioutil.ReadFile(filename)
    fmt.Printf("the logfile content is:\n%s", string(data))
    os.Remove(filename)
}
