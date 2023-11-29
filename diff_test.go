package jfile

import "testing"

func TestADiffB(t *testing.T) {
	ADiffB("D:\\T00ls\\01-GoProject\\AzureGov\\App\\outputs\\public-azure-global-dc-ip-range.csv-v4-original.txt", "D:\\T00ls\\01-GoProject\\AzureGov\\App\\outputs\\public-azure-gov-dc-ip-range.csv-v4-original.txt", "test.txt", 10)
}
