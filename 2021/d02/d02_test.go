package d02

import (
	"os"
	"strings"
	"testing"
)

func TestPartA(t *testing.T) {
	tests := []struct {
		name         string
		instructions []string
		expected     int
	}{
		{"sample", []string{"forward 5", "down 5", "forward 8", "up 3", "down 8", "forward 2"}, 150},
		{"puzzle", puzzleInput(), 1499229},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := PartA(test.instructions)

			if result != test.expected {
				t.Errorf("expected %d, got %d", test.expected, result)
			}
		})
	}
}

func TestPartB(t *testing.T) {
	tests := []struct {
		name         string
		instructions []string
		expected     int
	}{
		{"sample", []string{"forward 5", "down 5", "forward 8", "up 3", "down 8", "forward 2"}, 900},
		{"puzzle", puzzleInput(), 1340836560},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := PartB(test.instructions)

			if result != test.expected {
				t.Errorf("expected %d, got %d", test.expected, result)
			}
		})
	}
}

func puzzleInput() []string {
	file, _ := os.ReadFile("input")
	lines := strings.Split(string(file), "\n")

	return lines[:len(lines)-1]
}
