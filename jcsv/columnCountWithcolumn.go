package jcsv

import (
	"fmt"
	"github.com/chroblert/jarr"
	"github.com/chroblert/jfile"
	"github.com/chroblert/jlog"
	"strings"
)

// 在csv文件的指定列右侧添加两列，_count数量，_order排序
// column_id:从0开始，-1表示使用column_name
// ref_是参考的。比如：一个city下有多少个ip，city是ref_,ip是column
//
// 插入的列名：refname_columnname<suffix>
func StatisticUniqueColumnCountWithKey(src_file string, column_name string, ref_name string, delimiter string, column_id int, ref_id int, with_header bool, suffix string) (new_column_name string, err error) {
	if suffix == "" {
		suffix = "_count"
	}
	// 获取列名ID
	column_id = GetColumnID(src_file, delimiter, column_name, column_id, with_header)
	if column_id == -1 {
		return "", fmt.Errorf("指定的column不存在，或columnd_id不对")
	}
	column_name = GetColumnName(src_file, delimiter, column_id, true)
	ref_id = GetColumnID(src_file, delimiter, ref_name, ref_id, with_header)
	if ref_id == -1 {
		return "", fmt.Errorf("指定的column不存在，或ref_id不对")
	}
	ref_name = GetColumnName(src_file, delimiter, ref_id, true)
	// 判断文件是否已经存在指定列
	b_exist := ContainsSuffixByID(src_file, delimiter, ref_id, []string{
		fmt.Sprintf("_%s%s", column_name, suffix),
	})
	// 存在，则将指定列删除
	if b_exist {
		jlog.Info("[+] 已经存在列，将会删除指定列:", fmt.Sprintf("%s_%s%s", ref_name, column_name, suffix))
		DeleteColumn(src_file, delimiter, ref_id+1)
	}
	//
	var ref_column_map = make(map[string][]string)
	var ref_column_count_map = make(map[string]int)
	jfile.ProcessLine64(src_file, func(line_num int64, line string) error {
		if line_num == 1 {
			jfile.JCONTINUE()
		}
		word_list := strings.Split(line, delimiter)
		column_value := word_list[column_id]
		ref_value := word_list[ref_id]
		if _, ok := ref_column_map[ref_value]; !ok {
			ref_column_map[ref_value] = []string{column_value}
		} else {
			ref_column_map[ref_value] = append(ref_column_map[ref_value], column_value)
			ref_column_map[ref_value] = jarr.UniqueString(ref_column_map[ref_value])
		}
		return nil
	}, false)
	for k, v := range ref_column_map {
		ref_column_count_map[k] = len(v)
	}
	// 写回原文件
	new_column_name = fmt.Sprintf("%s_%s%s", ref_name, column_name, suffix)
	err = AddColumnByID(src_file, delimiter, ref_id, ref_column_count_map, new_column_name, true)
	if err != nil {
		return new_column_name, err
	}
	jlog.Infof("[*] 将添加了_%s_count的数据写入文件\n", column_name)
	return new_column_name, nil
}
