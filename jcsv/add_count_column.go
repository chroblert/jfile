package jcsv

import (
	"fmt"
	"github.com/chroblert/jfile"
	"github.com/chroblert/jstr"
	"os"
	"strings"
)

func AddCountColumn(src_file string, delimeter string, column_id int, new_column_name string, with_header bool) (column_count_map map[string]int, err error) {
	// 判断是否有列
	column_id = GetColumnID(src_file, delimeter, "", column_id, with_header)
	if column_id == -1 {
		return nil, fmt.Errorf("no such column_id:%d", column_id)
	}
	// 备份文件
	bak_file := fmt.Sprintf("%s-%s", src_file, jstr.GenerateRandomString(8))
	jfile.FileCopy(src_file, bak_file, true)
	// 删除源文件
	err = os.Remove(src_file)
	if err != nil {
		return nil, err
	}
	// 创建文件
	f, err := os.OpenFile(src_file, os.O_CREATE|os.O_TRUNC|os.O_TRUNC, 0666)
	if err != nil {
		return nil, err
	}
	defer func() {
		if f != nil {
			f.Close()
		}
	}()
	//
	column_count_map = GetCountOfColumn(src_file, delimeter, column_id, with_header)
	if column_count_map == nil || len(column_count_map) == 0 {
		return nil, fmt.Errorf("统计单元格失败")
	}
	//column_vlaue := GetColumnName(src_file, delimeter, column_id, with_header)
	// 添加一列
	err = AddColumnByID(src_file, delimeter, column_id, column_count_map, new_column_name, with_header)

	return
}

func GetCountOfColumn(src_file, delimeter string, column_id int, with_header bool) (column_count_map map[string]int) {
	column_id = GetColumnID(src_file, delimeter, "", column_id, with_header)
	if column_id == -1 {
		return
	}
	column_count_map = make(map[string]int)
	// 统计单元格数值
	jfile.ProcessLine(src_file, func(line_num int, line string) error {
		if with_header && line_num == 1 {
			return jfile.JCONTINUE()
		}
		word_list := strings.Split(line, delimeter)
		column_value := word_list[column_id]
		if _, ok := column_count_map[column_value]; !ok {
			column_count_map[column_value] = 1
		} else {
			column_count_map[column_value] += 1
		}
		return jfile.JCONTINUE()
	}, false)
	return
}
