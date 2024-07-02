package jcsv

import (
	"fmt"
	"github.com/chroblert/jfile"
	"github.com/chroblert/jlog"
	"github.com/chroblert/jstr"
	"os"
	"strings"
)

// 在指定的column_id后添加一列
func AddColumnByID(src_file string, delimeter string, column_id int, ref_column_map map[string]int, new_column_name string, with_header bool) (err error) {
	// 判断是否有列
	column_id = GetColumnID(src_file, delimeter, "", column_id, with_header)
	if column_id == -1 {
		return fmt.Errorf("no such column_id:%d", column_id)
	}
	// 备份文件
	bak_file := fmt.Sprintf("%s-%s", src_file, jstr.GenerateRandomString(8))
	jfile.FileCopy(src_file, bak_file, true)
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
	jfile.ProcessLine(bak_file, func(line_num int64, line string) error {
		word_list := strings.Split(line, ",")
		new_word_list := make([]string, len(word_list[:column_id+1]))
		copy(new_word_list, word_list[:column_id+1])
		// 加入新的word
		column_value := word_list[column_id]
		insert_word := ""
		if line_num == 1 {
			if with_header {
				insert_word = new_column_name
			} else {
				insert_word = fmt.Sprintf("%d", ref_column_map[column_value])
			}
		} else {
			insert_word = fmt.Sprintf("%d", ref_column_map[column_value])
		}
		new_word_list = append(new_word_list, insert_word)
		// 加入原word_list的剩余部分
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
