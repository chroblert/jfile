package jexcel

import (
	"github.com/chroblert/jlog"
	"github.com/tealeg/xlsx/v2"
	"testing"
)

func TestProcessLine64v2(t *testing.T) {
	type args struct {
		filename   string
		sheetName  string
		pf         func(lineNum int64, cells []*xlsx.Cell) (err error)
		isContinue bool
	}
	tests := []struct {
		name            string
		args            args
		wantBDone       bool
		wantDoneLineNum int64
		wantErr         bool
	}{
		{
			name: "1",
			args: args{
				filename:  "D:\\Data\\250411-目标对象资产\\B方向机构子域名-250411 - Copy.xlsx",
				sheetName: "",
				pf: func(lineNum int64, cells []*xlsx.Cell) error {
					cells[len(cells)-1].SetString("test")
					jlog.Debug(len(cells), cells)
					return nil
				},
				isContinue: false,
				//opts:       nil,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotBDone, gotDoneLineNum, err := ProcessLine64(tt.args.filename, tt.args.sheetName, tt.args.pf, tt.args.isContinue)
			if (err != nil) != tt.wantErr {
				t.Errorf("ProcessLine64v2() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotBDone != tt.wantBDone {
				t.Errorf("ProcessLine64v2() gotBDone = %v, want %v", gotBDone, tt.wantBDone)
			}
			if gotDoneLineNum != tt.wantDoneLineNum {
				t.Errorf("ProcessLine64v2() gotDoneLineNum = %v, want %v", gotDoneLineNum, tt.wantDoneLineNum)
			}
		})
	}
}
