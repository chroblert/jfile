package jfile

import (
	"github.com/chroblert/jasync"
	"github.com/chroblert/jlog"
)

// ADiffB A文件相对于B文件的差集（A中有但B中没有的行）
func ADiffB(AFile, BFile, dstFile string, coroutinCount int) (diffCount int, err error) {
	dst_f := jlog.New(jlog.LogConfig{LogFullPath: dstFile, InitCreateNewLog: true, StoreToFile: true})
	defer dst_f.Flush()
	diffCount = 0
	a := jasync.NewAR(int64(coroutinCount))
	_, _, err = ProcessLine64(AFile, func(line_num int64, line string) error {
		a.Init("1").CAdd(func(line string) {
			bFinded := false
			_, _, err = ProcessLine64(BFile, func(b_line_num int64, b_line string) error {
				if line == b_line {
					bFinded = true
					return JBREAK()
				}
				if err != nil {
					return err
				}
				return JCONTINUE()
			}, false)
			if !bFinded {
				diffCount += 1
				dst_f.NInfo(line)
			}
		}, line).CDO()
		return JCONTINUE()
	}, false)
	a.Wait()
	dst_f.Flush()
	if err != nil {
		return 0, err
	}
	return
}
