package jcsv

import (
	"github.com/chroblert/jlog"
	"testing"
)

func TestQuoteField(t *testing.T) {
	defer jlog.Flush()
	field := "test,test\""
	//if FieldNeedsQuotes(field) {
	//	field, _ = QuoteField(field, true)
	//	jlog.NInfo("test1," + field)
	//}
	jlog.Info("test1," + EncodeField(field))
	field = "test2"
	//if FieldNeedsQuotes(field) {
	//	field, _ = QuoteField(field, true)
	//	jlog.NInfo("test1," + field)
	//} else {
	//}
	jlog.Info("test1," + EncodeField(field))
	word_list := DecodeString2List("test1,test2")
	jlog.Info(word_list)
}
