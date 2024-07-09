package jcsv

import (
	"fmt"
	"github.com/chroblert/jfile"
	"github.com/chroblert/jlog"
	"os"
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
	jfile.ProcessLine64(src_file, func(line_num int64, line string) error {
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

func GetColumnStringValuesByID(src_file string, delimeter string, column_id int, with_header bool) (column_value_list []string, err error) {
	// 校验
	column_id = GetColumnID(src_file, delimeter, "", column_id, with_header)
	if column_id == -1 {
		jlog.Error("列数不足")
		return nil, fmt.Errorf("列数不足")
	}
	//
	jfile.ProcessLine64(src_file, func(line_num int64, line string) error {
		if with_header && line_num == 1 {
			return jfile.JCONTINUE()
		}
		word_list := strings.Split(line, delimeter)
		column_value := word_list[column_id]
		column_value_list = append(column_value_list, column_value)
		return jfile.JCONTINUE()
	}, false)
	return column_value_list, nil
}

func GetColumnStringValuesByName(src_file string, delimeter string, column_name string, with_header bool) (column_value_list []string, err error) {
	// 校验
	column_id := GetColumnID(src_file, delimeter, column_name, -1, with_header)
	if column_id == -1 {
		jlog.Error("列数不足")
		return nil, fmt.Errorf("列数不足")
	}
	//
	jfile.ProcessLine64(src_file, func(line_num int64, line string) error {
		if with_header && line_num == 1 {
			return jfile.JCONTINUE()
		}
		word_list := strings.Split(line, delimeter)
		column_value := word_list[column_id]
		column_value_list = append(column_value_list, column_value)
		return jfile.JCONTINUE()
	}, false)
	return column_value_list, nil
}

func GetColumnStringValues2FileByName(src_file string, delimeter string, column_name string, with_header bool, outputFileName string) (column_count int, err error) {
	// 校验
	column_id := GetColumnID(src_file, delimeter, column_name, -1, with_header)
	if column_id == -1 {
		jlog.Error("列数不足")
		return 0, fmt.Errorf("列数不足")
	}
	// 创建输出文件
	outputFile, err := os.Create(outputFileName)
	if err != nil {
		return 0, err
	}
	defer outputFile.Close()
	column_count = 0
	jfile.ProcessLine64(src_file, func(line_num int64, line string) error {
		if with_header && line_num == 1 {
			return jfile.JCONTINUE()
		}
		word_list := strings.Split(line, delimeter)
		column_value := word_list[column_id]
		column_count += 1
		outputFile.Write([]byte(fmt.Sprintf("%s\n", column_value)))
		return jfile.JCONTINUE()
	}, false)
	return column_count, nil
}

func GetColumnStringValues2FileByID(src_file string, delimeter string, column_id int, with_header bool, outputFileName string) (column_count int, err error) {
	// 校验
	column_id = GetColumnID(src_file, delimeter, "", column_id, with_header)
	if column_id == -1 {
		jlog.Error("列数不足")
		return 0, fmt.Errorf("列数不足")
	}
	// 创建输出文件
	outputFile, err := os.Create(outputFileName)
	if err != nil {
		return 0, err
	}
	defer outputFile.Close()
	column_count = 0
	jfile.ProcessLine64(src_file, func(line_num int64, line string) error {
		if with_header && line_num == 1 {
			return jfile.JCONTINUE()
		}
		word_list := strings.Split(line, delimeter)
		column_value := word_list[column_id]
		column_count += 1
		outputFile.Write([]byte(fmt.Sprintf("%s\n", column_value)))
		return jfile.JCONTINUE()
	}, false)
	return column_count, nil
}
