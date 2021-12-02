package d02

import (
	"strconv"
	"strings"
)

func PartA(instructions []string) int {
	var horizontal, depth int

	for _, instruction := range instructions {
		move := strings.Split(instruction, " ")

		dir := move[0]
		cnt, _ := strconv.Atoi(move[1])

		switch dir {
		case "forward":
			horizontal += cnt
		case "up":
			depth -= cnt
		case "down":
			depth += cnt
		}
	}

	return horizontal * depth
}
