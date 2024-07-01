package gocsv

import (
	"bytes"
	"errors"
	"fmt"
	"github.com/chroblert/jlog"
	"reflect"
	"strings"
	"unicode"
	"unicode/utf8"
)

// MarshalStructToString 将结构体指针转为csv字符串
// in：结构体指针
// omitHeaders：是否丢弃头部字段
func MarshalStructToString(in interface{}, omitHeaders bool) (headers, out string, err error) {
	// 确保传入的是结构体指针。因为指针是可寻址，结构体无法寻址
	if reflect.ValueOf(in).Kind() != reflect.Ptr {
		err = fmt.Errorf("only support *struct")
		return
	}
	inValue, inType := getConcreteReflectValueAndType(in) // Get the concrete type (not pointer) (Slice<?> or Array<?>)
	// 确保inType是结构体
	switch inType.Kind() {
	case reflect.Struct:
	default:
		err = fmt.Errorf("only support *struct")
		return
	}
	inInnerType := inType
	inInnerStructInfo := getStructInfo(inInnerType) // Get the inner struct info to get CSV annotations
	csvHeadersLabels := make([]string, len(inInnerStructInfo.Fields))
	for i, fieldInfo := range inInnerStructInfo.Fields { // Used to write the header (first line) in CSV
		csvHeadersLabels[i] = fieldInfo.getFirstKey()
	}
	if !omitHeaders {
		headers, err = EncodeStringList2String(csvHeadersLabels, ',', false)
		if err != nil {
			jlog.Error(err)
			return "", "", err
		}
	}
	//
	for j, fieldInfo := range inInnerStructInfo.Fields {
		csvHeadersLabels[j] = ""
		inInnerFieldValue, err := getInnerField(inValue, false, fieldInfo.IndexChain) // Get the correct field header <-> position
		if err != nil {
			return "", "", err
		}
		csvHeadersLabels[j] = inInnerFieldValue
	}
	out, err = EncodeStringList2String(csvHeadersLabels, ',', false)
	if err != nil {
		jlog.Error(err)
		return "", "", err
	}
	return headers, out, nil
}

var errInvalidDelim = errors.New("csv: invalid field or comment delimiter")

func validDelim(r rune) bool {
	return r != 0 && r != '"' && r != '\r' && r != '\n' && utf8.ValidRune(r) && r != utf8.RuneError
}

// FieldNeedsQuotes reports whether our field must be enclosed in quotes.
// Fields with a Comma, fields with a quote or newline, and
// fields which start with a space must be enclosed in quotes.
// We used to quote empty strings, but we do not anymore (as of Go 1.4).
// The two representations should be equivalent, but Postgres distinguishes
// quoted vs non-quoted empty string during database imports, and it has
// an option to force the quoted behavior for non-quoted CSV but it has
// no option to force the non-quoted behavior for quoted CSV, making
// CSV with quoted empty strings strictly less useful.
// Not quoting the empty string also makes this package match the behavior
// of Microsoft Excel and Google Drive.
// For Postgres, quote the data terminating string `\.`.
func FieldNeedsQuotes(field string, Comma rune) bool {
	if field == "" {
		return false
	}

	if field == `\.` {
		return true
	}

	if Comma < utf8.RuneSelf {
		for i := 0; i < len(field); i++ {
			c := field[i]
			if c == '\n' || c == '\r' || c == '"' || c == byte(Comma) {
				return true
			}
		}
	} else {
		if strings.ContainsRune(field, Comma) || strings.ContainsAny(field, "\"\r\n") {
			return true
		}
	}

	r1, _ := utf8.DecodeRuneInString(field)
	return unicode.IsSpace(r1)
}

// EncodeStringList2String 将字符串列表转为csv字符串
// Comma: 分割符
// UseCRLF：是否使用\r\n作为换行符
func EncodeStringList2String(record []string, Comma rune, UseCRLF bool) (out string, err error) {
	bufferString := bytes.NewBufferString(out)
	if !validDelim(Comma) {
		return "", errInvalidDelim
	}

	for n, field := range record {
		if n > 0 {
			if _, err := bufferString.WriteRune(Comma); err != nil {
				return "", err
			}
		}

		// If we don't have to have a quoted field then just
		// write out the field and continue to the next field.
		if !FieldNeedsQuotes(field, Comma) {
			if _, err := bufferString.WriteString(field); err != nil {
				return "", err
			}
			continue
		}

		if err := bufferString.WriteByte('"'); err != nil {
			return "", err
		}
		for len(field) > 0 {
			// Search for special characters.
			i := strings.IndexAny(field, "\"\r\n")
			if i < 0 {
				i = len(field)
			}

			// Copy verbatim everything before the special character.
			if _, err := bufferString.WriteString(field[:i]); err != nil {
				return "", err
			}
			field = field[i:]

			// Encode the special character.
			if len(field) > 0 {
				var err error
				switch field[0] {
				case '"':
					_, err = bufferString.WriteString(`""`)
				case '\r':
					if !UseCRLF {
						err = bufferString.WriteByte('\r')
					}
				case '\n':
					if UseCRLF {
						_, err = bufferString.WriteString("\r\n")
					} else {
						err = bufferString.WriteByte('\n')
					}
				}
				field = field[1:]
				if err != nil {
					return "", err
				}
			}
		}
		if err := bufferString.WriteByte('"'); err != nil {
			return "", err
		}
	}
	return bufferString.String(), nil
}
