package jtsv

// 将tsv转为csv
//
// with_header:src_file中是否有表头
//
// coroutin_count:协程数量
func TransTsv2CSVConcurrent(src_file, dst_file string, with_header bool, coroutin_count int) (err error) {
	//dst_f, err := os.OpenFile(dst_file, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0666)
	//if err != nil {
	//	return
	//}
	//defer dst_f.Close()
	////
	//// 异步
	//a := jasync.New(true)
	////total_line := jfile.GetLineCount(src_file)
	////bar := progressbar.Default(total_line)
	//jfile.ProcessLine64(src_file, func(line_num int64, line string) error {
	//	word_list := strings.Split(line, "\t")
	//	// 判断with_header
	//	if with_header && line_num == 1 {
	//		new_word_list := make([]string, len(word_list))
	//		for k, v := range word_list {
	//			if strings.Contains(v, ",") {
	//				if !strings.Contains(v, "\"") {
	//					new_word_list[k] = fmt.Sprintf("\"%s\"", v)
	//				} else {
	//					dubble_quote := strings.ReplaceAll(v, "\"", "\"\"")
	//					new_word_list[k] = fmt.Sprintf("\"%s\"", dubble_quote)
	//				}
	//			} else {
	//				new_word_list[k] = word_list[k]
	//			}
	//		}
	//		new_line := strings.Join(new_word_list, ",")
	//		dst_f.Write([]byte(new_line + "\n"))
	//		//bar.Add(1)
	//		return jfile.JCONTINUE()
	//	}
	//	a.Add("", func(in_line_num int, in_word_list []string, in_f *os.File) {
	//		new_word_list := make([]string, len(in_word_list))
	//		for k, v := range in_word_list {
	//			if strings.Contains(v, ",") {
	//				if !strings.Contains(v, "\"") {
	//					new_word_list[k] = fmt.Sprintf("\"%s\"", v)
	//				} else {
	//					dubble_quote := strings.ReplaceAll(v, "\"", "\"\"")
	//					new_word_list[k] = fmt.Sprintf("\"%s\"", dubble_quote)
	//				}
	//			} else {
	//				new_word_list[k] = in_word_list[k]
	//			}
	//		}
	//		new_line := strings.Join(new_word_list, ",")
	//		in_f.Write([]byte(new_line + "\n"))
	//	}, func() {}, line_num, word_list, dst_f)
	//	// 达到1000则运行
	//	if line_num%coroutin_count == 0 {
	//		a.Run(coroutin_count)
	//		a.Wait()
	//		a.Clean()
	//	}
	//	return jfile.JCONTINUE()
	//}, false)
	//jlog.Info("[!] last 异步执行")
	//a.Run(coroutin_count)
	//jlog.Info("[!] last 等待")
	//a.Wait()
	//a.Clean()
	//jlog.Info("[!] 所有执行完成")
	return
}
