package jexcel

import (
	"fmt"
	"github.com/tealeg/xlsx/v2"
)

// ProcessLine64 读取excel文件，按行输出
// filename: 文件名
// pf: 处理每一行的函数 int64:行号，从1开始；[]*xlsx.Cell：该行的cell
// isContinue: pf函数报错后是否继续处理下一行
// jfile.JCONTINUE() 进行下个循环
// jfile.JBREAK()退出循环
// returns
// bool: 是否遍历完全
// int64:  处理到哪一行，从1开始
// error: 报错
func ProcessLine64(filename, sheetName string, pf func(lineNum int64, cells []*xlsx.Cell) (err error), isContinue bool) (bDone bool, doneLineNum int64, err error) {
	var line_num int64 = 0
	// 打开 Excel 文件
	f, err := xlsx.OpenFile(filename)
	if err != nil {
		return false, -1, err
	}
	//
	// 获取工作表列表
	sheets := f.Sheets
	if len(sheets) == 0 {
		err = fmt.Errorf("Excel 文件中没有工作表")
		return false, -1, err
	}
	if sheetName == "" {
		sheetName = sheets[0].Name // 读取第一个工作表
	}
	bMatched := false
	var matchedSheet *xlsx.Sheet
	for _, sheet := range sheets {
		if sheet.Name == sheetName {
			bMatched = true
			matchedSheet = sheet
			break
		}
	}
	if !bMatched {
		err = fmt.Errorf("no such sheet")
		return false, -1, err
	}
	// 遍历行数据（跳过表头）
	for _, row := range matchedSheet.Rows {
		// 使用传进来的函数处理line
		line_num += 1
		// 业务逻辑
		//var colValueList = make([]string, len(row.Cells))
		//for i, cell := range row.Cells {
		//	colValueList[i] = cell.Value
		//}
		err = pf(line_num, row.Cells)
		if err != nil {
			if err.Error() == "JBREAK" {
				return false, line_num, nil
			} else if err.Error() == "JCONTINUE" {
				continue
			} else if !isContinue {
				return false, line_num, err
			}
		}
	}
	return true, line_num, nil
}

// ProcessLine64AndSave 读取excel文件，按行输出,修改后，可保存到新文件
// filename: 文件名
// pf: 处理每一行的函数 int64:行号，从1开始；[]*xlsx.Cell：该行的cell
// isContinue: pf函数报错后是否继续处理下一行
// jfile.JCONTINUE() 进行下个循环
// jfile.JBREAK()退出循环
// returns
// bool: 是否遍历完全
// int64:  处理到哪一行，从1开始
// error: 报错
func ProcessLine64AndSave(filename, sheetName, dstFile string, pf func(lineNum int64, cells []*xlsx.Cell) (err error), isContinue bool) (bDone bool, doneLineNum int64, err error) {
	if dstFile == "" {
		err = fmt.Errorf("invalid dstFile")
		return false, -1, err
	}
	var line_num int64 = 0
	// 打开 Excel 文件
	f, err := xlsx.OpenFile(filename)
	if err != nil {
		return false, -1, err
	}
	//
	// 获取工作表列表
	sheets := f.Sheets
	if len(sheets) == 0 {
		err = fmt.Errorf("Excel 文件中没有工作表")
		return false, -1, err
	}
	if sheetName == "" {
		sheetName = sheets[0].Name // 读取第一个工作表
	}
	bMatched := false
	var matchedSheet *xlsx.Sheet
	for _, sheet := range sheets {
		if sheet.Name == sheetName {
			bMatched = true
			matchedSheet = sheet
			break
		}
	}
	if !bMatched {
		err = fmt.Errorf("no such sheet")
		return false, -1, err
	}
	// 遍历行数据（跳过表头）
	for _, row := range matchedSheet.Rows {
		// 使用传进来的函数处理line
		line_num += 1
		// 业务逻辑
		//var colValueList = make([]string, len(row.Cells))
		//for i, cell := range row.Cells {
		//	colValueList[i] = cell.Value
		//}
		err = pf(line_num, row.Cells)
		if err != nil {
			if err.Error() == "JBREAK" {
				return false, line_num, nil
			} else if err.Error() == "JCONTINUE" {
				continue
			} else if !isContinue {
				return false, line_num, err
			}
		}
	}
	return true, line_num, nil
}
