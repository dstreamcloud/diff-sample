package logger

import "github.com/dstreamcloud/diff-sample/pkg/critical"

type Logger interface {
	Log(string) error
}

func init() {
	c := &critical.Critical{}
	c.DoNothing()
}
