## Golang logging library

This package is simplified from https://github.com/op/go-logging.

## Installing

### Using *go get*

    $ go get github.com/qjpcpu/go-logging


## Example

Print log to stdout

```go
package main

import "github.com/qjpcpu/go-logging"

func main() {
    logging.InitLogger(logging.DEBUG)
	log.Debug("hello", "golang")
	log.Info("hello", "golang")
	log.Debugf("hello %s", "golang")
	log.Infof("hello %s", "golang")
}
```

Print log to file

```go
package main

import "github.com/qjpcpu/go-logging"

func main() {
    logging.InitSimpleFileLogger("/tmp/logger",logging.DEBUG)
	log.Debug("hello", "golang")
	log.Info("hello", "golang")
	log.Debugf("hello %s", "golang")
	log.Infof("hello %s", "golang")
}
```
## FYI

This pkg is forked form https://github.com/op/go-logging, thanks for his excellent work. For more docs, please see http://godoc.org/github.com/op/go-logging.

