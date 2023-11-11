package jcsv

import (
	"github.com/chroblert/jstr"

	"fmt"
	"github.com/chroblert/jfile"
	"github.com/chroblert/jlog"
	"os"
	"strings"
)

// 删除csv文件的指定列
// column_id:从0开始
func DeleteColumn(src_file string, delimeter string, column_id int) (err error) {
	// 判断是否有列
	column_id = GetColumnID(src_file, delimeter, "", column_id, false)
	if column_id == -1 {
		return fmt.Errorf("no such column_id:%d", column_id)
	}
	// 备份文件
	bak_file := fmt.Sprintf("%s-%s", src_file, jstr.GenerateRandomString(8))
	jfile.FileCopy(src_file, bak_file)
	// 删除源文件
	err = os.Remove(src_file)
	if err != nil {
		return err
	}
	// 创建文件
	f, err := os.OpenFile(src_file, os.O_CREATE|os.O_TRUNC|os.O_TRUNC, 0666)
	if err != nil {
		return err
	}
	defer func() {
		if f != nil {
			f.Close()
		}
	}()
	//
	jfile.ProcessLine(bak_file, func(line_num int, line string) error {
		word_list := strings.Split(line, ",")
		new_word_list := make([]string, len(word_list[:column_id]))
		copy(new_word_list, word_list[:column_id])
		if len(word_list)-(column_id+1) > 0 {
			for _, v := range word_list[column_id+1:] {
				new_word_list = append(new_word_list, v)
			}
		}
		new_line := fmt.Sprintf("%s\n", strings.Join(new_word_list, delimeter))
		f.Write([]byte(new_line))
		return nil
	}, false)
	if f != nil {
		f.Close()
	}
	// 删除bak_file
	err = os.Remove(bak_file)
	if err != nil {
		jlog.Error(err)
	}
	return nil
}
