package jfile

import (
	"bufio"
	"crypto/md5"
	"fmt"
	"github.com/chroblert/jstr"
	"io"
	"os"
)

func Unique(srcFile, uniqueFile string) {
	//inputFile := srcFile
	outputFile := uniqueFile
	input, err := os.Open(srcFile)
	if err != nil {
		fmt.Println("无法打开输入文件:", err)
		return
	}
	defer input.Close()

	output, err := os.Create(outputFile)
	if err != nil {
		fmt.Println("无法创建输出文件:", err)
		return
	}
	defer output.Close()

	seenLines := make(map[[16]byte]bool)
	scanner := bufio.NewScanner(input)

	for scanner.Scan() {
		line := scanner.Text()
		lineBytes := []byte(line)
		hash := calculateHash(lineBytes)

		if !seenLines[hash] {
			seenLines[hash] = true
			_, err := io.WriteString(output, line+"\n")
			if err != nil {
				fmt.Println("无法写入输出文件:", err)
				return
			}
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("读取输入文件时发生错误:", err)
	}

	fmt.Println("去重完成，结果已写入", outputFile)
}

func UniqueInSameFile(srcFile string) (unique_count int, err error) {
	// 移动文件
	bak_file := fmt.Sprintf("%s-%s", srcFile, jstr.GenerateRandomString(8))
	FileMove(srcFile, bak_file)

	//inputFile := srcFile
	outputFile := srcFile
	input, err := os.Open(bak_file)
	if err != nil {
		fmt.Println("无法打开输入文件:", err)
		return
	}
	defer func() {
		input.Close()
		os.Remove(bak_file)
	}()

	output, err := os.Create(outputFile)
	if err != nil {
		fmt.Println("无法创建输出文件:", err)
		return
	}
	defer output.Close()

	seenLines := make(map[[16]byte]bool)
	scanner := bufio.NewScanner(input)
	unique_count = 0
	for scanner.Scan() {
		line := scanner.Text()
		lineBytes := []byte(line)
		hash := calculateHash(lineBytes)

		if !seenLines[hash] {
			seenLines[hash] = true
			unique_count += 1
			_, err = io.WriteString(output, line+"\n")
			if err != nil {
				fmt.Println("无法写入输出文件:", err)
				return
			}
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("读取输入文件时发生错误:", err)
	}

	fmt.Println("去重完成，结果已写入", outputFile)
	return
}

func calculateHash(data []byte) [16]byte {
	return md5.Sum(data)
}
