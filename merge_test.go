package jfile

import "testing"

func TestMerge(t *testing.T) {
	Merge("scan_res.csv", 1, "D:\\T00ls\\02-PyProject\\traceroute\\microsoft_scan_res_v1.csv", "D:\\T00ls\\02-PyProject\\traceroute\\trace_scan_res_2_v1.csv")
	//UniqueInSameFile("scan_res.csv")
}
