package jconvert

import (
	"encoding/json"
	"fmt"
	"github.com/chroblert/jfile"
	"github.com/chroblert/jfile/jcsv"
	"github.com/chroblert/jlog"
	"reflect"
	"slices"
	"strings"
)

// SimpleJson2CSV 将简单的json文件转为csv文件
//
// 以第一行json字符串中key作为csv的header，之后输出第一行中的key对应的value
//
// json文件中的每一行应该是一对一，不支持嵌套。支持数值、字符串、bool。数值可能会被转换为浮点型
func SimpleJson2CSV(srcFile, dstFile string) (err error) {
	dst_f := jlog.New(jlog.LogConfig{
		LogFullPath:      dstFile,
		InitCreateNewLog: true,
		StoreToFile:      true,
	})
	defer dst_f.Flush()
	// 示例 JSON 字符串
	// 记录第一行的key
	var keyList []string
	_, _, err = jfile.ProcessLine(srcFile, func(line_num int, line string) error {
		// 解析 JSON 数据到 map
		var data map[string]interface{}
		if err := json.Unmarshal([]byte(line), &data); err != nil {
			return err
		}
		// 将 map 中的值放入字符串列表
		var columnList []string
		var keyValue = make(map[string]string)
		tmpValue := ""
		bSkip := false
		for key, value := range data {
			switch reflect.TypeOf(value).Kind() {
			case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
				tmpValue = fmt.Sprintf("%d", reflect.ValueOf(value).Int())
				//stringList = append(stringList, fmt.Sprintf("%d", reflect.ValueOf(value).Int()))
			case reflect.Uint8, reflect.Uint, reflect.Uint16, reflect.Uint32, reflect.Uint64:
				//stringList = append(stringList, fmt.Sprintf("%d", reflect.ValueOf(value).Uint()))
				tmpValue = fmt.Sprintf("%d", reflect.ValueOf(value).Uint())
			case reflect.Float32, reflect.Float64:
				//stringList = append(stringList, fmt.Sprintf("%f", reflect.ValueOf(value).Float()))
				tmpValue = fmt.Sprintf("%f", reflect.ValueOf(value).Float())
			case reflect.String:
				//stringList = append(stringList, jcsv.EncodeField(value.(string)))
				tmpValue = jcsv.EncodeField(value.(string))
			case reflect.Bool:
				if value.(bool) {
					tmpValue = "true"
					//stringList = append(stringList, "true")
				} else {
					tmpValue = "false"
					//stringList = append(stringList, "false")
				}
			default:
				bSkip = true
			}
			if bSkip {
				continue
			}
			if line_num == 1 {
				keyList = append(keyList, jcsv.EncodeField(key))
			}
			keyValue[key] = tmpValue
		}
		if line_num == 1 {
			slices.Sort(keyList)
			dst_f.NInfo(strings.Join(keyList, ","))
		}
		for _, key := range keyList {
			if value, ok := keyValue[key]; ok {
				columnList = append(columnList, value)
			} else {
				columnList = append(columnList, "")
			}
		}
		dst_f.NInfo(strings.Join(columnList, ","))
		return jfile.JCONTINUE()
	}, false)
	dst_f.Flush()
	return err
}
