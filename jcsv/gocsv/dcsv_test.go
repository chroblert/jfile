package gocsv

import (
	"github.com/chroblert/jlog"
	"testing"
)

func TestEncodeStructToString(t *testing.T) {
	type args struct {
		in          interface{}
		omitHeaders bool
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "1",
			args: args{
				omitHeaders: false,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ipPrefixI := IPPrefixStruct{
				IPType:             1,
				IPPrefix:           "ipPrefix",
				Region:             "region",
				Service:            "service",
				NetworkBorderGroup: "networkBorderGroup",
				RegionType:         1,
			}
			_, str, err := MarshalStructToString(ipPrefixI, tt.args.omitHeaders)
			jlog.Debug(str)
			ipPrefixI2 := IPPrefixStruct{
				Service:            "service",
				IPPrefix:           "ipPrefix",
				IPType:             1,
				Region:             "region",
				NetworkBorderGroup: "networkBorderGroup",
				RegionType:         1,
			}
			_, str, err = MarshalStructToString(ipPrefixI2, tt.args.omitHeaders)
			if err != nil {
				jlog.Fatal(err)
			}
			jlog.Debug(str)

		})
	}
}

type IPPrefixStruct struct {
	IPType             int // ipv4 or ipv6
	IPPrefix           string
	Region             string
	RegionType         int // gov or not gov
	Service            string
	NetworkBorderGroup string
}

func TestMarshalList2String(t *testing.T) {
	type args struct {
		record  []string
		Comma   rune
		UseCRLF bool
	}
	tests := []struct {
		name    string
		args    args
		wantOut string
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "1",
			args: args{
				record:  []string{"1,\"\n", "22", "333"},
				Comma:   ',',
				UseCRLF: false,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotOut, err := EncodeStringList2String(tt.args.record, tt.args.Comma, tt.args.UseCRLF)
			jlog.Debug(gotOut)
			if (err != nil) != tt.wantErr {
				t.Errorf("EncodeStringList2String() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotOut != tt.wantOut {
				t.Errorf("EncodeStringList2String() gotOut = %v, want %v", gotOut, tt.wantOut)
			}
		})
	}
}
