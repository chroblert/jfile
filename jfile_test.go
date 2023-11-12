package jfile

import (
	"github.com/chroblert/jlog"
	"testing"
)

func TestGetFileSize(t *testing.T) {
	jlog.Info(GetFileSize("jfile.go"))
}
