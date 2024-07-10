package jfile

import (
	"fmt"
	"github.com/chroblert/jasync"
	"github.com/chroblert/jlog"
	"github.com/chroblert/jprogress"
	"math/rand"
	"path/filepath"
	"strings"
)

// SplitFile 对文件切片bRand: 是否开启随机分配
// splitCount: 分成多少份，最小为1
func SplitFile(srcFile, dstFile string, splitCount int, bRand bool) (err error) {
	if splitCount <= 0 {
		splitCount = 1
	}
	splitFMap := make(map[int]*jlog.FishLogger)
	for i := 0; i < splitCount; i++ {
		logFullPath := filepath.Join(filepath.Dir(dstFile), strings.TrimRight(filepath.Base(dstFile), filepath.Ext(dstFile))+fmt.Sprintf("-%d", i)+filepath.Ext(dstFile))
		tmpF := jlog.New(jlog.LogConfig{LogFullPath: logFullPath, StoreToFile: true, InitCreateNewLog: true, MaxSizePerLogFile: "-1"})
		splitFMap[i] = tmpF
	}
	totalCount := GetLineCount(srcFile)
	jprogress.Start()
	defer jprogress.Stop()
	bar := jprogress.Default(int(totalCount), "")
	pageCount := totalCount / int64(splitCount)
	_, _, err = ProcessLine64(srcFile, func(lineNum int64, line string) error {
		defer bar.Add(1)
		page := int((lineNum - 1) / pageCount)
		//jlog.Debug(lineNum, line)
		if page == splitCount {
			page -= 1
		}
		if bRand {
			page = rand.Intn(splitCount)
		}
		splitFMap[page].NInfof("%s\n", line)
		return JCONTINUE()
	}, false)
	if err != nil {
		//jlog.Errorf("%d,%s\n", i, err.Error())
		return
	}
	for i := 0; i < splitCount; i++ {
		splitFMap[i].Flush()
	}
	return
}

// AsyncSplitFile 并发对文件切片，bRand: 是否开启随机分配
// splitCount: 分成多少份，最小为1
// coCount: 协程数量
func AsyncSplitFile(srcFile, dstFile string, splitCount int, bRand bool, coCount int64) (err error) {
	if splitCount <= 0 {
		splitCount = 1
	}
	if coCount <= 0 {
		coCount = int64(splitCount)
	}
	splitFMap := make(map[int]*jlog.FishLogger)
	for i := 0; i < splitCount; i++ {
		logFullPath := filepath.Join(filepath.Dir(dstFile), strings.TrimRight(filepath.Base(dstFile), filepath.Ext(dstFile))+fmt.Sprintf("-%d", i)+filepath.Ext(dstFile))
		tmpF := jlog.New(jlog.LogConfig{LogFullPath: logFullPath, StoreToFile: true, InitCreateNewLog: true, MaxSizePerLogFile: "-1"})
		splitFMap[i] = tmpF
	}
	totalCount := GetLineCount(srcFile)
	jprogress.Start()
	defer jprogress.Stop()
	bar := jprogress.Default(int(totalCount), "")
	pageCount := totalCount / int64(splitCount)
	a := jasync.NewAR(coCount)
	_, _, err = ProcessLine64(srcFile, func(lineNum int64, line string) error {
		err = a.Init("").CAdd(func(inLineNum int64, inLine string) {
			defer bar.Add(1)
			page := int((inLineNum - 1) / pageCount)
			//jlog.Debug(lineNum, line)
			if page == splitCount {
				page -= 1
			}
			if bRand {
				page = rand.Intn(splitCount)
			}
			splitFMap[page].NInfof("%s\n", inLine)
		}, lineNum, line).CDO()
		if err != nil {
			return err
		}
		return JCONTINUE()
	}, false)
	if err != nil {
		//jlog.Errorf("%d,%s\n", i, err.Error())
		return
	}
	a.Wait()
	for i := 0; i < splitCount; i++ {
		splitFMap[i].Flush()
	}
	return
}
