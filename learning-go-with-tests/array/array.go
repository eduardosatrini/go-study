package array

func Sum(num []int) int {

	var sum int
	for _, v := range num {
		sum += v
	}

	return sum
}
