package jfile

import "testing"

func TestMerge(t *testing.T) {
	Merge("test.txt", 2, "D:\\T00ls\\jfile\\jcsv\\logs\\app.log", "D:\\T00ls\\jfile\\jcsv\\logs\\app.csv")
}
