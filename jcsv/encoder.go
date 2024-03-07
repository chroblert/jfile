package jcsv

import (
	"bytes"
	"encoding/csv"
	"strings"
	"unicode"
	"unicode/utf8"
)

// FieldNeedsQuotes 判断字段field是否需要用双引号括起来
//
// @param:comma_list: 默认rune(',')
func FieldNeedsQuotes(field string, comma_list ...rune) bool {
	if field == "" {
		return false
	}

	if field == `\.` {
		return true
	}
	comma := ','
	if len(comma_list) > 0 {
		comma = comma_list[0]
	}
	if comma < utf8.RuneSelf {
		for i := 0; i < len(field); i++ {
			c := field[i]
			if c == '\n' || c == '\r' || c == '"' || c == byte(comma) {
				return true
			}
		}
	} else {
		if strings.ContainsRune(field, comma) || strings.ContainsAny(field, "\"\r\n") {
			return true
		}
	}

	r1, _ := utf8.DecodeRuneInString(field)
	return unicode.IsSpace(r1)
}

// QuoteField 输出引起来的字段值
//
// @param: useCRLF: 是否使用CRLF换行
func QuoteField(field string, useCRLF bool) (quotedString string, err error) {
	buf := bytes.Buffer{}
	buf.WriteString("\"")
	for len(field) > 0 {
		// Search for special characters.
		i := strings.IndexAny(field, "\"\r\n")
		if i < 0 {
			i = len(field)
		}

		// Copy verbatim everything before the special character.

		if _, err = buf.WriteString(field[:i]); err != nil {
			return
		}
		field = field[i:]

		// Encode the special character.
		if len(field) > 0 {
			switch field[0] {
			case '"':
				_, err = buf.WriteString(`""`)
			case '\r':
				if !useCRLF {
					err = buf.WriteByte('\r')
				}
			case '\n':
				if useCRLF {
					_, err = buf.WriteString("\r\n")
				} else {
					err = buf.WriteByte('\n')
				}
			}
			field = field[1:]
			if err != nil {
				return
			}
		}
	}
	buf.WriteString("\"")
	return buf.String(), nil
}

// OutputField 判断后，输出字段对应的值
//
// @param: field: 输入值
//
// @param: comma: 文件由什么分割，默认：,
//
// @return: string: 空可能是报错
func EncodeField(field string, comma ...rune) string {
	if len(comma) == 0 {
		if FieldNeedsQuotes(field, ',') {
			field, err := QuoteField(field, true)
			if err != nil {
				return ""
			}
			return field
		}
		return field
	} else {
		if FieldNeedsQuotes(field, comma[0]) {
			field, err := QuoteField(field, true)
			if err != nil {
				return ""
			}
			return field
		}
		return field
	}

}

// EncodeStringList2StringList 判断字符串列表中的元素是否需要quote，并在quote之后转成字符串列表
func EncodeStringList2StringList(wordList []string, comma ...rune) (newWordList []string, err error) {
	delimiter := ','
	if len(comma) > 0 {
		delimiter = comma[0]
	}
	newWordList = make([]string, len(wordList))
	for k, word := range wordList {
		if FieldNeedsQuotes(word, delimiter) {
			field, err := QuoteField(word, true)
			if err != nil {
				return
			}
			newWordList[k] = field
		}
	}
	return
}

// DecodeString2List 将一行分割为字符串列表，使用传入的第一个comma作为分隔符。若不传，默认为逗号
func DecodeString2List(line string, comma_list ...rune) []string {
	comma := ','
	if len(comma_list) > 0 {
		comma = comma_list[0]
	}
	//
	rr := csv.NewReader(bytes.NewReader([]byte(line)))
	rr.Comma = comma
	word_list, err := rr.Read()
	if err != nil {
		return nil
	}
	return word_list
}
