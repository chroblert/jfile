package jfile

import (
	"fmt"
	"github.com/chroblert/jasync"
	"os"
)

// 合并多个文件到一个文件
// TODO 返回count
func Merge(dstFile string, coroutineCount int, srcFileList ...string) (count int64, err error) {
	if len(srcFileList) == 0 {
		return 0, fmt.Errorf("no srcFileList")
	}
	dst_f, err := os.OpenFile(dstFile, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0666)
	if err != nil {
		return 0, err
	}
	defer dst_f.Close()
	count = 0
	if coroutineCount < 1 {
		coroutineCount = 1
	}
	a := jasync.NewAR(int64(coroutineCount))
	//var resChan = make(chan struct{})
	//var taskCount = 0
	//var taskChan = make(chan struct{})
	for _, srcFile := range srcFileList {
		err = a.Init("1").CAdd(func(srcFile string) {
			_, _, err = ProcessLine(srcFile, func(line_num int, line string) error {
				dst_f.Write([]byte(line + "\n"))
				count += 1
				//resChan <- struct{}{}
				return JCONTINUE()
			}, false)
			//taskChan <- struct{}{}
			return
		}, srcFile).CDO()
		if err != nil {
			return 0, err
		}
	}
	a.Wait()
	//for {
	//	select {
	//	case <-taskChan:
	//	case <-resChan:
	//
	//	}
	//}
	return

}
