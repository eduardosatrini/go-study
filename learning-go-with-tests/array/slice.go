package array

func Sum(slice []int) int {

	var sum int

	for _, v := range slice {
		sum += v
	}

	return sum
}

func SumAllValues(slices ...[]int) (totalSumSlices []int) {

	for _, value := range slices {
		totalSumSlices = append(totalSumSlices, Sum(value))
	}

	return
}

func SumAllSlicesValues(slices ...[]int) (total int) {

	for _, value := range slices {
		total += Sum(value)
	}

	return
}
