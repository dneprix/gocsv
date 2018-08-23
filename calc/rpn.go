package calc

import (
	"math"
	"strconv"
	"strings"
)

// Rpn function calculatres value from expression
func Rpn(expression string) string {
	items := strings.Fields(expression)

	stack := make([]float64, 0, len(items))
	for _, itm := range items {
		if v, err := strconv.ParseFloat(itm, 64); err == nil {
			stack = append(stack, v)
			continue
		}

		l := len(stack)
		if l < 2 {
			return ""
		}
		a, b := stack[l-2], stack[l-1]
		var r float64
		switch itm {
		case "+":
			r = a + b
		case "-":
			r = a - b
		case "*":
			r = a * b
		case "/":
			r = a / b
		default:
			return ""
		}
		stack[l-2] = r
		stack = stack[:l-1]
	}

	if len(stack) < 1 {
		return "0"
	} else if len(stack) == 1 && !math.IsInf(stack[0], 1) {
		return strconv.FormatFloat(stack[0], 'f', -1, 64)
	} else {
		return ""
	}

}
