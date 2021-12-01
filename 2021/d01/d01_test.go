package d01

import (
	"bufio"
	"os"
	"strconv"
	"testing"
)

func TestPartA(t *testing.T) {
	tests := []struct {
		name         string
		measurements []int
		increases    int
	}{
		{"no input", []int{}, 0},
		{"single input", []int{0}, 0},
		{"sample", []int{199, 200, 208, 210, 200, 207, 240, 269, 260, 263}, 7},
		{"puzzle input", puzzleInput(), 1301},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			if PartA(test.measurements) != test.increases {
				t.Fail()
			}
		})
	}
}

func TestPartB(t *testing.T) {
	tests := []struct {
		name         string
		measurements []int
		increases    int
	}{
		{"no input", []int{}, 0},
		{"sample", []int{199, 200, 208, 210, 200, 207, 240, 269, 260, 263}, 5},
		{"puzzle input", puzzleInput(), 1346},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			increases := PartB(test.measurements)

			if increases != test.increases {
				t.Errorf("expected %d, got %d", test.increases, increases)
			}
		})
	}
}

func puzzleInput() []int {
	f, err := os.Open("input")

	if err != nil {
		panic("missing input file")
	}

	s := bufio.NewScanner(f)

	var measurements []int

	for s.Scan() {
		n, err := strconv.Atoi(s.Text())

		if err != nil {
			panic("invalid input file")
		}

		measurements = append(measurements, n)
	}

	return measurements
}
