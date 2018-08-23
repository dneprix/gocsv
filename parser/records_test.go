package parser

import (
	"reflect"
	"testing"
)

func TestRecords(t *testing.T) {
	type args struct {
		records [][]string
	}
	tests := []struct {
		name string
		args args
		want [][]string
	}{
		{
			"1: simple",
			args{
				[][]string{
					{"a2"},
					{"1 1 +"},
				},
			},
			[][]string{{"2"}, {"2"}},
		},
		{
			"2: example",
			args{
				[][]string{
					{"b1 b2 +", "2 b2 3 * -", "3 ", " 2 0 /"},
					{"a1 ", "5 ", " ", "7 2 /"},
					{"c2 3 * ", "1 2 ", " ", "5 1 2 + 4 * + 3 -"},
				},
			},
			[][]string{
				{"-8", "-13", "3", "#ERR"},
				{"-8", "5", "0", "3.5"},
				{"0", "#ERR", "0", "14"},
			},
		},
		{
			"3: advanced",
			args{[][]string{{"b1", "b2", "a1 b1 +"}, {"1 2 +", "3 4 +", "c1 b1 a2 - -"}}},
			[][]string{{"7", "7", "14"}, {"3", "7", "10"}},
		},
		{
			"4: simple circle",
			args{[][]string{{"b1 a1 +", "a1"}}},
			[][]string{{"#ERR", "#ERR"}},
		},
		{
			"5: advanced circle",
			args{[][]string{{"b1 c1 +", "c1", "a1"}}},
			[][]string{{"#ERR", "#ERR", "#ERR"}},
		},
		{
			"6: fail refs",
			args{[][]string{{"b1"}}},
			[][]string{{"#ERR"}},
		},
		{
			"7: wrong rpn in refs",
			args{[][]string{{"b1+c1", "1", "2"}}},
			[][]string{{"#ERR", "1", "2"}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Records(tt.args.records); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Records() = %v, want %v", got, tt.want)
			}
		})
	}
}
