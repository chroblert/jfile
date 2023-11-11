package jcsv

import (
	"fmt"
	"github.com/chroblert/jarr"
	"github.com/chroblert/jfile"
	"strings"
)

// 判断指定列后的值是否包含指定的后缀
func ContainsSuffixByID(src_file string, delimeter string, column_id int, suffix_list []string) bool {
	first_line := jfile.GetLineData(src_file, 1)
	word_list := strings.Split(first_line, delimeter)
	// 判断个数. 列数不能少于id+1+后缀个数
	if len(word_list) < column_id+1+len(suffix_list) {
		return false
	}
	column_name := word_list[column_id]
	for k, suffix := range suffix_list {
		if word_list[column_id+k+1] != fmt.Sprintf("%s%s", column_name, suffix) {
			return false
		}
	}
	return true
}

func ContainsSuffixByName(src_file string, delimeter string, column_name string, suffix_list []string) bool {
	first_line := jfile.GetLineData(src_file, 1)
	word_list := strings.Split(first_line, delimeter)
	// 没有column_name
	if !jarr.EleInArr(column_name, word_list) {
		return false
	}
	for k, v := range word_list {
		if v == column_name {
			// 表头剩余的列表个数，少于，指定的后缀个数，false
			if len(word_list)-(k+1) < len(suffix_list) {
				return false
			} else {
				for k2, suffix := range suffix_list {
					if word_list[k+k2+1] != suffix {
						return false
					}
				}
				return true
			}
		}
	}
	return false
}
