package calc

import "testing"

func TestRpn(t *testing.T) {
	type args struct {
		expression string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"1: +", args{"125 213 -8 + +"}, "330"},
		{"2: + spaces", args{"   125    213 -8   +  +"}, "330"},
		{"3: +-", args{"3 11 + 5 -4 - -"}, "5"},
		{"4: +-*", args{"2   3 -11 + 5 - *"}, "-26"},
		{"5: +-*/", args{"90 2     3 11 + 5 - * /"}, "5"},
		{"6: digit", args{"15"}, "15"},
		{"7: digit", args{"45   "}, "45"},
		{"8: fail", args{"23 67"}, ""},
		{"9: fail", args{"23 -"}, ""},
		{"10: fail", args{"+"}, ""},
		{"11: fail", args{"   +"}, ""},
		{"12: fail", args{"+   "}, ""},
		{"13: fail", args{"+ 10 1000"}, ""},
		{"14: empty", args{""}, "0"},
		{"15: empty", args{"10 10 -"}, "0"},
		{"16: error", args{"3 0 /"}, ""},
		{"17: error", args{"12 a +"}, ""},
		{"18: error", args{"abc"}, ""},
		{"19: float", args{"2.2 3.45 +"}, "5.65"},
		{"20: float", args{"-2.2 3.45 +"}, "1.25"},
		{"21: wrong rnp", args{"1+1"}, ""},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Rpn(tt.args.expression); got != tt.want {
				t.Errorf("Rpn() = %v, want %v", got, tt.want)
			}
		})
	}
}
