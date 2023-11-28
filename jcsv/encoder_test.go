package jcsv

import (
	"fmt"
	"github.com/chroblert/jasync"
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
	a := jasync.NewAR(1000)
	for i := 0; i < 10000; i++ {
		a.Init(fmt.Sprintf("task-%d", i)).CAdd(func(i int) {
			jlog.NInfof("test1,test%d\n", i)
			word_list := DecodeString2List(fmt.Sprintf("test1,test%d\n", i))
			jlog.Info(word_list)
		}, i).CDO()
	}
	a.Wait()
	jlog.Info("test1," + EncodeField(field))
	word_list := DecodeString2List("test1,test2")
	jlog.Info(word_list)
}
