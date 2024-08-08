package gotest

import (
	"errors"
	"testing"
)

func Test_Division_not_pass(t *testing.T) {
	if i, e := Division1(6, 2); i != 3 || e != nil {
		t.Error("not pass")
	} else {
		t.Log("pass")
	}
}

func Test_Division_pass(t *testing.T) {
	t.Error("not pass always")
}

func Division1(a, b float64) (float64, error) {
	if b == 0 {
		return 0, errors.New("error")
	}

	return a / b, nil
}
