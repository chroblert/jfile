package jcsv

import (
	"fmt"
	"github.com/chroblert/jarr"
	"github.com/chroblert/jlog"
	"sort"
)

func AddOrderColumnByID(src_file string, delimeter string, column_count_id int, new_column_name string, with_header bool) (err error) {
	// 校验
	column_count_id = GetColumnID(src_file, delimeter, "", column_count_id, with_header)
	if column_count_id == -1 {
		jlog.Error("列数不足")
		return
	}
	// 获取所有的数值
	column_values, err := GetColumnIntValues(src_file, delimeter, column_count_id, with_header)
	if err != nil {
		return
	}
	column_values = jarr.UniqueInt(column_values)
	// 降序排列
	sort.Slice(column_values, func(i, j int) bool {
		return column_values[i] > column_values[j]
	})
	var column_value_order_map = make(map[string]int)
	for order, v := range column_values {
		if _, ok := column_value_order_map[fmt.Sprintf("%d", v)]; !ok {
			column_value_order_map[fmt.Sprintf("%d", v)] = order + 1
		}
	}
	// 将顺序写入文件
	err = AddColumnByID(src_file, delimeter, column_count_id, column_value_order_map, new_column_name, with_header)
	return
}

func AddOrderColumnByName(src_file string, delimeter string, column_count_name string, new_column_name string, with_header bool) (err error) {
	column_count_id := GetColumnID(src_file, delimeter, column_count_name, -1, with_header)
	// 校验
	if column_count_id == -1 {
		jlog.Error("列数不足")
		return
	}
	// 获取所有的数值
	column_values, err := GetColumnIntValues(src_file, delimeter, column_count_id, with_header)
	if err != nil {
		return
	}
	column_values = jarr.UniqueInt(column_values)
	// 降序排列
	sort.Slice(column_values, func(i, j int) bool {
		return column_values[i] > column_values[j]
	})
	var column_value_order_map = make(map[string]int)
	for order, v := range column_values {
		if _, ok := column_value_order_map[fmt.Sprintf("%d", v)]; !ok {
			column_value_order_map[fmt.Sprintf("%d", v)] = order + 1
		}
	}
	// 将顺序写入文件
	err = AddColumnByID(src_file, delimeter, column_count_id, column_value_order_map, new_column_name, with_header)
	return
}
