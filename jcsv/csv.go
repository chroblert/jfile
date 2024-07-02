package jcsv

import (
	"github.com/chroblert/jarr"

	"github.com/chroblert/jstr"

	"fmt"
	"github.com/chroblert/jfile"
	"github.com/chroblert/jlog"
	"os"
	"strings"
)

// suffix表示要在指定字段追加的内容 如：_reverse,_count
// pf:对指定列处理后，返回的字符串
func KeyProcesstring(src_file, column_name, delimiter string, column_id int, with_header bool, suffix string, pf func(string) string) (err error) {
	// 获取column_id
	line_num := 0
	// 删除原来文件中的_count,_order列
	bak2_file := fmt.Sprintf("%s-%s-%s", src_file, suffix, jstr.GenerateRandomString(8))
	f, err := os.OpenFile(bak2_file, os.O_WRONLY|os.O_TRUNC|os.O_CREATE, 0666)
	if err != nil {
		jlog.Fatal(err)
	}
	defer func() {
		if f != nil {
			f.Close()
		}
	}()
	b_contain_count_order := false
	jfile.ProcessLine(src_file, func(tmp_line_num int64, line string) error {
		line_num += 1
		word_list := strings.Split(line, delimiter)
		// 获取column_id
		if column_id == -1 {
			if line_num == 1 {
				if !jarr.EleInArr(column_name, word_list) {
					jlog.Errorf("column_name:%s,none\n", column_name)
					return fmt.Errorf("format err")
				} else {
					//jlog.Info("word_list:", word_list, ",column_name:", column_name)
					for k, v := range word_list {
						if v == column_name {
							column_id = k
						}
					}
				}
			}
		} else {
			if len(word_list) < column_id+1 {
				return fmt.Errorf("format err")
			}
		}
		//
		if line_num == 1 {
			if with_header {
				// 处理
				// 长度不够，不包含suffix
				if len(word_list) < column_id+2 {
					return fmt.Errorf("JBREAK")
				} else {
					column_name = word_list[column_id]
					if word_list[column_id+1] == fmt.Sprintf("%s%s", column_name, suffix) {
						b_contain_count_order = true
					}
				}
			}
		}
		// 写入新文件
		if b_contain_count_order {
			new_word_list := make([]string, column_id+1)
			copy(new_word_list, word_list[:column_id+1])
			for k, v := range word_list {
				if k > column_id+1 {
					new_word_list = append(new_word_list, v)
				}
			}
			new_line := fmt.Sprintf("%s\n", strings.Join(new_word_list, delimiter))
			//jlog.Info(new_line)
			// 写入文件
			f.Write([]byte(new_line))
		}
		return nil
	}, false)
	if column_id == -1 {
		jlog.Info("column_id == -1")
		return fmt.Errorf("no such column")
	}
	if b_contain_count_order {
		jlog.Warnf("[!] 该文件包含%s%s，进行删除\n", column_name, suffix)
		// 删除原文件
		os.Remove(src_file)
		// 拷贝新文件
		jfile.FileCopy(bak2_file, src_file, true)
	} else {

		jlog.Infof("[!] 该文件不包含%s%s列\n", column_name, suffix)
	}
	if f != nil {
		f.Close()
		os.Remove(bak2_file)
	}
	// 存储特定列中值的数量
	line_num = 0
	// 遍历文件行
	// 备份原文件
	bak_file := fmt.Sprintf("%s-%s", src_file, jstr.GenerateRandomString(8))
	err = jfile.FileCopy(src_file, bak_file, true)
	if err != nil {
		jlog.Fatal(err)
	}
	// 删除原文件
	err = os.Remove(src_file)
	if err != nil {
		jlog.Fatal(err)
	}
	jlog.Infof("[!] 将源文件%s，移动到%s\n", src_file, bak_file)
	// 写入文件
	f, err = os.OpenFile(src_file, os.O_WRONLY|os.O_TRUNC|os.O_CREATE, 0666)
	if err != nil {
		jlog.Fatal(err)
	}
	defer func() {
		if f != nil {
			f.Close()
		}
	}()
	line_num = 0
	// 将添加_count,_order的数据写入文件
	jfile.ProcessLine(bak_file, func(tmp_line_num int64, line string) error {
		line_num += 1
		word_list := strings.Split(line, delimiter)
		if len(strings.TrimSpace(line)) == 0 || len(word_list) < column_id+1 {
			return nil
		}
		//jlog.Info("len:", len(word_list), strings.Join(word_list, delimiter))
		column_value := word_list[column_id]
		new_word_list := make([]string, column_id+1)
		//new_word_list = word_list[:column_id+1]
		copy(new_word_list, word_list[:column_id+1])
		if line_num == 1 {
			if with_header {
				column_suffix_name := fmt.Sprintf("%s%s", column_value, suffix)

				new_word_list = append(new_word_list, column_suffix_name)
			} else {
				// 若为空，则添加空字符串
				if len(strings.TrimSpace(word_list[column_id])) == 0 {
					new_word_list = append(new_word_list, "")
				} else {
					new_word_list = append(new_word_list, pf(column_value))
				}
			}
		} else {
			// 若为空，则添加空字符串
			if len(strings.TrimSpace(word_list[column_id])) == 0 {
				new_word_list = append(new_word_list, "")
			} else {
				new_word_list = append(new_word_list, pf(column_value))
			}
		}
		// 添加后面一部分
		if len(word_list) > 1 && len(word_list) > column_id+1 {
			new_word_list = append(new_word_list, word_list[column_id+1:]...)
		}
		// 写入文件
		f.Write([]byte(strings.Join(new_word_list, delimiter)))
		f.Write([]byte("\n"))
		return nil
	}, false)
	// 删除bak文件
	os.Remove(bak_file)
	jlog.Infof("[*] 将添加了%s的数据写入文件\n", suffix)
	return nil
}
