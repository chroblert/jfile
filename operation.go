package jfile

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

// 合并多个文件
func MergeFiles(outputFileName string, inputFiles ...string) error {
	// 创建输出文件
	outputFile, err := os.Create(outputFileName)
	if err != nil {
		return err
	}
	defer outputFile.Close()

	// 循环遍历输入文件
	for _, inputFile := range inputFiles {
		// 打开输入文件
		input, err := os.Open(inputFile)
		if err != nil {
			return err
		}
		defer input.Close()

		// 将输入文件的内容复制到输出文件中
		_, err = io.Copy(outputFile, input)
		if err != nil {
			return err
		}
	}

	fmt.Println("文件合并成功:", outputFileName)
	return nil
}

// 找到file1文件与file2文件不重复的内容
func FindUniqueContent(file1, file2 string) ([]string, error) {
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
	for scanner2.Scan() {
		line := scanner2.Text()
		contentMap[line] = true
	}

	// 读取第一个文件的内容，并记录不在第二个文件中的内容
	var uniqueContent []string
	scanner1 := bufio.NewScanner(f1)
	for scanner1.Scan() {
		line := scanner1.Text()
		if _, exists := contentMap[line]; !exists {
			uniqueContent = append(uniqueContent, line)
		}
	}

	if err := scanner1.Err(); err != nil {
		return nil, err
	}

	return uniqueContent, nil
}

// 找到file1文件与file2文件不重复的内容
func FindUniqueContent2File(file1, file2, outputFileName string) (int, error) {
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
	for scanner2.Scan() {
		line := scanner2.Text()
		contentMap[line] = true
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
	for scanner1.Scan() {
		line := scanner1.Text()
		if _, exists := contentMap[line]; !exists {
			uniqueCount += 1
			outputFile.Write([]byte(fmt.Sprintf("%s\n", line)))
		}
	}

	if err := scanner1.Err(); err != nil {
		return 0, err
	}

	return uniqueCount, nil
}
