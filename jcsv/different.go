package jcsv

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// 找到在第一个文件中，不在第二文件中的内容
func GetUnique2FileByID(file1, file2, delimeter string, file1_column_id, file2_column_id int, file1_with_header, file2_with_header bool, outputFileName string) (unique_count int, err error) {
	// 校验
	file1_column_id = GetColumnID(file1, delimeter, "", file1_column_id, file1_with_header)
	if file1_column_id == -1 {
		return 0, fmt.Errorf("%scolumn_id有误", file1)
	}
	file2_column_id = GetColumnID(file2, delimeter, "", file2_column_id, file2_with_header)
	if file2_column_id == -1 {
		return 0, fmt.Errorf("%scolumn_id有误", file2)
	}
	// 打开第一个文件
	f1, err := os.Open(file1)
	if err != nil {
		return 0, err
	}
	defer f1.Close()

	// 打开第二个文件
	f2, err := os.Open(file2)
	if err != nil {
		return 0, err
	}
	defer f2.Close()

	// 读取第二个文件的内容到一个map中
	contentMap := make(map[string]bool)
	scanner2 := bufio.NewScanner(f2)
	f2_line_num := 0
	for scanner2.Scan() {
		f2_line_num += 1
		if file2_with_header && f2_line_num == 1 {
			continue
		}
		line := scanner2.Text()
		word_list := strings.Split(line, delimeter)
		column_value := word_list[file2_column_id]
		contentMap[column_value] = true
	}

	// 读取第一个文件的内容，并记录不在第二个文件中的内容
	var uniqueCount int = 0
	// 创建输出文件
	outputFile, err := os.Create(outputFileName)
	if err != nil {
		return 0, err
	}
	defer outputFile.Close()
	scanner1 := bufio.NewScanner(f1)
	f1_line_num := 0
	for scanner1.Scan() {
		f1_line_num += 1
		if file1_with_header && f1_line_num == 1 {
			continue
		}
		line := scanner1.Text()
		word_list := strings.Split(line, delimeter)
		column_value := word_list[file1_column_id]
		if _, exists := contentMap[column_value]; !exists {
			uniqueCount += 1
			outputFile.Write([]byte(fmt.Sprintf("%s\n", column_value)))
		}
	}

	if err := scanner1.Err(); err != nil {
		return 0, err
	}

	return uniqueCount, nil
}

func GetUnique2FileByName(file1, file2, delimeter string, file1_column_name, file2_column_name string, file1_with_header, file2_with_header bool, outputFileName string) (unique_count int, err error) {
	file1_column_id := GetColumnID(file1, delimeter, file1_column_name, -1, file1_with_header)
	file2_column_id := GetColumnID(file2, delimeter, file2_column_name, -1, file2_with_header)
	// 校验
	if file1_column_id == -1 {
		return 0, fmt.Errorf("%scolumn_id有误", file1)
	}
	if file2_column_id == -1 {
		return 0, fmt.Errorf("%scolumn_id有误", file2)
	}
	// 打开第一个文件
	f1, err := os.Open(file1)
	if err != nil {
		return 0, err
	}
	defer f1.Close()

	// 打开第二个文件
	f2, err := os.Open(file2)
	if err != nil {
		return 0, err
	}
	defer f2.Close()

	// 读取第二个文件的内容到一个map中
	contentMap := make(map[string]bool)
	scanner2 := bufio.NewScanner(f2)
	f2_line_num := 0
	for scanner2.Scan() {
		f2_line_num += 1
		if file2_with_header && f2_line_num == 1 {
			continue
		}
		line := scanner2.Text()
		word_list := strings.Split(line, delimeter)
		column_value := word_list[file2_column_id]
		contentMap[column_value] = true
	}

	// 读取第一个文件的内容，并记录不在第二个文件中的内容
	var uniqueCount int = 0
	// 创建输出文件
	outputFile, err := os.Create(outputFileName)
	if err != nil {
		return 0, err
	}
	defer outputFile.Close()
	scanner1 := bufio.NewScanner(f1)
	f1_line_num := 0
	for scanner1.Scan() {
		f1_line_num += 1
		if file1_with_header && f1_line_num == 1 {
			continue
		}
		line := scanner1.Text()
		word_list := strings.Split(line, delimeter)
		column_value := word_list[file1_column_id]
		if _, exists := contentMap[column_value]; !exists {
			uniqueCount += 1
			outputFile.Write([]byte(fmt.Sprintf("%s\n", column_value)))
		}
	}

	if err := scanner1.Err(); err != nil {
		return 0, err
	}

	return uniqueCount, nil
}

func GetUniqueWordsByName(file1, file2, delimeter string, file1_column_name, file2_column_name string, file1_with_header, file2_with_header bool) (unique_word_list []string, err error) {
	file1_column_id := GetColumnID(file1, delimeter, file1_column_name, -1, file1_with_header)
	file2_column_id := GetColumnID(file2, delimeter, file2_column_name, -1, file2_with_header)
	// 校验
	if file1_column_id == -1 {
		return nil, fmt.Errorf("%scolumn_id有误", file1)
	}
	if file2_column_id == -1 {
		return nil, fmt.Errorf("%scolumn_id有误", file2)
	}
	// 打开第一个文件
	f1, err := os.Open(file1)
	if err != nil {
		return nil, err
	}
	defer f1.Close()

	// 打开第二个文件
	f2, err := os.Open(file2)
	if err != nil {
		return nil, err
	}
	defer f2.Close()

	// 读取第二个文件的内容到一个map中
	contentMap := make(map[string]bool)
	scanner2 := bufio.NewScanner(f2)
	f2_line_num := 0
	for scanner2.Scan() {
		f2_line_num += 1
		if file2_with_header && f2_line_num == 1 {
			continue
		}
		line := scanner2.Text()
		word_list := strings.Split(line, delimeter)
		column_value := word_list[file2_column_id]
		contentMap[column_value] = true
	}

	// 读取第一个文件的内容，并记录不在第二个文件中的内容
	var uniqueCount int = 0
	// 创建输出文件
	//outputFile, err := os.Create(outputFileName)
	//if err != nil {
	//	return 0, err
	//}
	//defer outputFile.Close()
	scanner1 := bufio.NewScanner(f1)
	f1_line_num := 0
	for scanner1.Scan() {
		f1_line_num += 1
		if file1_with_header && f1_line_num == 1 {
			continue
		}
		line := scanner1.Text()
		word_list := strings.Split(line, delimeter)
		column_value := word_list[file1_column_id]
		if _, exists := contentMap[column_value]; !exists {
			uniqueCount += 1
			unique_word_list = append(unique_word_list, column_value)
			//outputFile.Write([]byte(fmt.Sprintf("%s\n", column_value)))
		}
	}

	if err := scanner1.Err(); err != nil {
		return nil, err
	}

	return
}

// 读取在file1文件中但不在file2文件中的指定单元格所对应的行，且指定的单元格值符合给定的后缀列表
func GetQualifiedUnique2FileByName(suffix_list []string, file1, file2, delimeter string, file1_column_name, file2_column_name string, file1_with_header, file2_with_header bool, outputFileName string) (unique_count int, err error) {
	file1_column_id := GetColumnID(file1, delimeter, file1_column_name, -1, file1_with_header)
	file2_column_id := GetColumnID(file2, delimeter, file2_column_name, -1, file2_with_header)
	// 校验
	if file1_column_id == -1 {
		return 0, fmt.Errorf("%scolumn_id有误", file1)
	}
	if file2_column_id == -1 {
		return 0, fmt.Errorf("%scolumn_id有误", file2)
	}
	// 打开第一个文件
	f1, err := os.Open(file1)
	if err != nil {
		return 0, err
	}
	defer f1.Close()

	// 打开第二个文件
	f2, err := os.Open(file2)
	if err != nil {
		return 0, err
	}
	defer f2.Close()

	// 读取第二个文件的内容到一个map中
	contentMap := make(map[string]bool)
	scanner2 := bufio.NewScanner(f2)
	f2_line_num := 0
	for scanner2.Scan() {
		f2_line_num += 1
		if file1_with_header && f2_line_num == 1 {
			continue
		}
		line := scanner2.Text()
		word_list := strings.Split(line, delimeter)
		column_value := word_list[file2_column_id]
		contentMap[column_value] = true
	}

	// 读取第一个文件的内容，并记录不在第二个文件中的内容
	var uniqueCount int = 0
	// 创建输出文件
	outputFile, err := os.Create(outputFileName)
	if err != nil {
		return 0, err
	}
	defer outputFile.Close()
	scanner1 := bufio.NewScanner(f1)
	f1_line_num := 0
	for scanner1.Scan() {
		f1_line_num += 1
		if file1_with_header && f1_line_num == 1 {
			continue
		}

		line := scanner1.Text()
		word_list := strings.Split(line, delimeter)
		column_value := word_list[file1_column_id]
		// 过滤掉不符合后缀的
		for _, suffix := range suffix_list {
			if !strings.HasSuffix(column_value, suffix) {
				continue
			}
		}
		if _, exists := contentMap[column_value]; !exists {
			uniqueCount += 1
			outputFile.Write([]byte(fmt.Sprintf("%s\n", column_value)))
		}
	}

	if err := scanner1.Err(); err != nil {
		return 0, err
	}

	return uniqueCount, nil
}

// 从指定的csv文件，读取符合后缀的行到outpuFileName文件中，根据指定的file1_column_name或file1_column_id来判断
//
// 返回的line_count包含表头
func GetQualifiedUniqueLine2FileByName(suffix_list []string, file1, delimeter, file1_column_name string, file1_column_id int, file1_with_header bool, outputFileName string) (line_count int, err error) {
	file1_column_id = GetColumnID(file1, delimeter, file1_column_name, file1_column_id, file1_with_header)
	// 校验
	if file1_column_id == -1 {
		return 0, fmt.Errorf("%scolumn_id有误", file1)
	}
	// 打开第一个文件
	f1, err := os.Open(file1)
	if err != nil {
		return 0, err
	}
	defer f1.Close()

	// 读取第一个文件的内容
	line_count = 0
	// 创建输出文件
	outputFile, err := os.Create(outputFileName)
	if err != nil {
		return 0, err
	}
	defer outputFile.Close()
	scanner1 := bufio.NewScanner(f1)
	f1_line_num := 0
	for scanner1.Scan() {
		line := scanner1.Text()
		f1_line_num += 1
		if file1_with_header && f1_line_num == 1 {
			line_count += 1
			outputFile.Write([]byte(fmt.Sprintf("%s\n", line)))
			continue
		}

		word_list := strings.Split(line, delimeter)
		column_value := word_list[file1_column_id]
		// 只记录符合后缀的
		for _, suffix := range suffix_list {
			if strings.HasSuffix(column_value, suffix) {
				line_count += 1
				outputFile.Write([]byte(fmt.Sprintf("%s\n", column_value)))
				break
			}
		}
	}

	if err := scanner1.Err(); err != nil {
		return 0, err
	}

	return line_count, nil
}
