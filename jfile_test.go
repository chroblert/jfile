package jfile

import (
	"github.com/chroblert/jlog"
	"testing"
	"time"
)

func TestGetFileSize(t *testing.T) {
	jlog.Info(GetFileSize("jfile.go"))
	ProcessLine64("test2.txt", func(line_num int64, line string) error {
		jlog.Infof("%d:%s_\n", line_num, line)
		return JCONTINUE()
	}, false)
}

func TestProcessLineReverse(t *testing.T) {
	defer jlog.Flush()
	jlog.SetLogFullPath("logs/app2.log")
	jlog.Info(ProcessLineReverse64("test2.txt", func(line_num int64, line string) error {
		jlog.Infof("%d:%s_\n", line_num, line)
		return JCONTINUE()
	}, false))
}

func TestGetLineCount(t *testing.T) {
	filePath := "C:\\T00ls\\01-GoProject\\HttpUtils\\core\\fingerprint\\assets_export_jz_domainscan_app3_goby.jsonl"
	//var count int64
	//timeStart := time.Now()
	//_, _, err := ProcessLine64(filePath, func(lineNum int64, line string) error {
	//	count = int64(lineNum)
	//	return JCONTINUE()
	//}, false)
	//if err != nil {
	//	return
	//}
	jlog.Debug(GetLineCount(filePath))
	return
}

func TestGetLineCount2(t *testing.T) {
	filePath := "C:\\T00ls\\01-GoProject\\HttpUtils\\core\\fingerprint\\assets_export_jz_domainscan_app3_goby.jsonl"
	//filePath = "C:\\T00ls\\01-GoProject\\HttpUtils\\core\\fingerprint\\assets_export_jz_domainscan.jsonl"
	var count int64
	timeStart := time.Now()
	//a := jasync.NewAR(100)
	_, allLineNum, err := ProcessLine64(filePath, func(lineNum int64, line string) error {

		return JCONTINUE()
	}, false)
	if err != nil {
		return
	}
	jlog.Debug(count, allLineNum, time.Now().Sub(timeStart).Seconds())
	return
}
