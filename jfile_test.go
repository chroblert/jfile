package jfile

import (
	"github.com/chroblert/jlog"
	"testing"
)

func TestGetFileSize(t *testing.T) {
	jlog.Info(GetFileSize("jfile.go"))
	ProcessLine("test2.txt", func(line_num int, line string) error {
		jlog.Infof("%d:%s_\n", line_num, line)
		return JCONTINUE()
	}, false)
}

func TestProcessLineReverse(t *testing.T) {
	defer jlog.Flush()
	jlog.SetLogFullPath("logs/app2.log")
	jlog.Info(ProcessLineReverse("test2.txt", func(line_num int64, line string) error {
		jlog.Infof("%d:%s_\n", line_num, line)
		return JCONTINUE()
	}, false))
}
