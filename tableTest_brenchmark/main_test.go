package main

import (
	"fmt"
	"testing"
)

func TestCalulate(t *testing.T) {
	if Calculate(2) != 4 {
		t.Error("Expected 2 + 2 to equal 4")
	}
}

func TestTableCalculate(t *testing.T) {
	var tests = []struct {
		input    int
		expected int
	}{ // attaching another curly braces attached to a struct inputs it as a "IFFY"
		{2, 4},
		{3, 5},
		{-1, 1},
		{0, 2},
		{999999, 1000001},
	}

	for _, test := range tests {
		if output := Calculate(test.input); output != test.expected {
			t.Error("Test failed: {}  inputed, {} expected, received: {}", test.input, test.expected, output)
		}
	}
}
func TestOther(t *testing.T) {
	fmt.Println("Testing something else")
	fmt.Println("This shouldn't run with -run=calc")
}


func benchmarkCalculate(input int, b *testing.B) {
	for n := 0; n < b.N; n++ {
		Calculate(input)
	}
}

func BenchmarkCalculate100(b *testing.B) {
	benchmarkCalculate(100, b)
}
func BenchmarkCalculateNegative100(b *testing.B) {
	benchmarkCalculate(-100, b)
}
func BenchmarkCalculateNegative1(b *testing.B) {
	benchmarkCalculate(-1, b)
}
