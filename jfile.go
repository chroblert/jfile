package jfile

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

// 判断传入的文件或目录是否存在
func PathExists(path string) (bool, error) {
	path, err := GetAbsPath(path)
	if err != nil {
		return false, err
	}
	_, err = os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

// 获取当前运行的可执行文件的路径
func GetWorkPath() (string, error) {
	file, err := exec.LookPath(os.Args[0])
	if err != nil {
		return "", err
	}
	path, err := filepath.Abs(file)
	if err != nil {
		return "", err
	}
	index := strings.LastIndex(path, string(os.PathSeparator))
	return path[:index], nil
}

// 获取绝对路径
// 若传入的参数是绝对路径，则返回
// 若是相对路径，则将其拼接到当前的工作目录，并返回
func GetAbsPath(path string) (string, error) {
	if !filepath.IsAbs(path) {
		workPath, err := GetWorkPath()
		if err != nil {
			return "", err
		}
		path = filepath.FromSlash(workPath + "/" + path)
	}
	return filepath.Abs(path)
}

// 枚举某个目录下所有的文件
func GetFilenamesByDir(root string) ([]string, error) {
	root, err := GetAbsPath(root)
	if err != nil {
		return nil, err
	}
	var files []string
	fileInfo, err := ioutil.ReadDir(root)
	if err != nil {
		return files, err
	}
	var absPath string
	for _, file := range fileInfo {
		absPath, err = GetAbsPath(root + "/" + file.Name())
		if err != nil {
			return nil, err
		}
		files = append(files, filepath.FromSlash(absPath))
	}
	return files, nil
}

// ProcessLine 可以用于处理大文件，按行读取
// filename: 文件名
// pf: 处理每一行的函数 int:行号，从1开始；string：该行的数据
// isContinue: pf函数报错后是否继续处理下一行
// jfile.JCONTINUE() 进行下个循环
// jfile.JBREAK()退出循环
//
// returns
//
// bool: 是否遍历完全
//
// int:  处理到哪一行，从1开始
//
// error: 报错
func ProcessLine(filename string, pf func(int64, string) error, isContinue bool) (bool, int64, error) {
	f, err := os.OpenFile(filename, os.O_RDONLY, 0)
	if err != nil {
		return false, -1, err
	}
	defer func() {
		f.Close()
	}()
	r := bufio.NewReader(f)
	var line_num int64 = 0
	for {
		line, err := readLine(r)
		if err != nil {
			if err == io.EOF {
				return true, line_num, nil
			}
			return false, line_num, err
		}
		// 使用传进来的函数处理line
		line_num += 1
		err = pf(line_num, line)
		if err != nil {
			if err.Error() == "JBREAK" {
				return false, line_num, nil
			} else if err.Error() == "JCONTINUE" {
				continue
			} else if !isContinue {
				return false, line_num, err
			}
		}

	}
}

// ProcessLineReverse 可以用于处理大文件，按行读取
// filename: 文件名
// pf: 处理每一行的函数 int:行号，从1开始，从最后一行开始；string：该行的数据，不包含该行末尾的\r\n或\n
// isContinue: pf函数报错后是否继续处理下一行
// jfile.JCONTINUE() 进行下个循环
// jfile.JBREAK()退出循环
//
// returns
//
// bool: 是否遍历完全
//
// int:  处理到哪一个偏移，从-1开始，从后向前
//
// error: 报错
func ProcessLineReverse(filename string, pf func(int64, string) error, isContinue bool) (bComplete bool, offset int64, err error) {
	f, err := os.OpenFile(filename, os.O_RDONLY, 0)
	if err != nil {
		return false, -1, err
	}
	defer func() {
		f.Close()
	}()
	buff := make([]byte, 0, 4096)
	char := make([]byte, 1)

	// 查询文件大小
	stat, err := f.Stat()
	if err != nil {
		return false, 0, err
	}
	filesize := stat.Size()
	if filesize < 1 {
		return false, 0, fmt.Errorf("file size(%d) < 1", filesize)
	}
	// 文件偏移
	offset = 0
	// 从后向前的行数
	var lineCount int64 = 0
	lastByteIsSlashN := false
	for {
		offset -= 1
		_, err = f.Seek(offset, io.SeekEnd)
		if err != nil {
			return
		}
		_, err = f.Read(char)
		if err != nil {
			return
		}
		//jlog.Debug("char:", char)
		if char[0] == '\n' {
			lastByteIsSlashN = true
			if len(buff) > 0 {
				revers(buff)
			}
			// 读取到的行
			lineCount++
			err = pf(lineCount, string(buff))
			if err != nil {
				if err.Error() == "JBREAK" {
					return false, offset, nil
				} else if err.Error() == "JCONTINUE" {
					continue
				} else if !isContinue {
					return false, offset, err
				}
			}
			buff = buff[:0]
		} else if char[0] == '\r' && lastByteIsSlashN {
			lastByteIsSlashN = false
			buff = buff[:0]
		} else {
			buff = append(buff, char[0])
		}
		// 判断是否是文件开头
		if offset == -filesize {
			lineCount++
			err = pf(lineCount, string(buff))
			if err != nil {
				if err.Error() == "JBREAK" {
					return false, offset, nil
				} else if err.Error() == "JCONTINUE" {
					break
				} else if !isContinue {
					return false, offset, err
				}
			}
			break
		}
	}
	return true, offset, nil
}

func revers(s []byte) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}

// 解决单行超过4096字节的文本读取问题
func readLine(r *bufio.Reader) (string, error) {
	line2 := []byte{}
	line, isprefix, err := r.ReadLine()
	line2 = append(line2, line...)
	for isprefix && err == nil {
		var bs []byte
		bs, isprefix, err = r.ReadLine()
		line2 = append(line2, bs...)
	}
	return string(line2), err
}

// 判断文件内包含某个字节数组的数量,没有重叠 如：kkkk中包含两个kk
func containsBytesCount(filepa string, cbytes []byte) int {
	f, err := os.Open(filepa)
	if err != nil {
		return 0
	}
	defer f.Close()
	// 每次读500字节
	buf := make([]byte, 50)
	cbytes2 := make([]byte, len(cbytes))
	var seek int64 = 0
	var count = 0
	for {
		rLens, err := f.Read(buf)
		if err != nil {
			break
		}
		//if bytes.Contains(buf,cbytes){
		//	return 1
		//}else{
		// 判断当前读取出来的是否含有第一个字节
		var k = 0
		for ; k < len(buf); k++ {
			//for k,v := range buf{
			if buf[k] == cbytes[0] {
				f.ReadAt(cbytes2, seek+int64(k))
				if bytes.Compare(cbytes, cbytes2) == 0 {
					//jlog.Debug(seek+int64(k))
					k = k + len(cbytes)
					count++
					//return true
				}
			}
		}
		//}
		if rLens < k {
			seek += int64(k)
		} else {
			seek += int64(rLens)
		}
		f.Seek(seek, io.SeekStart)
	}
	return count
}

// FileCopy 文件复制从src到dst
//
// b_trunc:覆盖目的文件
func FileCopy(src string, dst string, b_trunc bool) error {
	srcf, err := os.Open(src)
	if err != nil {
		return err
	}
	defer srcf.Close()
	var dstf = &os.File{}
	if b_trunc {
		dstf, err = os.OpenFile(dst, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0755)
		if err != nil {
			return err
		}
	} else {
		dstf, err = os.OpenFile(dst, os.O_RDWR|os.O_CREATE, 0755)
		if err != nil {
			return err
		}
	}

	defer dstf.Close()
	buf := make([]byte, 500)
	for {
		n, err := srcf.Read(buf)
		if err != nil && err != io.EOF {
			return err
		}
		if n == 0 {
			break
		}
		if _, err := dstf.Write(buf[:n]); err != nil {
			return err
		}
	}
	return nil
}

// FileMove 文件移动从src到dst
func FileMove(src string, dst string) error {
	err := FileCopy(src, dst, true)
	if err != nil {
		return err
	}
	err = os.Remove(src)
	if err != nil {
		return err
	}
	return nil
}

// GetLineCount 获取文件的行数
func GetLineCount(filePath string) int64 {
	//Open file
	var count int64
	_, _, err := ProcessLine(filePath, func(lineNum int64, line string) error {
		count = int64(lineNum)
		return JCONTINUE()
	}, false)
	if err != nil {
		return 0
	}
	return count
}

// GetLineData 获取特定第几行数据
// line_num:从1开始
func GetLineData(file_path string, line_num int64) (line_data string) {
	ProcessLine(file_path, func(inner_num int64, line string) error {
		if inner_num == line_num {
			line_data = line
			return JBREAK()
		}
		return nil
	}, false)
	return
}

// GetAllLines 获取所有的文件行数据
func GetAllLines(file_path string) (lines []string, total_line_count int64) {
	ProcessLine(file_path, func(inner_num int64, line string) error {
		lines = append(lines, line)
		total_line_count = inner_num
		return nil
	}, false)
	return
}

// GetHeadNLines 获取前n行数据
func GetHeadNLines(file_path string, n int64) (lines []string, total_line_count int64) {
	ProcessLine(file_path, func(inner_num int64, line string) error {
		lines = append(lines, line)
		total_line_count = inner_num
		if n == inner_num {
			return JBREAK()
		}
		return JCONTINUE()
	}, false)
	return
}

// GetTailNLines 获取后n行数据
func GetTailNLines(file_path string, n int64) (lines []string, start_line_count int64) {
	all_line_count := GetLineCount(file_path)
	start_line_count = all_line_count - int64(n) + 1
	ProcessLine(file_path, func(inner_num int64, line string) error {
		if int64(inner_num) < start_line_count {
			return JBREAK()
		}
		lines = append(lines, line)
		return JCONTINUE()
	}, false)
	return
}

// JBREAK 退出循环
func JBREAK() error {
	return fmt.Errorf("JBREAK")
}

// JCONTINUE 进行下次循环
func JCONTINUE() error {
	return fmt.Errorf("JCONTINUE")
}

// GetFileSize 返回文件的大小
func GetFileSize(filename string) (int64, error) {
	// 获取文件信息
	fileInfo, err := os.Stat(filename)
	if err != nil {
		return 0, err
	}

	// 获取文件大小
	size := fileInfo.Size()
	return size, nil
}
