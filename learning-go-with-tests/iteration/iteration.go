package iteration

func Iterations(c string, total int) string {

	var value string
	for i := 0; i < total; i++ {
		value += c
	}

	return value
}
