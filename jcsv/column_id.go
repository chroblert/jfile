package jcsv

import (
	"fmt"
	"github.com/chroblert/jfile"
	"strings"
)

// with_header: bool 首行是否是表头
//
// 有header 列名优先
// 若指定了id和name，则列数不能少于id+1,返回name所在的id
// 若指定了name，则返回第一个name所在的id
// 若指定了id，则判断列数不少于id+1，否则返回-1
// 没有header，只判断id
func GetColumnID(src_file, delimiter string, column_name string, column_id int, with_header bool) (real_column_id int) {
	line_num := 0
	real_column_id = -1
	jfile.ProcessLine(src_file, func(tmp int64, line string) error {
		line_num += 1
		word_list := strings.Split(line, delimiter)
		if line_num == 1 {
			if with_header {
				// 指定了列名
				if column_name != "" {
					for k, v := range word_list {
						if v == column_name {
							real_column_id = k
							return fmt.Errorf("JBREAK")
						}
					}
				} else if column_id != -1 {
					// 列数少于column_id + 1
					if len(word_list) < column_id+1 {
						real_column_id = -1
						return fmt.Errorf("format err")
					}
					real_column_id = column_id
					return fmt.Errorf("JBREAK")
				} else {
					real_column_id = -1
					return fmt.Errorf("JBREAK")
				}
			} else {
				if column_id != -1 {
					// 列数少于column_id + 1
					if len(word_list) < column_id+1 {
						real_column_id = -1
						return fmt.Errorf("format err")
					}
					real_column_id = column_id
					return fmt.Errorf("JBREAK")
				} else {
					real_column_id = -1
					return fmt.Errorf("JBREAK")
				}
			}

		}
		return nil
	}, false)
	return
}

func GetColumnName(src_file, delimiter string, column_id int, with_header bool) (column_name string) {
	line := jfile.GetLineData(src_file, 1)
	word_list := strings.Split(line, delimiter)
	if len(word_list) < (column_id + 1) {
		return ""
	}
	return word_list[column_id]
}
