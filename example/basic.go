package main

import (
    "github.com/qjpcpu/go-logging"
)

func main() {
    logging.InitLogger(logging.DEBUG)
    logging.Debug("hello")
    logging.Info("hi")
    logging.Errorf("hello %s", "jason")
    logging.Criticalf("hi %s", "jason")
}
