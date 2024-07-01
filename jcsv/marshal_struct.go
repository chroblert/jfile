package jcsv

import (
	"errors"
	"github.com/chroblert/jfile/jcsv/gocsv"
	"unicode/utf8"
)

// MarshalStructToString 将结构体转为csv字符串
// in：结构体
// omitHeaders：是否丢弃头部字段
func MarshalStructToString(in interface{}, omitHeaders bool) (headers, out string, err error) {
	return gocsv.MarshalStructToString(in, omitHeaders)
}

var errInvalidDelim = errors.New("csv: invalid field or comment delimiter")

func validDelim(r rune) bool {
	return r != 0 && r != '"' && r != '\r' && r != '\n' && utf8.ValidRune(r) && r != utf8.RuneError
}

// EncodeStringList2String 将字符串列表转为csv字符串
// Comma: 分割符
// UseCRLF：是否使用\r\n作为换行符
func EncodeStringList2String(record []string, Comma rune, UseCRLF bool) (out string, err error) {
	return gocsv.EncodeStringList2String(record, Comma, UseCRLF)
}
