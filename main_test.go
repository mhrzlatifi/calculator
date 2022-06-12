package main

import "testing"

func TestCalc(t *testing.T) {
	tests := []struct {
		num1, num2 int
		op         string
		result     float32
		errMsg     string
	}{
		{4, 5, "*", 20.0, ""},
		{-4, -6, "*", 24.0, ""},
		{-2, 5, "*", -10.0, ""},
		{8, 4, "/", 2.0, ""},
		{16, -2, "/", -8.0, ""},
		{-126, -2, "/", 63.0, ""},
		{58, 0, "/", 0, "division by zero"},
		{10, 3, "-", 7.0, ""},
		{25, -5, "-", 30.0, ""},
		{-25, -5, "-", -20, ""},
		{55, 5, "+", 60.0, ""},
		{0, -5, "+", -5.0, ""},
		{-25, -5, "+", -30.0, ""},
		{10, 2, "%", 0.0, ""},
		{10, 3, "%", 1, ""},
		{-25, 5, "%", 0, ""},
		{-25, 5, "^", 0, "operator is not supported"},
	}

	for i := range tests {

		res, err := calc(tests[i].num1, tests[i].num2, tests[i].op)
		if err != nil && err.Error() != tests[i].errMsg {
			t.Errorf("error Msg should be equal to %s but get %s", tests[i].errMsg, err.Error())
			return
		}

		if res != tests[i].result {
			t.Errorf("%d %s %d : expected %g but get %g", tests[i].num1, tests[i].op, tests[i].num2, tests[i].result, res)
		}

	}
}
