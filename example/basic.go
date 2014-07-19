package main

import (
    "github.com/qjpcpu/go-logging"
)

func main() {
    logging.InitLogger(logging.DEBUG)
    logging.Debug("hello")
    logging.Info("hi")
    logging.Notice("hi")
    logging.Warning("hi")
    logging.Error("hi")
    logging.Critical("hi")
    logging.Debugf("%s", "hi")
    logging.Infof("%s", "hi")
    logging.Noticef("%s", "hi")
    logging.Warningf("%s", "hi")
    logging.Errorf("hello %s", "jason")
    logging.Criticalf("hi %s", "jason")
    logging.Fatalf("hi %s, fatal and exit.", "jason")
}
