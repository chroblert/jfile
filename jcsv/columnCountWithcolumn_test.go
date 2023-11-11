package jcsv

import (
	"testing"
)

func TestStatisticColumnCountWithKey(t *testing.T) {
	StatisticUniqueColumnCountWithKey("test.csv", "unknown_ip", "city", ",", -1, -1, true, "_count")
}
