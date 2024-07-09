package jfile

import (
	"crypto/md5"
	"fmt"
	"github.com/chroblert/jstr"
	"os"
)

func Unique(srcFile, uniqueFile string) (uniqueCount int64, err error) {
	outputFile := uniqueFile
	seenLines := make(map[[16]byte]bool)
	uniqueCount = 0
	output, err := os.OpenFile(outputFile, os.O_WRONLY|os.O_TRUNC|os.O_CREATE, 0666)
	if err != nil {
		fmt.Println("无法创建输出文件:", err)
		return
	}
	defer output.Close()
	_, _, err = ProcessLine64(srcFile, func(line_num int64, line string) error {
		hash := calculateHash([]byte(line))
		if !seenLines[hash] {
			seenLines[hash] = true
			uniqueCount += 1
			_, err = output.Write([]byte(line + "\n"))
			if err != nil {
				fmt.Println("无法写入输出文件:", err)
				return JCONTINUE()
			}
		}
		return JCONTINUE()
	}, false)

	if err != nil {
		fmt.Println("读取输入文件时发生错误:", err)
		return
	}

	fmt.Println("去重完成，结果已写入", outputFile)
	return
}

func UniqueInSameFile(srcFile string) (uniqueCount int, err error) {
	// 移动文件
	bak_file := fmt.Sprintf("%s-%s", srcFile, jstr.GenerateRandomString(8))
	FileMove(srcFile, bak_file)

	//inputFile := srcFile
	outputFile := srcFile
	seenLines := make(map[[16]byte]bool)
	uniqueCount = 0
	output, err := os.Create(outputFile)
	if err != nil {
		fmt.Println("无法创建输出文件:", err)
		return
	}
	defer output.Close()
	_, _, err = ProcessLine64(bak_file, func(line_Num int64, line string) error {
		hash := calculateHash([]byte(line))
		if !seenLines[hash] {
			seenLines[hash] = true
			uniqueCount += 1
			_, err = output.Write([]byte(line + "\n"))
			//_, err = io.WriteString(output, line+"\n")
			if err != nil {
				fmt.Println("无法写入输出文件:", err)
				return JCONTINUE()
			}
		}
		return JCONTINUE()
	}, false)
	os.Remove(bak_file)
	if err != nil {
		fmt.Println("读取输入文件时发生错误:", err)
		return
	}

	fmt.Println("去重完成，结果已写入", outputFile)
	return
}

func calculateHash(data []byte) [16]byte {
	return md5.Sum(data)
}
