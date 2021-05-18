package sum

func Sum(values ...int) int {
	if len(values) == 0 {
		return 0
	}

	var total int
	for _, v := range values {
		total += v
	}

	return total
}
