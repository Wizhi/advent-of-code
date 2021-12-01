package d01

func PartA(measurements []int) int {
	var count int

	for i := 1; i < len(measurements); i++ {
		if measurements[i-1] < measurements[i] {
			count++
		}
	}

	return count
}

func PartB(measurements []int) int {
	const window int = 3

	var count int

	for i := 0; i < len(measurements)-window; i++ {
		var a, b int

		for j := i; j < i+window; j++ {
			a += measurements[j]
			b += measurements[j+1]
		}

		if a < b {
			count++
		}
	}

	return count
}
