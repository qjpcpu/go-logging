package main

import "github.com/qjpcpu/go-logging"

func main() {
    logging.InitLogger(logging.DEBUG)
    logging.Debug("hee")
    logging.Info("heelo")
    logging.Error("heelo")
    logging.Critical("%s", "heelo")
    logging.Debugf("%s", "hee")
    logging.Infof("%s", "heelo")
    logging.Errorf("%s", "heelo")
    logging.Criticalf("%s", "heelo")
    logging.Warningf("%s", "heelo")
}
