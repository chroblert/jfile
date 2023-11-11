package jcsv

import (
	"github.com/chroblert/jstr"
)

func ReverseString(src_file, column_name, delimiter string, column_id int, with_header bool) (err error) {
	return KeyProcesstring(src_file, column_name, delimiter, column_id, with_header, "_reverse", jstr.ReverseString)
}
