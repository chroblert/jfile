package jcsv

import (
	"github.com/chroblert/jfile"
	"github.com/chroblert/jlog"
	"strconv"
	"strings"
)

func GetColumnIntValues(src_file string, delimeter string, column_id int, with_header bool) (column_value_list []int, err error) {
	// 校验
	column_id = GetColumnID(src_file, delimeter, "", column_id, with_header)
	if column_id == -1 {
		jlog.Error("列数不足")
		return
	}
	//
	jfile.ProcessLine(src_file, func(line_num int, line string) error {
		if with_header && line_num == 1 {
			return jfile.JCONTINUE()
		}
		word_list := strings.Split(line, delimeter)
		column_value := word_list[column_id]
		int_value, err := strconv.Atoi(column_value)
		if err != nil {
			jlog.Error(err)
			int_value = 0
		}
		column_value_list = append(column_value_list, int_value)
		return nil
	}, false)
	return column_value_list, nil
}
