package jcsv

import (
	"github.com/chroblert/jfile"
	"github.com/chroblert/jlog"
	"testing"
)

func TestGetUnique2FileByName(t *testing.T) {
	file1 := "oneforall-subdomain-collect.csv"
	file2 := "endpoints_flint_custom.csv"
	suffix_file := "all_endpoints_suffix.txt"
	suffix_list, _ := GetColumnStringValuesByID(suffix_file, ",", 0, false)
	filtered_file := "unique_subdoamin.txt"
	jlog.Info(GetQualifiedUnique2FileByName(suffix_list, file1, file2, ",", "subdomain", "fqdn", true, true, filtered_file))
	jlog.Info(jfile.UniqueInSameFile(filtered_file))
}
