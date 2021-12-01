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
