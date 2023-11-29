package jconvert

import "testing"

func TestJson2CSV(t *testing.T) {
	SimpleJson2CSV("../../output/traceroute-nodesinfo.json", "../../output/test2.csv")
}
