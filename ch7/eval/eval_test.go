package eval

import (
	"fmt"
	"testing"
)

func TestEval(t *testing.T) {
	tests := []struct {
		expr string
		env  Env
		want string
	}{
		{},
	}
	var preExpr string
	for _, test := range tests {
		if test.expr != preExpr {
			fmt.Println(test.expr)
			preExpr = test.expr
		}
		expr, err := Parse(test.expr)
		if err != nil {
			t.Error(err)
			continue
		}
		got := fmt.Sprintf("%.6g", expr.Eval(test.env))
	}
}
