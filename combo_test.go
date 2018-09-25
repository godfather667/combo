// combo_test.go - This test package confirms that combo functions are working.
// This is a bit of overkill - But testing is included for a more complete demo package.
package main

import (
	"fmt"
	"testing"
)

// var s1D - Variable containing expected data from run1D() Function.
var s1D = []string{"[100]",
	"[111]",
	"[100 111]",
	"[223]",
	"[100 223]",
	"[111 223]",
	"[100 111 223]",
	"[44]",
	"[100 44]",
	"[111 44]",
	"[100 111 44]",
	"[223 44]",
	"[100 223 44]",
	"[111 223 44]",
	"[100 111 223 44]",
	"[552]",
	"[100 552]",
	"[111 552]",
	"[100 111 552]",
	"[223 552]",
	"[100 223 552]",
	"[111 223 552]",
	"[100 111 223 552]",
	"[44 552]",
	"[100 44 552]",
	"[111 44 552]",
	"[100 111 44 552]",
	"[223 44 552]",
	"[100 223 44 552]",
	"[111 223 44 552]",
	"[100 111 223 44 552]",
}

// var s2D - Variable containing expected data from run2D() Function
var s2D = []string{"[[0 5 9]]",
	"[[1 1 5]]",
	"[[0 5 9] [1 1 5]]",
	"[[6 2 3]]",
	"[[0 5 9] [6 2 3]]",
	"[[1 1 5] [6 2 3]]",
	"[[0 5 9] [1 1 5] [6 2 3]]",
	"[[4 2 4]]",
	"[[0 5 9] [4 2 4]]",
	"[[1 1 5] [4 2 4]]",
	"[[0 5 9] [1 1 5] [4 2 4]]",
	"[[6 2 3] [4 2 4]]",
	"[[0 5 9] [6 2 3] [4 2 4]]",
	"[[1 1 5] [6 2 3] [4 2 4]]",
	"[[0 5 9] [1 1 5] [6 2 3] [4 2 4]]",
}

// TestRun1D - Tests the function run1D and combo1D and verifies correct result.
func TestRun1D(t *testing.T) {
	c1 := run1D()
	for i, v := range c1 {
		if s1D[i] != fmt.Sprint(v.cm) {
			fmt.Println("Error: s1D = ", s1D[i], "  v = ", v.cm)
			t.Fail()
		}
	}
}

// TestRun2D - Tests the function run2D and combo2D and verifies correct result.
func TestRun2D(t *testing.T) {
	c2 := run2D()
	for i, v := range c2 {
		if s2D[i] != fmt.Sprint(v.cm) {
			fmt.Println("Error: s2D = ", s2D[i], "  v = ", v.cm)
			t.Fail()
		}
	}
}
