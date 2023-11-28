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
	jlog.Info("test1," + OutputField(field))
	field = "test2"
	//if FieldNeedsQuotes(field) {
	//	field, _ = QuoteField(field, true)
	//	jlog.NInfo("test1," + field)
	//} else {
	//}
	jlog.Info("test1," + OutputField(field))
}
