package main

import (
	"flag"
	"github.com/dstreamcloud/diff-sample/pkg/logger"
	"github.com/dstreamcloud/diff-sample/pkg/sirupsenlogrus"
	"github.com/dstreamcloud/diff-sample/pkg/uberzap"
)

var pLogger = flag.String("logger", "logrus", "logger")

func main() {
	var l logger.Logger
	if *pLogger == "zap" {
		l = uberzap.New()
	} else {
		l = sirupsenlogrus.New()
	}
	l.Log("hello")
}
