package parser

import (
	"reflect"
	"testing"
)

func TestCsv(t *testing.T) {
	type args struct {
		filepath string
	}
	tests := []struct {
		name    string
		args    args
		want    [][]string
		wantErr bool
	}{
		{
			"file",
			args{"../files/test1.csv"},
			[][]string{{"1", "2", "3"}, {"4", "5", "6"}, {"7", "8", "9"}},
			false,
		},
		{"fail", args{"file.csv"}, nil, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Csv(tt.args.filepath)
			if (err != nil) != tt.wantErr {
				t.Errorf("Csv() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Csv() = %v, want %v", got, tt.want)
			}
		})
	}
}
